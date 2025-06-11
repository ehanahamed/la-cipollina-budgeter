<script>
import { base } from "$app/paths";
import BackArrowIcon from "$lib/icons/BackArrow.svelte";
import PlusIcon from "$lib/icons/Plus.svelte";
import CheckmarkIcon from "$lib/icons/Checkmark.svelte";
import { onMount } from "svelte";
let { data } = $props();
let floorHoursArray = $state([]);
let kitchenHoursArray = $state([]);
if (data.new) {
    data.employees.forEach(function (employee) {
        if (employee?.type == "floor") {
            floorHoursArray.push({
                name: employee.name,
                hours: ""
            })
        }
        if (employee?.type == "kitchen") {
            kitchenHoursArray.push({
                name: employee.name,
                hours: ""
            })
        }
    })
}
let ingredientExpensesArray = $state([
    {
        increase: "",
        decrease: ""
    }
]);
let dayOfTheWeek = $state();
let weekDays = [
    "Sunday",
    "Monday",
    "Tuesday",
    "Wednesday",
    "Thrusday",
    "Friday",
    "Saturday"
]
let showSetBudget = $state(false);
onMount(function () {
    let today = new Date()
    dayOfTheWeek = weekDays[today.getDay()]
    if (
        dayOfTheWeek == "Monday" ||
        dayOfTheWeek == "Tuesday" ||
        dayOfTheWeek == "Wednesday"
    ) {
        showSetBudget = true;
    }
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
.gridtablelayout input[type=text] {
    width: 8rem;
}
.centerbutnotonmobile {
    text-align: center;
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
<div class="grid page" style="margin-top: 2rem;">
    <div class="content">
        <div class="centerbutnotonmobile">
            {#if data.new}
                <p>Run Day</p>
            {:else}
                <p>Edit Day</p>
            {/if}
            <h3 style="margin-top: 0px;">{dayOfTheWeek}</h3>
        </div>
        {#if showSetBudget}
            <div>
                Set budget for this week:
                <table class="setbudgettable">
                    <thead>
                        <tr>
                            <th>Floor Pay</th>
                            <th>Kitchen Pay</th>
                            <th>Food Cost</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr>
                            <td><input type="text"></td>
                        </tr>
                    </tbody>
                </table>
            </div>
        {/if}
        <div class="gridtablelayout">
            <div class="areaback">
                <a href={base} class="button faint">
                    <BackArrowIcon></BackArrowIcon>
                    Back
                </a>
            </div>
            <div class="areafloor">
                <p>Floor Workers</p>
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
                                <td>{row.name}</td>
                                <td><input type="text" bind:value={row.hours}></td>
                            </tr>
                        {/each}
                    </tbody>
                </table>
            </div>
            <div class="areakitchen">
                <p>Kitchen Workers</p>
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
                                <td>{row.name}</td>
                                <td><input type="text" bind:value={row.hours}></td>
                            </tr>
                        {/each}
                    </tbody>
                </table>
            </div>
            <div class="areafood">
                <p>Ingredient Expenses</p>
                <table>
                    <thead>
                        <tr>
                            <th>Increases</th>
                            <th>Decreases</th>
                        </tr>
                    </thead>
                    <tbody>
                        {#each ingredientExpensesArray as row}
                            <tr>
                                <td><input type="text" bind:value={row.increase}></td>
                                <td><input type="text" bind:value={row.decrease}></td>
                            </tr>
                        {/each}
                        <tr>
                            <td>
                                <button onclick={function () {
                                    ingredientExpensesArray.push({increase:"",decrease:""})
                                }}>
                                    <PlusIcon></PlusIcon>
                                    Add
                                </button>
                            </td>
                        </tr>
                    </tbody>
                </table>
            </div>
            <div class="areasave">
                <button>
                    <CheckmarkIcon></CheckmarkIcon>
                    Save
                </button>
            </div>
        </div>
    </div>
</div>
