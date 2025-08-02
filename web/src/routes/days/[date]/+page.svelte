<script>
import { base } from "$app/paths";
import BackArrowIcon from "$lib/icons/BackArrow.svelte";
import PlusIcon from "$lib/icons/Plus.svelte";
import CheckmarkIcon from "$lib/icons/Checkmark.svelte";
import XMarkIcon from "$lib/icons/CloseXMark.svelte";
import PencilIcon from "$lib/icons/Pencil.svelte";
import DocIcon from "$lib/icons/DocFilePage.svelte";
import { calculateDay, remainingBudgetFromDays } from "$lib/budget.js";
import { dateToYMDString } from "$lib/dateToYMDString.js";
import { dateGetWeekNum } from "$lib/dateGetWeekNum.js";
import { onMount } from "svelte";
let { data } = $props();
const date = (() => {
    const [year, month, day] = data.date.split("-");
    const d = new Date(year, month - 1, day);
    if (!isNaN(d)) {
        return d;
    } else {
        window.location = `${base}/days/${data.date}/err/invalid-date`
    }
})();
const dateYMD = dateToYMDString(date);
const weekDayName = [
    "Sunday", "Monday", "Tuesday", "Wednesday",
    "Thrusday", "Friday", "Saturday"
][date.getDay()];
const weekDayKey = [
    "sun", "mon", "tue", "wed",
    "thu", "fri", "sat"
][date.getDay()];
const monthName = [
    "January", "February", "March", "April", "May",
    "June", "July", "August", "September",
    "October", "November", "December"
][date.getMonth()];
const weekNum = dateGetWeekNum(date);
let prevDaysArray = [];
let selectedDay;
const dateAsNum = parseInt(
    data.date.replaceAll("-", "")
);
data.days.forEach((day) => {
    const iterationDateAsNum = parseInt(
        day.date.replaceAll("-", "")
    );
    if (iterationDateAsNum < dateAsNum) {
        prevDaysArray.push(day);
    } else if (iterationDateAsNum == dateAsNum) {
        selectedDay = day;
    }
})
if (!selectedDay) {
    window.location = `${base}/days/${data.date}/err/dayNotFound`
}

let {
    foodBudget: dayStartFoodBudget,
    kitchenBudget: dayStartKitchenBudget,
    floorBudget: dayStartFloorBudget
} = remainingBudgetFromDays(
    data.week.startFoodBudget,
    data.week.startKitchenBudget,
    data.week.startFloorBudget,
    prevDaysArray,
    data.week.weeklyPay
);

let dayResults = calculateDay(
    dayStartFoodBudget,
    dayStartKitchenBudget,
    dayStartFloorBudget,
    selectedDay,
    prevDaysArray?.length >= 1 ?
        null : data.week.weeklyPay
        /* use weeklyPay on this day
        if there are no prevDays,
        which would mean today is the first day */
);

const foodCostsTotal = dayResults.foodExpenses;
const totalKitchenHourlyEarned = dayResults.totalKitchenHourlyEarned;
const totalFloorHourlyEarned = dayResults.totalFloorHourlyEarned;
const totalKitchenSpecialEarned = dayResults.totalKitchenSpecialEarned;
const totalFloorSpecialEarned = dayResults.totalFloorSpecialEarned;
const floorHourlyArray = dayResults.floorHourly;
const floorSpecialArray = dayResults.floorSpecialPay;
const kitchenHourlyArray = dayResults.kitchenHourly;
const kitchenSpecialArray = dayResults.kitchenSpecialPay;

const dayFinalFoodBudget = dayResults.foodBudgetFinal;
const dayFinalKitchenBudget = dayResults.kitchenBudgetFinal;
const dayFinalFloorBudget = dayResults.floorBudgetFinal;
</script>
<div class="grid page" style="margin-top: 4rem; margin-bottom: 10rem;">
    <div class="content">
        <div class="flex">
            <a class="button faint" href={base}>
                <BackArrowIcon></BackArrowIcon>
                Back
            </a>
        </div>
        <h3 style="margin-bottom: 0px;">{weekDayName}'s Report</h3>
        <p style="margin-top: 0.4rem;">{monthName} {date.getDate()}, {date.getFullYear()}{
            localStorage.getItem("budgeter:showWeekNums") == "true" ?
                `, W${weekNum}` : ""
        }</p>
        <div class="flex" style="margin-bottom: 2rem;">
            <a class="button alt" href="{base}/edit-day/{dateYMD}">
                <PencilIcon></PencilIcon>
                Edit Day
            </a>
            <a class="button alt" href="{base}/reports/{dateYMD}">
                <DocIcon></DocIcon>
                View Week Report
            </a>
        </div>
        <div>
            <div style="display: inline-block; margin-bottom: 2rem;">
                <div style="border: 0.2rem solid var(--border); border-radius: 0.8rem; min-width: 18rem;">
                <table style="border: none;">
                    <thead>
                        <tr>
                            <th></th>
                            <th>Food Budget</th>
                            <th>Kitchen Budget</th>
                            <th>Floor Budget</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr>
                            <td>Day Start</td>
                            <td>${dayStartFoodBudget.toLocaleString(
                                "en-US", {
                                minimumFractionDigits: 2,
                                maximumFractionDigits: 2
                            })}</td>
                            <td>${dayStartKitchenBudget.toLocaleString(
                                "en-US", {
                                minimumFractionDigits: 2,
                                maximumFractionDigits: 2
                            })}</td>
                            <td>${dayStartFloorBudget.toLocaleString(
                                "en-US", {
                                minimumFractionDigits: 2,
                                maximumFractionDigits: 2
                            })}</td>
                        </tr>
                        <tr>
                            <td>Increases</td>
                            <td>+${(selectedDay.foodBudgetIncrease ?? 0).toLocaleString(
                                "en-US", {
                                minimumFractionDigits: 2,
                                maximumFractionDigits: 2
                            })}</td>
                            <td>+${(selectedDay.kitchenBudgetIncrease ?? 0).toLocaleString(
                                "en-US", {
                                minimumFractionDigits: 2,
                                maximumFractionDigits: 2
                            })}</td>
                            <td>+${(selectedDay.floorBudgetIncrease ?? 0).toLocaleString(
                                "en-US", {
                                minimumFractionDigits: 2,
                                maximumFractionDigits: 2
                            })}</td>
                        </tr>
                        <tr>
                            <td>Expenses</td>
                            <td>-${foodCostsTotal.toLocaleString(
                                "en-US", {
                                minimumFractionDigits: 2,
                                maximumFractionDigits: 2
                            })}</td>
                            <td>-${
                                dayResults.kitchenExpenses.toLocaleString(
                                    "en-US", {
                                    minimumFractionDigits: 2,
                                    maximumFractionDigits: 2
                                })
                            }</td>
                            <td>-${
                                dayResults.floorExpenses.toLocaleString(
                                    "en-US", {
                                    minimumFractionDigits: 2,
                                    maximumFractionDigits: 2
                                })
                            }</td>
                        </tr>
                        <tr>
                            <td>Day Final</td>
                            <td>${dayFinalFoodBudget.toLocaleString(
                                "en-US", {
                                minimumFractionDigits: 2,
                                maximumFractionDigits: 2
                            })}</td>
                            <td>${dayFinalKitchenBudget.toLocaleString(
                                "en-US", {
                                minimumFractionDigits: 2,
                                maximumFractionDigits: 2
                            })}</td>
                            <td>${dayFinalFloorBudget.toLocaleString(
                                "en-US", {
                                minimumFractionDigits: 2,
                                maximumFractionDigits: 2
                            })}</td>
                        </tr>
                    </tbody>
                </table>
                </div>
            </div>
        </div>
        <div>
            <div style="display: inline-block; margin-bottom: 2rem;">
                <p>Food Costs</p>
                <div style="border: 0.2rem solid var(--border); border-radius: 0.8rem; min-width: 18rem;">
                <table style="border: none;">
                    <thead>
                        <tr>
                            <th>Expense</th>
                            <th>Notes</th>
                        </tr>
                    </thead>
                    <tbody>
                        {#each selectedDay.foodCosts as row}
                            <tr>
                                <td>${row.cost.toLocaleString(
                                    "en-US", {
                                    minimumFractionDigits: 2,
                                    maximumFractionDigits: 2
                                })}</td>
                                <td style="white-space: pre-wrap;">{row.notes}</td>
                            </tr>
                        {/each}
                        <tr>
                            <th>${foodCostsTotal.toLocaleString(
                                "en-US", {
                                minimumFractionDigits: 2,
                                maximumFractionDigits: 2
                            })}</th>
                            <th>Total</th>
                        </tr>
                    </tbody>
                </table>
                </div>
            </div>
        </div>
        <div>
        <div style="display: inline-block; margin-bottom: 2rem;">
            <div class="flex" style="justify-content: space-between;">
                <p>Kitchen Workers</p>
                <p class="fg0">
                    {kitchenHourlyArray.length} employees
                </p>
            </div>
            <table>
                <thead>
                    <tr>
                        <th>Employee</th>
                        <th>Hours Worked</th>
                        <th>Amount Earned</th>
                    </tr>
                </thead>
                <tbody>
                    {#each kitchenHourlyArray as row}
                        <tr>
                            <td>{row.employee.name}</td>
                            <td>{row.hours}</td>
                            <td>${
                                (row.hours * row.employee.wage).toLocaleString(
                                    "en-US", {
                                    minimumFractionDigits: 2,
                                    maximumFractionDigits: 2
                                })
                            }</td>
                        </tr>
                    {/each}
                    <tr>
                        <th>Total</th>
                        <th></th>
                        <th>${
                            totalKitchenHourlyEarned.toLocaleString(
                                "en-US", {
                                minimumFractionDigits: 2,
                                maximumFractionDigits: 2
                            })
                        }</th>
                    </tr>
                </tbody>
            </table>
        </div>
        </div>
        {#if kitchenSpecialArray.length >= 1}
            <div>
            <div style="display: inline-block; margin-bottom: 2rem;">
                <div class="flex" style="justify-content: space-between;">
                    <p>Kitchen Workers <span class="fg0">(Special Pay)</span></p>
                    <p class="fg0">
                        {#if kitchenSpecialArray.length == 1}
                            1 employee
                        {:else}
                            {kitchenSpecialArray.length} employees
                        {/if}
                    </p>
                </div>
                <table>
                    <thead>
                        <tr>
                            <th>Employee</th>
                            <th>Worked Today</th>
                            <th>Hours Worked</th>
                            <th>Amount Earned</th>
                        </tr>
                    </thead>
                    <tbody>
                        {#each kitchenSpecialArray as row}
                            <tr>
                                <td>{row.employee.name}</td>
                                <td>
                                    {#if row.workedToday === true}
                                        <div class="flex compact-gap text yay" style="align-items: center; align-content: center;">
                                            <CheckmarkIcon></CheckmarkIcon>
                                            <span>Yes</span>
                                        </div>
                                    {:else if row.workedToday === false}
                                        <div class="flex compact-gap text ohno" style="align-items: center; align-content: center;">
                                            <XMarkIcon></XMarkIcon>
                                            <span>No</span>
                                        </div>
                                    {:else if row.workedToday == null}
                                        <span class="fg0">N/A</span>
                                    {/if}
                                </td>
                                {#if row.hours == null}
                                    <td><span class="fg0">N/A</span></td>
                                {:else}
                                    <td>{row.hours}</td>
                                {/if}
                                <td>${
                                    row.earned.toLocaleString(
                                        "en-US", {
                                        minimumFractionDigits: 2,
                                        maximumFractionDigits: 2
                                    })
                                }</td>
                            </tr>
                        {/each}
                        <tr>
                            <th>Total</th>
                            <th></th>
                            <th></th>
                            <th>${
                                totalKitchenSpecialEarned.toLocaleString(
                                    "en-US", {
                                    minimumFractionDigits: 2,
                                    maximumFractionDigits: 2
                                })
                            }</th>
                        </tr>
                    </tbody>
                </table>
            </div>
            </div>
        {/if}
        {#if dayResults.kitchenWeeklyEmployees.length >= 1}
            <div>
            <div style="display: inline-block; margin-bottom: 2rem;">
                <div class="flex" style="justify-content: space-between;">
                    <p>Floor Workers <span class="fg0">(Weekly Pay)</span></p>
                    <p class="fg0">
                        {#if dayResults.kitchenWeeklyEmployees.length == 1}
                            1 employee
                        {:else}
                            {dayResults.kitchenWeeklyEmployees.length} employees
                        {/if}
                    </p>
                </div>
                <table style="min-width: 20rem;">
                    <thead>
                        <tr>
                            <th>Employee</th>
                            <th>Amount Earned</th>
                        </tr>
                    </thead>
                    <tbody>
                        {#each dayResults.kitchenWeeklyEmployees as employee}
                            <tr>
                                <td>{employee.name}</td>
                                <td>${
                                    employee.weeklyPay.toLocaleString(
                                        "en-US", {
                                        minimumFractionDigits: 2,
                                        maximumFractionDigits: 2
                                    })
                                }</td>
                            </tr>
                        {/each}
                        <tr>
                            <th>Total</th>
                            <th>${
                                dayResults.totalKitchenWeeklyEarned.toLocaleString(
                                    "en-US", {
                                    minimumFractionDigits: 2,
                                    maximumFractionDigits: 2
                                })
                            }</th>
                        </tr>
                    </tbody>
                </table>
            </div>
            </div>
        {/if}
        <div>
        <div style="display: inline-block; margin-bottom: 2rem;">
            <div class="flex" style="justify-content: space-between;">
                <p>Floor Workers</p>
                <p class="fg0">
                    {floorHourlyArray.length} employees
                </p>
            </div>
            <table>
                <thead>
                    <tr>
                        <th>Employee</th>
                        <th>Hours Worked</th>
                        <th>Amount Earned</th>
                    </tr>
                </thead>
                <tbody>
                    {#each floorHourlyArray as row}
                        <tr>
                            <td>{row.employee.name}</td>
                            <td>{row.hours}</td>
                            <td>${
                                (row.hours * row.employee.wage).toLocaleString(
                                    "en-US", {
                                    minimumFractionDigits: 2,
                                    maximumFractionDigits: 2
                                })
                            }</td>
                        </tr>
                    {/each}
                    <tr>
                        <th>Total</th>
                        <th></th>
                        <th>${
                            totalFloorHourlyEarned.toLocaleString(
                                "en-US", {
                                minimumFractionDigits: 2,
                                maximumFractionDigits: 2
                            })
                        }</th>
                    </tr>
                </tbody>
            </table>
        </div>
        </div>
        {#if floorSpecialArray.length >= 1}
            <div>
            <div style="display: inline-block; margin-bottom: 2rem;">
                <div class="flex" style="justify-content: space-between;">
                    <p>Floor Workers <span class="fg0">(Special Pay)</span></p>
                    <p class="fg0">
                        {#if floorSpecialArray.length == 1}
                            1 employee
                        {:else}
                            {floorSpecialArray.length} employees
                        {/if}
                    </p>
                </div>
                <table>
                    <thead>
                        <tr>
                            <th>Employee</th>
                            <th>Worked Today</th>
                            <th>Hours Worked</th>
                            <th>Amount Earned</th>
                        </tr>
                    </thead>
                    <tbody>
                        {#each floorSpecialArray as row}
                            <tr>
                                <td>{row.employee.name}</td>
                                <td>
                                    {#if row.workedToday === true}
                                        <div class="flex compact-gap text yay" style="align-items: center; align-content: center;">
                                            <CheckmarkIcon></CheckmarkIcon>
                                            <span>Yes</span>
                                        </div>
                                    {:else if row.workedToday === false}
                                        <div class="flex compact-gap text ohno" style="align-items: center; align-content: center;">
                                            <XMarkIcon></XMarkIcon>
                                            <span>No</span>
                                        </div>
                                    {:else if row.workedToday == null}
                                        <span class="fg0">N/A</span>
                                    {/if}
                                </td>
                                {#if row.hours == null}
                                    <td><span class="fg0">N/A</span></td>
                                {:else}
                                    <td>{row.hours}</td>
                                {/if}
                                <td>${
                                    row.earned.toLocaleString(
                                        "en-US", {
                                        minimumFractionDigits: 2,
                                        maximumFractionDigits: 2
                                    })
                                }</td>
                            </tr>
                        {/each}
                        <tr>
                            <th>Total</th>
                            <th></th>
                            <th></th>
                            <th>${
                                totalFloorSpecialEarned.toLocaleString(
                                    "en-US", {
                                    minimumFractionDigits: 2,
                                    maximumFractionDigits: 2
                                })
                            }</th>
                        </tr>
                    </tbody>
                </table>
            </div>
            </div>
        {/if}
        {#if dayResults.floorWeeklyEmployees.length >= 1}
            <div>
            <div style="display: inline-block; margin-bottom: 2rem;">
                <div class="flex" style="justify-content: space-between;">
                    <p>Floor Workers <span class="fg0">(Weekly Pay)</span></p>
                    <p class="fg0">
                        {#if dayResults.floorWeeklyEmployees.length == 1}
                            1 employee
                        {:else}
                            {dayResults.floorWeeklyEmployees.length} employees
                        {/if}
                    </p>
                </div>
                <table style="min-width: 20rem;">
                    <thead>
                        <tr>
                            <th>Employee</th>
                            <th>Amount Earned</th>
                        </tr>
                    </thead>
                    <tbody>
                        {#each dayResults.floorWeeklyEmployees as employee}
                            <tr>
                                <td>{employee.name}</td>
                                <td>${
                                    employee.weeklyPay.toLocaleString(
                                        "en-US", {
                                        minimumFractionDigits: 2,
                                        maximumFractionDigits: 2
                                    })
                                }</td>
                            </tr>
                        {/each}
                        <tr>
                            <th>Total</th>
                            <th>${
                                dayResults.totalFloorWeeklyEarned.toLocaleString(
                                    "en-US", {
                                    minimumFractionDigits: 2,
                                    maximumFractionDigits: 2
                                })
                            }</th>
                        </tr>
                    </tbody>
                </table>
            </div>
            </div>
        {/if}
    </div>
</div>
