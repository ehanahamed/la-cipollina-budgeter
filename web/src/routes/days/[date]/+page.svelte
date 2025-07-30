<script>
import { base } from "$app/paths";
import BackArrowIcon from "$lib/icons/BackArrow.svelte";
import PlusIcon from "$lib/icons/Plus.svelte";
import CheckmarkIcon from "$lib/icons/Checkmark.svelte";
import XMarkIcon from "$lib/icons/CloseXMark.svelte";
import { onMount } from "svelte";
let { data } = $props();
let floorHourlyArray = $state([]);
let floorSpecialArray = $state([]);
let kitchenHourlyArray = $state([]);
let kitchenSpecialArray = $state([]);
const date = (() => {
    const [year, month, day] = data.date.split("-");
    const d = new Date(year, month - 1, day);
    if (!isNaN(d)) {
        return d;
    } else {
        window.location = `${base}/days/${data.date}/err/invalid-date`
    }
})();
function dateToYMDString(dateObj) {
    return dateObj.getFullYear() + "-" +
        (dateObj.getMonth() + 1).toString().padStart(2, "0") + "-" +
        dateObj.getDate().toString().padStart(2, "0")
}
const dateYMD = dateToYMDString(date);
const weekDayName = [
    "Sunday", "Monday", "Tuesday", "Wednesday",
    "Thrusday", "Friday", "Saturday"
][date.getDay()];
const weekDayKey = [
    "mon", "tue", "wed", "thu",
    "fri", "sat", "sun"
][date.getDay()];
const monthName = [
    "January", "February", "March", "April", "May",
    "June", "July", "August", "September",
    "October", "November", "December"
][date.getDay()];
let weekData = $state(null);
let showWeekBudget = $state(false);
let startingNewWeek = $state(false);
let prevDaysArray = $state([]);
let selectedDay = $state(null);
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
let foodCostsTotal = $state(0);
selectedDay.foodCosts.forEach((row) => {
    foodCostsTotal += row.cost;
})
selectedDay.hoursWorked.forEach((row) => {
    if (row.employee.specialPay) {
        if (row.employee.type == "FLOOR") {
            floorSpecialArray.push(row);
        } else if (row.employee.type == "KITCHEN") {
            kitchenSpecialArray.push(row);
        }
    }
    if (row.employee.type == "FLOOR") {
        floorHourlyArray.push(row);
    } else if (row.employee.type == "KITCHEN") {
        kitchenHourlyArray.push(row);
    }
})
selectedDay.workedToday.forEach((row) => {
    if (row.employee.type == "FLOOR") {
        let found = false;
        for (
            let index = 0;
            index < floorSpecialArray.length;
            index++
        ) {
            if (
                floorSpecialArray[index].employee.id ==
                row.employee.id
            ) {
                found = true;
                floorSpecialArray[
                    index
                ].workedToday = row.workedToday;
            }
        }
        if (!found) {
            floorSpecialArray.push(row);
        }
    } else if (row.employee.type == "KITCHEN") {
        let found = false;
        for (
            let index = 0;
            index < kitchenSpecialArray.length;
            index++
        ) {
            if (
                kitchenSpecialArray[index].employee.id ==
                row.employee.id
            ) {
                found = true;
                kitchenSpecialArray[
                    index
                ].workedToday = row.workedToday;
            }
        }
        if (!found) {
            kitchenSpecialArray.push(row);
        }
    }
})
let totalKitchenHourlyEarned = $state(0);
let totalFloorHourlyEarned = $state(0);
let totalKitchenSpecialEarned = $state(0);
let totalFloorSpecialEarned = $state(0);
kitchenHourlyArray.forEach((row) => {
    totalKitchenHourlyEarned += (
        row.hours * row.employee.wage
    );
})
floorHourlyArray.forEach((row) => {
    totalFloorHourlyEarned += (
        row.hours * row.employee.wage
    );
})
for (
    let index = 0;
    index < kitchenSpecialArray.length;
    index++
) {
    const row = kitchenSpecialArray[index];
    let earned = 0;
    if (
        row.workedToday &&
        row?.employee?.specialPay?.[
            weekDayKey
        ]?.perDay != null
    ) {
        earned += row.employee.specialPay[
            weekDayKey
        ].perDay;
    }
    if (
        row.hours != null &&
        row?.employee?.specialPay?.[
            weekDayKey
        ]?.perHour != null
    ) {
        earned += hours * row.employee.specialPay[
            weekDayKey
        ].perDay;
    }
    kitchenSpecialArray[index].earned = earned;
    totalKitchenSpecialEarned += earned;
}
for (
    let index = 0;
    index < floorSpecialArray.length;
    index++
) {
    const row = floorSpecialArray[index];
    let earned = 0;
    if (
        row.workedToday &&
        row?.employee?.specialPay?.[
            weekDayKey
        ]?.perDay != null
    ) {
        earned += row.employee.specialPay[
            weekDayKey
        ].perDay;
    }
    if (
        row.hours != null &&
        row?.employee?.specialPay?.[
            weekDayKey
        ]?.perHour != null
    ) {
        earned += hours * row.employee.specialPay[
            weekDayKey
        ].perDay;
    }
    floorSpecialArray[index].earned = earned;
    totalFloorSpecialEarned += earned;
}

let dayStartFoodBudget = data.week.startFoodBudget;
let dayStartKitchenBudget = data.week.startKitchenBudget;
let dayStartFloorBudget = data.week.startFloorBudget;
prevDaysArray.forEach((day) => {
    dayStartFoodBudget += day.foodBudgetIncrease;
    dayStartKitchenBudget += day.kitchenBudgetIncrease;
    dayStartFloorBudget += day.floorBudgetIncrease;

    const iDate = (() => {
        const [year, month, day] = data.date.split("-");
        return new Date(year, month - 1, day);
    })();
    const iWeekDayKey = [
        "mon", "tue", "wed", "thu",
        "fri", "sat", "sun"
    ][iDate.getDay()];


    let iFoodCostsTotal = 0;
    let iKitchenHourlyArray = [];
    let iKitchenSpecialArray = [];
    let iFloorHourlyArray = [];
    let iFloorSpecialArray = [];

    day.foodCosts.forEach((row) => {
        iFoodCostsTotal += row.cost;
    })

    day.hoursWorked.forEach((row) => {
        if (row.employee.specialPay) {
            if (row.employee.type == "FLOOR") {
                iFloorSpecialArray.push(row);
            } else if (row.employee.type == "KITCHEN") {
                iKitchenSpecialArray.push(row);
            }
        }
        if (row.employee.type == "FLOOR") {
            iFloorHourlyArray.push(row);
        } else if (row.employee.type == "KITCHEN") {
            iKitchenHourlyArray.push(row);
        }
    })
    day.workedToday.forEach((row) => {
        if (row.employee.type == "FLOOR") {
            let found = false;
            for (
                let index = 0;
                index < iFloorSpecialArray.length;
                index++
            ) {
                if (
                    iFloorSpecialArray[index].employee.id ==
                    row.employee.id
                ) {
                    found = true;
                    iFloorSpecialArray[
                        index
                    ].workedToday = row.workedToday;
                }
            }
            if (!found) {
                iFloorSpecialArray.push(row);
            }
        } else if (row.employee.type == "KITCHEN") {
            let found = false;
            for (
                let index = 0;
                index < iKitchenSpecialArray.length;
                index++
            ) {
                if (
                    iKitchenSpecialArray[index].employee.id ==
                    row.employee.id
                ) {
                    found = true;
                    iKitchenSpecialArray[
                        index
                    ].workedToday = row.workedToday;
                }
            }
            if (!found) {
                iKitchenSpecialArray.push(row);
            }
        }
    })
    let iTotalKitchenEarned = 0;
    let iTotalFloorEarned = 0;
    iKitchenHourlyArray.forEach((row) => {
        iTotalKitchenEarned += (
            row.hours * row.employee.wage
        );
    })
    iFloorHourlyArray.forEach((row) => {
        iTotalFloorEarned += (
            row.hours * row.employee.wage
        );
    })
    for (
        let index = 0;
        index < iKitchenSpecialArray.length;
        index++
    ) {
        const row = iKitchenSpecialArray[index];
        if (
            row.workedToday &&
            row?.employee?.specialPay?.[
                iWeekDayKey
            ]?.perDay != null
        ) {
            iTotalKitchenEarned += row.employee.specialPay[
                iWeekDayKey
            ].perDay;
        }
        if (
            row.hours != null &&
            row?.employee?.specialPay?.[
                iWeekDayKey
            ]?.perHour != null
        ) {
            iTotalKitchenEarned += hours * row.employee.specialPay[
                iWeekDayKey
            ].perDay;
        }
    }
    for (
        let index = 0;
        index < iFloorSpecialArray.length;
        index++
    ) {
        const row = iFloorSpecialArray[index];
        if (
            row.workedToday &&
            row?.employee?.specialPay?.[
                iWeekDayKey
            ]?.perDay != null
        ) {
            iTotalFloorEarned += row.employee.specialPay[
                iWeekDayKey
            ].perDay;
        }
        if (
            row.hours != null &&
            row?.employee?.specialPay?.[
                weekDayKey
            ]?.perHour != null
        ) {
            iTotalFloorEarned += hours * row.employee.specialPay[
                iWeekDayKey
            ].perDay;
        }
    }

    dayStartFoodBudget -= iFoodCostsTotal;
    dayStartKitchenBudget -= iTotalKitchenEarned;
    dayStartFloorBudget -= iTotalFloorEarned;
})
let dayFinalFoodBudget = $derived(
    dayStartFoodBudget +
    (selectedDay.foodBudgetIncrease ?? 0) -
    foodCostsTotal
);
let dayFinalKitchenBudget = $derived(
    dayStartKitchenBudget +
    (selectedDay.kitchenBudgetIncrease ?? 0) - (
        totalKitchenHourlyEarned +
        totalKitchenSpecialEarned
    )
);
let dayFinalFloorBudget = $derived(
    dayStartFloorBudget +
    (selectedDay.floorBudgetIncrease ?? 0) - (
        totalFloorHourlyEarned +
        totalFloorSpecialEarned
    )
)
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
        <p style="margin-top: 0.4rem; margin-bottom: 2rem;">{monthName} {date.getDate()}, {date.getFullYear()}</p>
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
                                (
                                    totalKitchenHourlyEarned +
                                    totalKitchenSpecialEarned
                                ).toLocaleString(
                                    "en-US", {
                                    minimumFractionDigits: 2,
                                    maximumFractionDigits: 2
                                })
                            }</td>
                            <td>-${
                                (
                                    totalFloorHourlyEarned +
                                    totalFloorSpecialEarned
                                ).toLocaleString(
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
                                <td>{row.notes}</td>
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
    </div>
</div>
