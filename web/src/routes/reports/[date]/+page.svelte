<script>
import { base } from "$app/paths";
import BackArrowIcon from "$lib/icons/BackArrow.svelte";
import PlusIcon from "$lib/icons/Plus.svelte";
import CheckmarkIcon from "$lib/icons/Checkmark.svelte";
import XMarkIcon from "$lib/icons/CloseXMark.svelte";
import PencilIcon from "$lib/icons/Pencil.svelte";
import DocIcon from "$lib/icons/DocFilePage.svelte";
import CopyIcon from "$lib/icons/Copy.svelte";
import LinkIcon from "$lib/icons/Link.svelte";
import { calculateDay, remainingBudgetFromDays } from "$lib/budget.js";
import { dateToYMDString } from "$lib/dateToYMDString.js";
import { dateGetWeekNum } from "$lib/dateGetWeekNum.js";
import { onMount } from "svelte";
import { plainTextTableFrom2DArray } from "$lib/plainTextTableFrom2DArray.js"
import { tableTo2DArray } from "$lib/tableTo2DArray.js"
import { fade } from "svelte/transition"
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
const endDate = (() => {
    const [year, month, day] = data.week.endDate.split("-");
    return new Date(year, month - 1, day);
})();
const dateYMD = dateToYMDString(startDate);
const weekDayDates = {
    mon: startDate,
    tue: new Date(
        startDate.getFullYear(),
        startDate.getMonth(),
        startDate.getDate() + 1
    ),
    wed: new Date(
        startDate.getFullYear(),
        startDate.getMonth(),
        startDate.getDate() + 2
    ),
    thu: new Date(
        startDate.getFullYear(),
        startDate.getMonth(),
        startDate.getDate() + 3
    ),
    fri: new Date(
        startDate.getFullYear(),
        startDate.getMonth(),
        startDate.getDate() + 4
    ),
    sat: new Date(
        startDate.getFullYear(),
        startDate.getMonth(),
        startDate.getDate() + 5
    ),
    sun: endDate
};
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

let dayResults = $state({});
let foodBudget = $state(data.week.startFoodBudget);
let kitchenBudget = $state(data.week.startKitchenBudget);
let floorBudget = $state(data.week.startFloorBudget);
let dayNum = 1;
data.days.forEach((day) => {
    const [y, m, d] = day.date.split("-");
    const date = new Date(y, m - 1, d);
    const result = calculateDay(
        foodBudget,
        kitchenBudget,
        floorBudget,
        day,
        dayNum == 1 ?
            data.week.weeklyPay : null
    );
    dayResults[weekDayKeys[date.getDay()]] = result;
    foodBudget = result.foodBudgetFinal;
    kitchenBudget = result.kitchenBudgetFinal;
    floorBudget = result.floorBudgetFinal;
})

const shareLink = `${window.location.origin+base}/reports/${dateYMD}?s=${
    data.week.shareLinkToken
}`;

let tableElement;
let copyYay = $state(false);
let copyYayTimeout;
let copyLinkYay = $state(false);
let copyLinkYayTimeout;

let showNoData = $state(false);
let noDataWeekDayKey = $state(-1);
</script>
<div class="grid page" style="margin-top: 4rem; margin-bottom: 10rem;">
    <div class="content">
        <div class="flex">
            <a class="button faint" href="{base}/reports">
                <BackArrowIcon></BackArrowIcon>
                Back
            </a>
        </div>
        <h3 style="margin-bottom: 0px;">Week Report</h3>
        <p style="margin-top: 0.4rem;">
            {#if localStorage?.getItem("budgeter:showWeekNums") == "true"}
                Week {weekNum} ({monthName} {startDate.getDate()}, {startDate.getFullYear()})
            {:else}
                Week of {monthName} {startDate.getDate()}, {startDate.getFullYear()}
            {/if}
        </p>
        <div style="margin-top: 2rem;">
            <div style="display: inline-block; margin-bottom: 2rem;">
                <div style="border: 0.2rem solid var(--border); border-radius: 0.8rem; min-width: 18rem;">
                <table style="border: none;" bind:this={tableElement}>
                    <thead>
                        <tr>
                            <th></th>
                            <th>Food Budget</th>
                            <th>Kitchen Budget</th>
                            <th>Floor Budget</th>
                        </tr>
                    </thead>
                    <tbody>
                        {#each Object.keys(dayResults) as dayKey}
                            <tr>
                                <td><b>{{
                                    "mon": "Monday",
                                    "tue": "Tuesday",
                                    "wed": "Wednesday",
                                    "thu": "Thursday",
                                    "fri": "Friday",
                                    "sat": "Saturday",
                                    "sun": "Sunday"
                                }[dayKey]}</b></td>
                                <td>${dayResults[dayKey].foodBudgetStart.toLocaleString(
                                    "en-US", {
                                    minimumFractionDigits: 2,
                                    maximumFractionDigits: 2
                                })}</td>
                                <td>${dayResults[dayKey].kitchenBudgetStart.toLocaleString(
                                    "en-US", {
                                    minimumFractionDigits: 2,
                                    maximumFractionDigits: 2
                                })}</td>
                                <td>${dayResults[dayKey].floorBudgetStart.toLocaleString(
                                    "en-US", {
                                    minimumFractionDigits: 2,
                                    maximumFractionDigits: 2
                                })}</td>
                            </tr>
                            <tr>
                                <td>Increases</td>
                                <td>+${dayResults[dayKey].foodBudgetIncrease.toLocaleString(
                                    "en-US", {
                                    minimumFractionDigits: 2,
                                    maximumFractionDigits: 2
                                })}</td>
                                <td>+${dayResults[dayKey].kitchenBudgetIncrease.toLocaleString(
                                    "en-US", {
                                    minimumFractionDigits: 2,
                                    maximumFractionDigits: 2
                                })}</td>
                                <td>+${dayResults[dayKey].floorBudgetIncrease.toLocaleString(
                                    "en-US", {
                                    minimumFractionDigits: 2,
                                    maximumFractionDigits: 2
                                })}</td>
                            </tr>
                            <tr>
                                    <td>Expenses</td>
                                <td>-${dayResults[dayKey].foodExpenses.toLocaleString(
                                    "en-US", {
                                    minimumFractionDigits: 2,
                                    maximumFractionDigits: 2
                                })}</td>
                                <td>-${
                                    dayResults[dayKey].kitchenExpenses.toLocaleString(
                                        "en-US", {
                                        minimumFractionDigits: 2,
                                        maximumFractionDigits: 2
                                    })
                                }</td>
                                <td>-${
                                    dayResults[dayKey].floorExpenses.toLocaleString(
                                        "en-US", {
                                        minimumFractionDigits: 2,
                                        maximumFractionDigits: 2
                                    })
                                }</td>
                            </tr>
                        {/each}
                        <tr>
                            <td><b>Week Final</b></td>
                            <td>${foodBudget.toLocaleString(
                                "en-US", {
                                minimumFractionDigits: 2,
                                maximumFractionDigits: 2
                            })}</td>
                            <td>${kitchenBudget.toLocaleString(
                                "en-US", {
                                minimumFractionDigits: 2,
                                maximumFractionDigits: 2
                            })}</td>
                            <td>${floorBudget.toLocaleString(
                                "en-US", {
                                minimumFractionDigits: 2,
                                maximumFractionDigits: 2
                            })}</td>
                        </tr>
                    </tbody>
                </table>
                </div>
                {#if !data.usingShareLink}
                <details>
                    <summary style="cursor: pointer;">Copy, Save, or Send this Report</summary>
                    <div>
                        <p>You can copy &amp; paste the table</p>
                        <button class="alt" onclick={async () => {
                            try {
                                clearTimeout(copyYayTimeout)

                                let table = tableElement.cloneNode(true);
                                table.removeAttribute("style");

                                const plainTextTable = plainTextTableFrom2DArray(
                                    tableTo2DArray(tableElement),
                                    2
                                );

                                await navigator.clipboard.write([
                                  new ClipboardItem({
                                    'text/html': new Blob([table.outerHTML], { type: 'text/html' }),
                                    'text/plain': new Blob([plainTextTable], { type: 'text/plain' })
                                  })
                                ]);
                                copyYay = true;
                                copyYayTimeout = setTimeout(
                                    () => copyYay = false, 1000
                                );
                            } catch (error) {
                                console.error(error)
                                alert("idk why it didn't work 💀")
                            }
                        }}>
                            {#if copyYay}
                                <CheckmarkIcon></CheckmarkIcon>
                            {:else}
                                <CopyIcon></CopyIcon>
                            {/if}
                            Copy
                        </button>
                        <p style="margin-top: 2rem;">Or you can send a link to this page;<br>
                        anyone with this special link can view this report<br>
                        without having to log in<br></p>
                        <div class="flex compact-gap">
                            <input type="text" value={shareLink} disabled>
                            <button class="alt" onclick={async () => {
                                try {
                                    clearTimeout(copyLinkYayTimeout)
                                    
                                    await navigator.clipboard.writeText(shareLink)
                                    copyLinkYay = true;
                                    copyLinkYayTimeout = setTimeout(
                                        () => copyLinkYay = false, 1000
                                    );
                                } catch (error) {
                                    console.error(error)
                                    alert("idk why it didn't work 💀")
                                }
                            }}>
                                {#if copyLinkYay}
                                    <CheckmarkIcon></CheckmarkIcon>
                                {:else}
                                    <LinkIcon></LinkIcon>
                                {/if}
                                Copy Link
                            </button>
                        </div>
                    </div>
                </details>
                {/if}
            </div>
        </div>
        <p>
            {#if data.usingShareLink}
                Day Reports & Details:
            {:else}
                View/Edit Days:
            {/if}
        </p>
        <div class="flex" style="max-width: 16rem; flex-direction: column; flex-wrap: nowrap;">
            {#each ["mon","tue","wed","thu","fri","sat","sun"] as weekDayKey}
                {#if data.usingShareLink && dayResults[weekDayKey] == null}
                <button class="button button-box" onclick={() => {
                    showNoData = true;
                    noDataWeekDayKey = weekDayKey;
                }}>
                    {{
                        mon: "Monday",
                        tue: "Tuesday",
                        wed: "Wednesday",
                        thu: "Thursday",
                        fri: "Friday",
                        sat: "Saturday",
                        sun: "Sunday"
                    }[weekDayKey]}, <span class="fg0">{
                        [
                            "Jan", "Feb", "Mar", "Apr",
                            "May", "Jun", "Jul", "Aug",
                            "Sep", "Oct", "Nov", "Dec"
                        ][weekDayDates[weekDayKey].getMonth()]
                    } {weekDayDates[weekDayKey].getDate()}</span>
                </button>
                {:else}
                <a class="button button-box" href={
                    data.usingShareLink ?
                        `${base}/days/${dateToYMDString(
                            weekDayDates[weekDayKey]
                        )}?s=${data.week.shareLinkToken}` : (
                            dayResults[weekDayKey] ?
                                `${base}/days/${dateToYMDString(
                                    weekDayDates[weekDayKey]
                                )}` :
                                `${base}/record-day?d=${dateToYMDString(
                                    weekDayDates[weekDayKey]
                                )}`
                        )
                }>
                    {{
                        mon: "Monday",
                        tue: "Tuesday",
                        wed: "Wednesday",
                        thu: "Thursday",
                        fri: "Friday",
                        sat: "Saturday",
                        sun: "Sunday"
                    }[weekDayKey]}, <span class="fg0">{
                        [
                            "Jan", "Feb", "Mar", "Apr",
                            "May", "Jun", "Jul", "Aug",
                            "Sep", "Oct", "Nov", "Dec"
                        ][weekDayDates[weekDayKey].getMonth()]
                    } {weekDayDates[weekDayKey].getDate()}</span>
                </a>
                {/if}
            {/each}
        </div>
    </div>
</div>
{#if showNoData}
    <div class="modal" transition:fade={{ duration: 100 }}>
        <div class="content">
            <h4>
                {{
                    mon: "Monday",
                    tue: "Tuesday",
                    wed: "Wednesday",
                    thu: "Thursday",
                    fri: "Friday",
                    sat: "Saturday",
                    sun: "Sunday"
                }[noDataWeekDayKey]}, {
                    [
                        "Jan", "Feb", "Mar", "Apr",
                        "May", "Jun", "Jul", "Aug",
                        "Sep", "Oct", "Nov", "Dec"
                    ][weekDayDates[noDataWeekDayKey].getMonth()]
                } {weekDayDates[noDataWeekDayKey].getDate()}
            </h4>
            <p>No data for this day</p>
            <div class="flex">
            <button onclick={() => {
                showNoData = false;
                noDataWeekDayKey = -1;
            }}>
                Okay
            </button>
            </div>
        </div>
    </div>
{/if}
