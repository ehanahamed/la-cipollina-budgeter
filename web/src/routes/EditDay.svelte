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
const dateYMD = date.getFullYear() + "-" +
    (date.getMonth() + 1).toString().padStart(2, "0") + "-" +
    date.getDate().toString().padStart(2, "0")
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
let weekData = $state(null);
let showWeekBudget = $state(false);
let startingNewWeek = $state(false);
let newWeekFloorBudget = $state("");
let newWeekKitchenBudget = $state("");
let newWeekFoodBudget = $state("");
try {
    fetch(
        data.PUBLIC_API_URL + `/weeks/${dateYMD}`,
        {
            method: "GET",
            headers: {
                "Authorization": "Bearer " + localStorage.getItem("budgeter:auth")
            }
        }
    ).then((res) => {
        res.json().then((resJson) => {
            if (resJson.statusCode == 404) {
                startingNewWeek = true;
                showWeekBudget = true;
            } else if (resJson) {
                weekData = resJson;
                showWeekBudget = true;
            } else {
                console.log(weekResJson);
                alert("something didn't work idk 💀");
            }
        })
    });
} catch (err) {
    console.error(err);
    alert("something went wrong idk 💀");
}
</script>
<style>
.gridtablelayout {
    display: grid;
    gap: 1rem;
    grid-template-rows: auto auto auto;
    grid-template-columns: auto auto auto;
    grid-template-areas:
        "areainfo areainfo areainfo"
        "areaback areaback areaback"
        "areabudget areabudget areabudget"
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
.gridtablelayout > .areabudget {
    grid-area: areabudget;
    justify-self: start;
}
.gridtablelayout > .areafloor {
    grid-area: areafloor;
    justify-self: start;
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
.gridtablelayout table td {
    padding-top: 0px;
}
.areabudget table input[type="text"] {
    min-width: 8rem;
    field-sizing: content;
}
@media only screen and (max-width: 1000px) {
    .gridtablelayout {
        grid-template-rows: auto auto auto auto auto;
        grid-template-columns: auto;
        grid-template-areas:
            "areainfo"
            "areaback"
            "areabudget"
            "areafloor"
            "areakitchen"
            "areafood"
            "areasave";
    }
    .gridtablelayout > .areabudget,
    .gridtablelayout > .areafloor,
    .gridtablelayout > .areakitchen,
    .gridtablelayout > .areafood {
        justify-self: start;
    }
    .areabudget table input[type="text"] {
        min-width: 5rem;
        max-width: 8rem;
        field-sizing: content;
    }
    .areabudget table td+td,
    .areabudget table th+th,
    .areabudget table th+td,
    .areabudget table td+th {
        padding-left: 0px;
    }
}
</style>
<div class="grid page" style="margin-top: 4rem; margin-bottom: 10rem;">
    <div class="content">
        <div class="gridtablelayout">
            <div class="areainfo">
                <div style="text-align: center;">
                    {#if data.new}
                        <p>Record Day</p>
                    {:else}
                        <p>Edit Day</p>
                    {/if}
                    <h3 style="margin-top: 0px; margin-bottom: 0px;">{weekDayName}</h3>
                    <p style="margin-top: 0.2rem;">{[
                        "January", "February", "March", "April",
                        "May", "June", "July", "August", "September",
                        "October", "November", "December"
                    ][date.getMonth()]} {date.getDate()}, {date.getFullYear()}</p>
                </div>
            </div>
            <div class="areaback" style="margin-top: 0px">
                <a href={base} class="button faint">
                    <BackArrowIcon></BackArrowIcon>
                    Back
                </a>
            </div>
            <div class="areabudget" style="margin-top: 0px; margin-bottom: 1rem;">
                {#if showWeekBudget && startingNewWeek}
                    <p>Set this week's starting budget</p>
                    <div style="border: 0.2rem solid var(--border); border-radius: 0.8rem;">
                        <table style="border: none;">
                            <thead>
                                <tr>
                                    <th></th>
                                    <th>Floor Budget</th>
                                    <th>Kitchen Budget</th>
                                    <th>Food Budget</th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr>
                                    <td>Start</td>
                                    <td><div class="flex compact-gap nowrap" style="align-items: center; align-content: center;">
                                        <span>$</span>
                                        <input type="text" placeholder="0.00" bind:value={newWeekFloorBudget}>
                                    </div></td>
                                    <td><div class="flex compact-gap nowrap" style="align-items: center; align-content: center;">
                                        <span>$</span>
                                        <input type="text" placeholder="0.00" bind:value={newWeekKitchenBudget}>
                                    </div></td>
                                    <td><div class="flex compact-gap nowrap" style="align-items: center; align-content: center;">
                                        <span>$</span>
                                        <input type="text" placeholder="0.00" bind:value={newWeekFoodBudget}>
                                    </div></td>
                                </tr>
                            </tbody>
                        </table>
                        <div class="flex" style="padding-left: 1rem; padding-bottom: 1rem;">
                            <button onclick={async () => {
                                
                            }}>
                                <CheckmarkIcon></CheckmarkIcon>
                                Save
                            </button>
                        </div>
                    </div>
                {:else if showWeekBudget}
                    Budget
                    <div style="border: 0.2rem solid var(--border); border-radius: 0.8rem; margin-top: 1rem;">
                        <table style="border: none;">
                            <thead>
                                <tr>
                                    <th></th>
                                    <th>Floor Pay</th>
                                    <th>Kitchen Pay</th>
                                    <th>Food Cost</th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr>
                                    <td>Start</td>
                                    <td>{weekData.startFloorBudget}</td>
                                </tr>
                            </tbody>
                        </table>
                    </div>
                {/if}
            </div>
            <div class="areafloor">
                <div class="flex" style="justify-content: space-between;">
                    <p>Floor Workers</p>
                    <p class="fg0">
                        {floorEmployeesCount} employees
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
                        {kitchenEmployeesCount} employees
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
                <button onclick={() => {
                    for (let index = 0; index < floorHoursArray.length; index++) {
                        if (isNaN(parseFloat(floorHoursArray[index].hours))) {
                            floorHoursArray[index].hours = 0;
                        } else {
                            floorHoursArray[index].hours = parseFloat(
                                floorHoursArray[index].hours
                            );
                        }
                    }
                    for (let index = 0; index < kitchenHoursArray.length; index++) {
                        if (isNaN(parseFloat(kitchenHoursArray[index].hours))) {
                            kitchenHoursArray[index].hours = 0;
                        } else {
                            kitchenHoursArray[index].hours = parseFloat(
                                kitchenHoursArray[index].hours
                            );
                        }
                    }
                    const hoursWorked = floorHoursArray.concat(kitchenHoursArray);

                    const workedToday = floorWorkedTodayArray.concat(kitchenWorkedTodayArray);
                    for (let index = 0; index < foodCostsArray.length; index++) {
                        foodCostsArray[index].cost = foodCostsArray[index].cost.replaceAll(
                            ",", "" /* remove commas for thousands before parsing as float */
                        )
                        if (isNaN(parseFloat(foodCostsArray[index].cost))) {
                            foodCostsArray[index].cost = 0;
                        } else {
                            foodCostsArray[index].cost = parseFloat(
                                foodCostsArray[index].cost
                            );
                        }
                    }

                    console.log(JSON.stringify({
                            date: dateYMD,
                            hoursWorked: hoursWorked,
                            workedToday: workedToday,
                            foodCosts: foodCostsArray
                    }))
                    fetch(data.PUBLIC_API_URL + "/days", {
                        method: "POST",
                        headers: {
                            "Content-Type": "application/json",
                            "Authorization": "Bearer " + localStorage.getItem("budgeter:auth")
                        },
                        body: JSON.stringify({
                            date: dateYMD,
                            hoursWorked: hoursWorked,
                            workedToday: workedToday,
                            foodCosts: foodCostsArray
                        })
                    })
                }}>
                    <CheckmarkIcon></CheckmarkIcon>
                    Save
                </button>
            </div>
        </div>
    </div>
</div>
