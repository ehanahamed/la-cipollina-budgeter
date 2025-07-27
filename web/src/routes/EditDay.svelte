<script>
import { base } from "$app/paths";
import BackArrowIcon from "$lib/icons/BackArrow.svelte";
import PlusIcon from "$lib/icons/Plus.svelte";
import CheckmarkIcon from "$lib/icons/Checkmark.svelte";
import XMarkIcon from "$lib/icons/CloseXMark.svelte";
import { onMount } from "svelte";
let { data } = $props();
let floorHoursArray = $state([]);
let floorWorkedTodayArray = $state([]);
let kitchenHoursArray = $state([]);
let kitchenWorkedTodayArray = $state([]);
let floorEmployeesCount = $state(0);
let kitchenEmployeesCount = $state(0);
const date = (() => {
    if (data.date) {
        const [year, month, day] = data.date.split("-");
        const d = new Date(year, month - 1, day);
        if (!isNaN(d)) {
            return d;
        }
    }
    return new Date()
})();
const weekDayName = [
    "Sunday", "Monday", "Tuesday", "Wednesday",
    "Thrusday", "Friday", "Saturday"
][date.getDay()]
const weekDayKey = [
    "mon", "tue", "wed", "thu",
    "fri", "sat", "sun"
][date.getDay()]
if (data.new) {
    data.employees.forEach(function (employee) {
        if (employee?.type == "FLOOR") {
            floorEmployeesCount++;
            if (employee?.specialPay) {
                if (employee.specialPay[weekDayKey]?.perDay != null) {
                    floorWorkedTodayArray.push({
                        employee: employee,
                        workedToday: false
                    });
                }
                if (employee.specialPay[weekDayKey]?.perHour != null) {
                    floorHoursArray.push({
                        employee: employee,
                        hours: ""
                    });
                }
            } else {
                floorHoursArray.push({
                    employee: employee,
                    hours: ""
                });
            }
        } else if (employee?.type == "KITCHEN") {
            kitchenEmployeesCount++;
            if (employee?.specialPay) {
                if (employee.specialPay[weekDayKey]?.perDay != null) {
                    kitchenWorkedTodayArray.push({
                        employee: employee,
                        workedToday: false
                    });
                }
                if (employee.specialPay[weekDayKey]?.perHour != null) {
                    kitchenHoursArray.push({
                        employee: employee,
                        hours: ""
                    });
                }
            } else {
                kitchenHoursArray.push({
                    employee: employee,
                    hours: ""
                });
            }
        }
    })
}
let foodCostsArray = $state([
    {
        cost: "",
        notes: ""
    }
]);
let foodCostsTotal = $derived.by(() => {
    let total = 0;
    foodCostsArray.forEach((row) => {
        if (typeof row.cost === "string") {
            const cost = parseFloat(row.cost.replaceAll(",", ""));
            if (!isNaN(cost)) {
                total += cost;
            }
        } else if (typeof row.cost === "number") {
            total += row.cost;
        }
    })
    return total;
})
</script>
<style>
.gridtablelayout {
    display: grid;
    gap: 1rem;
    grid-template-rows: auto auto auto;
    grid-template-columns: auto auto auto;
    grid-template-areas:
        "areaback areaback areaback"
        "areafloor areakitchen areafood"
        "areasave areasave areasave";
}
.gridtablelayout > .areaback {
    grid-area: areaback;
    justify-self: start;
}
.gridtablelayout > .areainfo {
    grid-area: areainfo;
    justify-self: center;
}
.gridtablelayout > .areafloor {
    grid-area: areafloor;
    justify-self: center;
}
.gridtablelayout > .areakitchen {
    grid-area: areakitchen;
    justify-self: center;
}
.gridtablelayout > .areafood {
    grid-area: areafood;
    justify-self: center;
}
.gridtablelayout > .areasave {
    grid-area: areasave;
    justify-self: start;
}
.centerbutnotonmobile {
    text-align: center;
}
.gridtablelayout table td {
    padding-top: 0px;
}
@media only screen and (max-width: 1000px) {
    .gridtablelayout {
        grid-template-rows: auto auto auto auto auto;
        grid-template-columns: auto;
        grid-template-areas:
            "areaback"
            "areainfo"
            "areafloor"
            "areakitchen"
            "areafood"
            "areasave";
    }
    .gridtablelayout > .areainfo,
    .gridtablelayout > .areafloor,
    .gridtablelayout > .areakitchen,
    .gridtablelayout > .areafood {
        justify-self: start;
    }
    .centerbutnotonmobile {
        text-align: start;
    }
}
.setbudgettable {
    
}
</style>
<div class="grid page" style="margin-top: 4rem; margin-bottom: 10rem;">
    <div class="content">
        <div class="areainfo">
            <div class="centerbutnotonmobile">
                {#if data.new}
                    <p>Run Day</p>
                {:else}
                    <p>Edit Day</p>
                {/if}
                <h3 style="margin-top: 0.2rem; margin-bottom: 0px;">{weekDayName}</h3>
                <p style="margin-top: 0.2rem;">{[
                    "January", "February", "March", "April",
                    "May", "June", "July", "August", "September",
                    "October", "November", "December"
                ][date.getMonth()]} {date.getDate()}, {date.getFullYear()}</p>
            </div>
        </div>
        <div class="gridtablelayout">
            <div class="areaback">
                <a href={base} class="button faint">
                    <BackArrowIcon></BackArrowIcon>
                    Back
                </a>
            </div>
            <div class="areafloor">
                <div class="flex" style="justify-content: space-between;">
                    <p>Floor Workers</p>
                    <p class="fg0">
                        {floorHoursArray.length + floorWorkedTodayArray.length} total
                    </p>
                </div>
                <table>
                    <thead>
                        <tr>
                            <th>Employee</th>
                            <th>Hours Worked</th>
                        </tr>
                    </thead>
                    <tbody>
                        {#each floorHoursArray as row}
                            <tr>
                                <td>{row.employee.name}</td>
                                <td><input type="text" bind:value={row.hours}
                                    style="min-width: 6rem; field-sizing: content;"></td>
                            </tr>
                        {/each}
                        {#if floorWorkedTodayArray.length >= 1}
                        <tr>
                            <th>Employee</th>
                            <th>Worked Today?</th>
                        </tr>
                        {/if}
                        {#each floorWorkedTodayArray as row}
                            <tr>
                                <td>{row.employee.name}</td>
                                <td>
                                    <div class="combo-select">
                                        <button class="left {
                                            row.workedToday ?
                                                "selected" : ""
                                        }" onclick={() => row.workedToday = true}>
                                            <CheckmarkIcon class="combo-selected-icon"></CheckmarkIcon>
                                            Yes
                                        </button>
                                        <button class="right {
                                            row.workedToday ?
                                                "" : "selected"
                                        }" onclick={() => row.workedToday = false}>
                                            <CheckmarkIcon class="combo-selected-icon"></CheckmarkIcon>
                                            No
                                        </button>
                                    </div>
                                </td>
                            </tr>
                        {/each}
                    </tbody>
                </table>
            </div>
            <div class="areakitchen">
                <div class="flex" style="justify-content: space-between;">
                    <p>Kitchen Workers</p>
                    <p class="fg0">
                        {kitchenHoursArray.length + kitchenWorkedTodayArray.length} total
                    </p>
                </div>
                <table>
                    <thead>
                        <tr>
                            <th>Employee</th>
                            <th>Hours Worked</th>
                        </tr>
                    </thead>
                    <tbody>
                        {#each kitchenHoursArray as row}
                            <tr>
                                <td>{row.employee.name}</td>
                                <td><input type="text" bind:value={row.hours}
                                    style="min-width: 6rem; field-sizing: content;"></td>
                            </tr>
                        {/each}
                        {#if kitchenWorkedTodayArray.length >= 1}
                        <tr>
                            <th>Employee</th>
                            <th>Worked Today?</th>
                        </tr>
                        {/if}
                        {#each kitchenWorkedTodayArray as row}
                            <tr>
                                <td>{row.employee.name}</td>
                                <td>
                                    <div class="combo-select">
                                        <button class="left {
                                            row.workedToday ?
                                                "selected" : ""
                                        }" onclick={() => row.workedToday = true}>
                                            <CheckmarkIcon class="combo-selected-icon"></CheckmarkIcon>
                                            Yes
                                        </button>
                                        <button class="right {
                                            row.workedToday ?
                                                "" : "selected"
                                        }" onclick={() => row.workedToday = false}>
                                            <CheckmarkIcon class="combo-selected-icon"></CheckmarkIcon>
                                            No
                                        </button>
                                    </div>
                                </td>
                            </tr>
                        {/each}
                    </tbody>
                </table>
            </div>
            <div class="areafood">
                <div class="flex" style="justify-content: space-between;">
                    <p>Food Cost</p>
                    <p class="fg0">
                        {foodCostsArray.length} {foodCostsArray.length == 1 ?
                            "item" : "items"}
                    </p>
                </div>
                <div style="border: 0.2rem solid var(--border); border-radius: 0.8rem;">
                <table style="border: none;">
                    <thead>
                        <tr>
                            <th>Expenses</th>
                            <th>Notes</th>
                        </tr>
                    </thead>
                    <tbody>
                        {#each foodCostsArray as row, rowIndex}
                            <tr>
                                <td style="align-content: start; padding-right: 0px;">
                                    <div class="flex compact-gap nowrap" style="align-items: center; align-content: center;">
                                        <span>$</span>
                                        <input type="text" placeholder="0.00" bind:value={row.cost} style="min-width: 5rem; field-sizing: content;">
                                    </div>
                                </td>
                                <td><textarea style="width: 10rem;" placeholder="Optional note for each expense" bind:value={row.notes}></textarea></td>
                                    <td style="align-content: start; padding-left: 0px;">
                                    <button class="icon-only-button" onclick={() => {
                                        foodCostsArray.splice(rowIndex, 1)
                                    }}>
                                        <XMarkIcon></XMarkIcon>
                                    </button>
                                </td>
                            </tr>
                        {/each}
                        <tr>
                            <td style="padding-right: 0px;">
                                <button onclick={function () {
                                    foodCostsArray.push({cost:"",notes:""})
                                }}>
                                    <PlusIcon></PlusIcon>
                                    Add
                                </button>
                            </td>
                        </tr>
                    </tbody>
                </table>
                <div style="padding: 0px 0px 1rem 1rem; margin-top: 0.6rem; font-size: 1.1rem;"
                >Total: ${
                    foodCostsTotal.toLocaleString('en-US', {
                        minimumFractionDigits: 2,
                        maximumFractionDigits: 2
                    })
                }</div>
                </div>
            </div>
            <div class="areasave">
            </div>
        </div>
    </div>
</div>
