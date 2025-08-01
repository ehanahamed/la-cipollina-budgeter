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
const startDate = (() => {
    const [year, month, day] = data.week.startDate.split("-");
    const d = new Date(year, month - 1, day);
    if (!isNaN(d)) {
        return d;
    } else {
        window.location = `${base}/reports/${data.week.startDate}/err/invalid-date`
    }
})();
const dateYMD = dateToYMDString(startDate);
const weekDayKeys = [
    "sun", "mon", "tue", "wed",
    "thu", "fri", "sat"
];
const monthName = [
    "January", "February", "March", "April", "May",
    "June", "July", "August", "September",
    "October", "November", "December"
][startDate.getMonth()];
const weekNum = dateGetWeekNum(startDate);
let prevDaysArray = [];
let selectedDay;

let dayResults = $state({});
let foodBudget = $state(data.week.startFoodBudget);
let kitchenBudget = $state(data.week.startKitchenBudget);
let floorBudget = $state(data.week.startFloorBudget);
data.days.forEach((day) => {
    const [y, m, d] = day.date.split("-");
    const date = new Date(y, m - 1, d);
    dayResults[weekDayKeys[date.getDay()]] = calculateDay(
        foodBudget,
        kitchenBudget,
        floorBudget,
        day
    );
})
console.log(dayResults);
</script>
<div class="grid page" style="margin-top: 4rem; margin-bottom: 10rem;">
    <div class="content">
        <div class="flex">
            <a class="button faint" href={base}>
                <BackArrowIcon></BackArrowIcon>
                Back
            </a>
        </div>
        <h3 style="margin-bottom: 0px;">Week Report</h3>
        <p style="margin-top: 0.4rem;">
            {#if localStorage.getItem("budgeter:showWeekNums") == "true"}
                Week of {monthName} {startDate.getDate()}, {startDate.getFullYear()}
                <!-- `, W${weekNum}` : "" -->
            {/if}
        </p>
        <div style="margin-top: 2rem;">
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
                        {#if dayResults?.wed}
                        <tr>
                            <td><b>Monday</b></td>
                            <td>${dayResults.wed.foodBudgetStart.toLocaleString(
                                "en-US", {
                                minimumFractionDigits: 2,
                                maximumFractionDigits: 2
                            })}</td>
                            <td>${dayResults.wed.kitchenBudgetStart.toLocaleString(
                                "en-US", {
                                minimumFractionDigits: 2,
                                maximumFractionDigits: 2
                            })}</td>
                            <td>${dayResults.wed.floorBudgetStart.toLocaleString(
                                "en-US", {
                                minimumFractionDigits: 2,
                                maximumFractionDigits: 2
                            })}</td>
                        </tr>
                        <tr>
                            <td>Increases</td>
                            <td>+${dayResults.wed.foodBudgetIncrease.toLocaleString(
                                "en-US", {
                                minimumFractionDigits: 2,
                                maximumFractionDigits: 2
                            })}</td>
                            <td>+${dayResults.wed.kitchenBudgetIncrease.toLocaleString(
                                "en-US", {
                                minimumFractionDigits: 2,
                                maximumFractionDigits: 2
                            })}</td>
                            <td>+${dayResults.wed.floorBudgetIncrease.toLocaleString(
                                "en-US", {
                                minimumFractionDigits: 2,
                                maximumFractionDigits: 2
                            })}</td>
                        </tr>
                        <tr>
                            <td>Expenses</td>
                            <td>-${dayResults.wed.foodExpenses.toLocaleString(
                                "en-US", {
                                minimumFractionDigits: 2,
                                maximumFractionDigits: 2
                            })}</td>
                            <td>-${
                                dayResults.wed.kitchenExpenses.toLocaleString(
                                    "en-US", {
                                    minimumFractionDigits: 2,
                                    maximumFractionDigits: 2
                                })
                            }</td>
                            <td>-${
                                dayResults.wed.floorExpenses.toLocaleString(
                                    "en-US", {
                                    minimumFractionDigits: 2,
                                    maximumFractionDigits: 2
                                })
                            }</td>
                        </tr>
                        {/if}
                        <!-- <tr> -->
                        <!--     <td>Day Final</td> -->
                        <!--     <td>${dayFinalFoodBudget.toLocaleString( -->
                        <!--         "en-US", { -->
                        <!--         minimumFractionDigits: 2, -->
                        <!--         maximumFractionDigits: 2 -->
                        <!--     })}</td> -->
                        <!--     <td>${dayFinalKitchenBudget.toLocaleString( -->
                        <!--         "en-US", { -->
                        <!--         minimumFractionDigits: 2, -->
                        <!--         maximumFractionDigits: 2 -->
                        <!--     })}</td> -->
                        <!--     <td>${dayFinalFloorBudget.toLocaleString( -->
                        <!--         "en-US", { -->
                        <!--         minimumFractionDigits: 2, -->
                        <!--         maximumFractionDigits: 2 -->
                        <!--     })}</td> -->
                        <!-- </tr> -->
                    </tbody>
                </table>
                </div>
            </div>
        </div>
    </div>
</div>
