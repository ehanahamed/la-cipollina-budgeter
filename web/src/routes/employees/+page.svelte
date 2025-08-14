<script>
import BackArrowIcon from "$lib/icons/BackArrow.svelte";
import PencilIcon from "$lib/icons/Pencil.svelte";
import MoreDotsIcon from "$lib/icons/MoreDotsVertical.svelte";
import CheckmarkIcon from "$lib/icons/Checkmark.svelte";
import TrashIcon from "$lib/icons/Trash.svelte";
import IconPlus from "$lib/icons/Plus.svelte";
import IconSettings from "$lib/icons/SettingsCogGear.svelte";
import { base } from "$app/paths";
import { fade } from "svelte/transition";
let { data } = $props();
let employees = $state(data.employees ?? []);
let showDeleteEmployeeConfirmation = $state(false);
let employeeToDeleteIndex = $state(-1);
let showSpecialPayEditing = $state(false);
let specialPayEditingEmployeeIndex = $state(-1);
let floorEmployeesCount = $derived.by(() => {
    let count = 0;
    employees.forEach((employee) => {
        if (employee.type.toUpperCase() == "FLOOR") {
            count++;
        }
    });
    return count;
});
let kitchenEmployeesCount = $derived.by(() => {
    let count = 0;
    employees.forEach((employee) => {
        if (employee.type.toUpperCase() == "KITCHEN") {
            count++;
        }
    });
    return count;
});
async function saveEmployee(employee) {
    if (
        employee.weeklyPay === "" &&
        employee.specialPay == null &&
        employee.wage == null
    ) {
        employee.weeklyPay = 0;
    }
    if (employee.specialPay) {
        employee.wage = null;
        employee.weeklyPay = null;
    } else if (!isNaN(parseFloat(employee.weeklyPay))) {
        employee.wage = null;
        employee.specialPay = null;
        if (typeof employee.weeklyPay === "string") {
            /* remove any user-inputted commas for thousands */
            employee.weeklyPay = employee.weeklyPay.replaceAll(",", "");
        }
        employee.weeklyPay = parseFloat(employee.weeklyPay);
    } else if (isNaN(parseFloat(employee.wage))) {
        /* isNaN(null) returns false
        isNaN(parseFloat(null)) returns true */
        employee.wage = 0;
    } else {
        if (typeof employee.wage === "string") {
            /* remove any user-inputted commas for thousands */
            employee.wage = employee.wage.replaceAll(",", "");
        }
        employee.wage = parseFloat(employee.wage);
    }
    if (employee.id == null) {
        try {
            const res = await (
                await fetch(data.PUBLIC_API_URL + "/employees", {
                    method: "POST",
                    headers: {
                        "Authorization": "Bearer " + localStorage.getItem("budgeter:auth"),
                        "Content-Type": "application/json"
                    },
                    body: JSON.stringify({
                        name: employee.name,
                        type: employee.type.toUpperCase(),
                        wage: employee.wage,
                        specialPay: employee.specialPay,
                        weeklyPay: employee.weeklyPay
                    })
                })
            ).json();
            if (res?.id) {
                employee.id = res.id;
            } else {
                console.log(res);
                console.error("id not returned?")
            }
        } catch (err) {
            console.error("Error adding employee: ", err);
            alert("Error while adding employee :( ");
        }
    } else {
        try {
            const res = await (
                await fetch(data.PUBLIC_API_URL + "/employees/" + employee.id, {
                    method: "PUT",
                    headers: {
                        "Authorization": "Bearer " + localStorage.getItem("budgeter:auth"),
                        "Content-Type": "application/json"
                    },
                    body: JSON.stringify({
                        name: employee.name,
                        type: employee.type.toUpperCase(),
                        wage: employee.wage,
                        specialPay: employee.specialPay,
                        weeklyPay: employee.weeklyPay
                    })
                })
            ).json();
        } catch (err) {
            console.error("Error adding employee: ", err);
            alert("Error while adding employee :( ");
        }
    }
}
async function removeEmployee() {
    try {
        const res = await (
            await fetch(data.PUBLIC_API_URL + "/employees/" + employees[employeeToDeleteIndex].id, {
                method: "DELETE",
                headers: {
                    "Authorization": "Bearer " + localStorage.getItem("budgeter:auth")
                }
            })
        ).json();
        if (res?.success) {
            employees.splice(employeeToDeleteIndex, 1);
            employees = employees; /* to update state/trigger reactivity after splice-ing */
        } else {
            console.log(res);
            alert("Error, couldn't delete ???");
        }
    } catch (err) {
        console.error("Error while deleting employee: ", err);
        alert("Error deleting employee :( ");
    } finally {
        showDeleteEmployeeConfirmation = false;
        employeeToDeleteIndex = -1;
    }
}
</script>
<div class="grid page" style="margin-top: 4rem; margin-bottom: 10rem;">
    <div class="content">
                <div style="margin-bottom: 1rem;"><a href={base} class="button faint">
                    <BackArrowIcon></BackArrowIcon>
                    Back
                </a></div>
<div>
<div style="display: inline-block;">
<div class="flex" style="justify-content: space-between;">
    <p>Add/Edit Employees</p>
    <p class="fg0">{employees.length} total, {floorEmployeesCount} floor, {kitchenEmployeesCount} kitchen</p>
</div>
<table>
    <thead>
        {#if employees.some((employee) => employee.weeklyPay != null)}
            <tr>
                <th>Name</th>
                <th>Type</th>
                <th>Weekly Pay</th>
            </tr>
        {:else}
            <tr>
                <th>Name</th>
                <th>Type</th>
                <th>Wage</th>
            </tr>
        {/if}
    </thead>
    <tbody>
        {#each employees as employee, employeeIndex}
            {#if employee.weeklyPay != null}
            <tr>
                {#if employee.edit}
                    <td style="vertical-align: top;"><input type="text" placeholder="Name" bind:value={employee.name} style="min-width: 6rem; field-sizing: content;"></td>
                    <td style="vertical-align: top;">
                        <div class="select-wrapper" style="width: 7rem;">
                            <select bind:value={
                                () => employee.type.toLowerCase(), /* get */
                                (newType) => employee.type = newType.toUpperCase() /* set */
                            }>
                                <option value="floor">floor</option>
                                <option value="kitchen">kitchen</option>
                            </select>
                        </div>
                    </td>
                    <td>
                        <div class="flex col">
                            {#if employee.specialPay}
                                <button class="alt" style="min-width: 9rem;" onclick={() => {
                                    showSpecialPayEditing = true;
                                    specialPayEditingEmployeeIndex = employeeIndex;
                                }}>
                                    <IconSettings></IconSettings>
                                    Special pay
                                </button>
                            {:else}
                            <div class="flex compact-gap nowrap" style="align-items: center; align-content: center;"><span>$</span><input type="text" placeholder="12.34" bind:value={employee.weeklyPay} style="min-width: 6rem; field-sizing: content;"></div>
                            <button class="alt" onclick={() => {
                                showSpecialPayEditing = true;
                                specialPayEditingEmployeeIndex = employeeIndex;
                            }}>
                                <IconSettings></IconSettings>
                                Settings
                            </button>
                            {/if}
                        </div>
                    </td>
                    <td style="vertical-align: top;"><button onclick={() => {
                        employee.edit = false;
                        saveEmployee(employee);
                    }}><CheckmarkIcon></CheckmarkIcon> Save</button></td>
                {:else}
                    <td>{employee.name}</td>
                    <td>{employee.type.toLowerCase()}</td>
                    <td>
                    {#if employee.specialPay}
                        <span class="fg0">special pay</span>
                    {:else}
                        ${parseFloat(employee.weeklyPay).toLocaleString('en-US', {
                            minimumFractionDigits: 2,
                            maximumFractionDigits: 2
                        })}
                    {/if}
                    </td>
                    <td>
                        <div class="flex" style="flex-wrap: nowrap;">
                            <button class="alt" onclick={function () {
                                employee.edit = true
                            }}><PencilIcon></PencilIcon> Edit</button>
                            <div class="dropdown" tabindex="0">
                                <button class="icon-only-button" aria-label="dropdown">
                                    <MoreDotsIcon></MoreDotsIcon>
                                </button>
                                <div class="content">
                                    <button class="ohno" onclick={() => {
                                        employeeToDeleteIndex = employeeIndex;
                                        if (employee.name) {
                                            showDeleteEmployeeConfirmation = true;
                                        } else {
                                            /* don't show confirmation modal/alert
                                            if name is blank, just directly delete it */
                                            removeEmployee();
                                        }
                                    }}>
                                        <TrashIcon></TrashIcon> Delete
                                    </button>
                                </div>
                            </div>
                        </div>
                    </td>
                {/if}
            </tr>
            {/if}
        {/each}
        {#if employees.some((employee) => employee.weeklyPay != null)}
            <tr>
                <th>Name</th>
                <th>Type</th>
                <th>Wage</th>
            </tr>
        {/if}
        {#each employees as employee, employeeIndex}
            {#if employee.weeklyPay == null}
            <tr>
                {#if employee.edit}
                    <td style="vertical-align: top;"><input type="text" placeholder="Name" bind:value={employee.name} style="min-width: 6rem; field-sizing: content;"></td>
                    <td style="vertical-align: top;">
                        <div class="select-wrapper" style="width: 7rem;">
                            <select bind:value={
                                () => employee.type.toLowerCase(), /* get */
                                (newType) => employee.type = newType.toUpperCase() /* set */
                            }>
                                <option value="floor">floor</option>
                                <option value="kitchen">kitchen</option>
                            </select>
                        </div>
                    </td>
                    <td>
                        <div class="flex col">
                            {#if employee.specialPay}
                                <button class="alt" style="min-width: 9rem;" onclick={() => {
                                    showSpecialPayEditing = true;
                                    specialPayEditingEmployeeIndex = employeeIndex;
                                }}>
                                    <IconSettings></IconSettings>
                                    Special pay
                                </button>
                            {:else}
                            <div class="flex compact-gap nowrap" style="align-items: center; align-content: center;"><span>$</span><input type="text" placeholder="12.34" bind:value={employee.wage} style="min-width: 6rem; field-sizing: content;"></div>
                            <button class="alt" onclick={() => {
                                showSpecialPayEditing = true;
                                specialPayEditingEmployeeIndex = employeeIndex;
                            }}>
                                <IconSettings></IconSettings>
                                Settings
                            </button>
                            {/if}
                        </div>
                    </td>
                    <td style="vertical-align: top;"><button onclick={() => {
                        employee.edit = false;
                        saveEmployee(employee);
                    }}><CheckmarkIcon></CheckmarkIcon> Save</button></td>
                {:else}
                    <td>{employee.name}</td>
                    <td>{employee.type.toLowerCase()}</td>
                    <td>
                    {#if employee.specialPay}
                        <span class="fg0">special pay</span>
                    {:else}
                        ${parseFloat(employee.wage).toLocaleString('en-US', {
                            minimumFractionDigits: 2,
                            maximumFractionDigits: 2
                        })}
                    {/if}
                    </td>
                    <td>
                        <div class="flex" style="flex-wrap: nowrap;">
                            <button class="alt" onclick={function () {
                                employee.edit = true
                            }}><PencilIcon></PencilIcon> Edit</button>
                            <div class="dropdown" tabindex="0">
                                <button class="icon-only-button" aria-label="dropdown">
                                    <MoreDotsIcon></MoreDotsIcon>
                                </button>
                                <div class="content">
                                    <button class="ohno" onclick={() => {
                                        employeeToDeleteIndex = employeeIndex;
                                        if (employee.name) {
                                            showDeleteEmployeeConfirmation = true;
                                        } else {
                                            /* don't show confirmation modal/alert
                                            if name is blank, just directly delete it */
                                            removeEmployee();
                                        }
                                    }}>
                                        <TrashIcon></TrashIcon> Delete
                                    </button>
                                </div>
                            </div>
                        </div>
                    </td>
                {/if}
            </tr>
            {/if}
        {/each}
        <tr>
            <td>
                <button onclick={function () {
                    employees.push({
                        name: "",
                        type: "FLOOR",
                        wage: "",
                        edit: true
                    })
                }}>
                    <IconPlus></IconPlus>
                    Add
                </button>
            </td>
        </tr>
    </tbody>
</table>
</div>
</div>
</div>
</div>
{#if showDeleteEmployeeConfirmation}
    <div class="modal" transition:fade={{ duration: 200 }}>
        <div class="content">
            <p>Are you sure you want to remove "{employees[employeeToDeleteIndex].name}"?</p>
            <div class="flex">
                <button class="ohno" onclick={removeEmployee}>
                    <TrashIcon></TrashIcon>
                    Remove
                </button>
                <button class="alt" onclick={() => {
                    showDeleteEmployeeConfirmation = false;
                    employeeToDeleteIndex = -1;
                }}>Cancel</button>
            </div>
        </div>
    </div>
{/if}
{#if showSpecialPayEditing}
    <div class="modal" transition:fade={{ duration: 200 }}>
        <div class="content">
            {#if employees[specialPayEditingEmployeeIndex].name}
                <p>Special Pay Settings for {employees[specialPayEditingEmployeeIndex].name}</p>
            {:else}
                <p>Special Pay Settings</p>
            {/if}
            <div class="combo-select">
                <button class="left {
                    employees[specialPayEditingEmployeeIndex].specialPay ||
                    employees[specialPayEditingEmployeeIndex].weeklyPay != null ?
                        "" : "selected"
                }" onclick={() => {
                    employees[specialPayEditingEmployeeIndex].specialPay = undefined;
                    employees[specialPayEditingEmployeeIndex].weeklyPay = undefined;
                }}>
                    <CheckmarkIcon class="combo-selected-icon"></CheckmarkIcon>
                    Hourly wage
                </button>
                <button class="mid {
                    employees[specialPayEditingEmployeeIndex].specialPay ? "selected" : ""
                }" onclick={() => {
                    employees[specialPayEditingEmployeeIndex].weeklyPay = undefined;
                    employees[specialPayEditingEmployeeIndex].specialPay = {
                        mon: {
                            perDay: null,
                            perHour: null
                        },
                        tue: {
                            perDay: null,
                            perHour: null
                        },
                        wed: {
                            perDay: null,
                            perHour: null
                        },
                        thu: {
                            perDay: null,
                            perHour: null
                        },
                        fri: {
                            perDay: null,
                            perHour: null
                        },
                        sat: {
                            perDay: null,
                            perHour: null
                        },
                        sun: {
                            perDay: null,
                            perHour: null
                        },
                    };
                }}>
                    <CheckmarkIcon class="combo-selected-icon"></CheckmarkIcon>
                    Special pay
                </button>
                <button class="right {
                    employees[specialPayEditingEmployeeIndex].weeklyPay != null ? "selected" : ""
                }" onclick={() => {
                    employees[specialPayEditingEmployeeIndex].specialPay = undefined;
                    employees[specialPayEditingEmployeeIndex].wage = null;
                    employees[specialPayEditingEmployeeIndex].weeklyPay = 0;
                }}>
                    <CheckmarkIcon class="combo-selected-icon"></CheckmarkIcon>
                    Weekly Pay
                </button>
            </div>
            {#if employees[specialPayEditingEmployeeIndex].specialPay}
                <table style="border: none;">
                    <thead>
                        <tr>
                            <th>Day</th>
                            <th>Daily</th>
                            <th>Hourly</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr>
                            <td>Monday</td>
                            <td>
                                <div class="flex compact-gap nowrap" style="align-items: center; align-content: center;">
                                    <span>$</span>
                                    <input type="text" placeholder="" bind:value={
                                        employees[
                                            specialPayEditingEmployeeIndex
                                        ].specialPay.mon.perDay
                                    } style="min-width: 5rem; field-sizing: content;">
                                </div>
                            </td>
                            <td>
                                <div class="flex compact-gap nowrap" style="align-items: center; align-content: center;">
                                    <span>$</span>
                                    <input type="text" placeholder="" bind:value={
                                        employees[
                                            specialPayEditingEmployeeIndex
                                        ].specialPay.mon.perHour
                                    } style="min-width: 5rem; field-sizing: content;">
                                </div>
                            </td>
                        </tr>
                        <tr>
                            <td>Tuesday</td>
                            <td>
                                <div class="flex compact-gap nowrap" style="align-items: center; align-content: center;">
                                    <span>$</span>
                                    <input type="text" placeholder="" bind:value={
                                        employees[
                                            specialPayEditingEmployeeIndex
                                        ].specialPay.tue.perDay
                                    } style="min-width: 5rem; field-sizing: content;">
                                </div>
                            </td>
                            <td>
                                <div class="flex compact-gap nowrap" style="align-items: center; align-content: center;">
                                    <span>$</span>
                                    <input type="text" placeholder="" bind:value={
                                        employees[
                                            specialPayEditingEmployeeIndex
                                        ].specialPay.tue.perHour
                                    } style="min-width: 5rem; field-sizing: content;">
                                </div>
                            </td>
                        </tr>
                        <tr>
                            <td>Wednesday</td>
                            <td>
                                <div class="flex compact-gap nowrap" style="align-items: center; align-content: center;">
                                    <span>$</span>
                                    <input type="text" placeholder="" bind:value={
                                        employees[
                                            specialPayEditingEmployeeIndex
                                        ].specialPay.wed.perDay
                                    } style="min-width: 5rem; field-sizing: content;">
                                </div>
                            </td>
                            <td>
                                <div class="flex compact-gap nowrap" style="align-items: center; align-content: center;">
                                    <span>$</span>
                                    <input type="text" placeholder="" bind:value={
                                        employees[
                                            specialPayEditingEmployeeIndex
                                        ].specialPay.wed.perHour
                                    } style="min-width: 5rem; field-sizing: content;">
                                </div>
                            </td>
                        </tr>
                        <tr>
                            <td>Thursday</td>
                            <td>
                                <div class="flex compact-gap nowrap" style="align-items: center; align-content: center;">
                                    <span>$</span>
                                    <input type="text" placeholder="" bind:value={
                                        employees[
                                            specialPayEditingEmployeeIndex
                                        ].specialPay.thu.perDay
                                    } style="min-width: 5rem; field-sizing: content;">
                                </div>
                            </td>
                            <td>
                                <div class="flex compact-gap nowrap" style="align-items: center; align-content: center;">
                                    <span>$</span>
                                    <input type="text" placeholder="" bind:value={
                                        employees[
                                            specialPayEditingEmployeeIndex
                                        ].specialPay.thu.perHour
                                    } style="min-width: 5rem; field-sizing: content;">
                                </div>
                            </td>
                        </tr>
                        <tr>
                            <td>Friday</td>
                            <td>
                                <div class="flex compact-gap nowrap" style="align-items: center; align-content: center;">
                                    <span>$</span>
                                    <input type="text" placeholder="" bind:value={
                                        employees[
                                            specialPayEditingEmployeeIndex
                                        ].specialPay.fri.perDay
                                    } style="min-width: 5rem; field-sizing: content;">
                                </div>
                            </td>
                            <td>
                                <div class="flex compact-gap nowrap" style="align-items: center; align-content: center;">
                                    <span>$</span>
                                    <input type="text" placeholder="" bind:value={
                                        employees[
                                            specialPayEditingEmployeeIndex
                                        ].specialPay.fri.perHour
                                    } style="min-width: 5rem; field-sizing: content;">
                                </div>
                            </td>
                        </tr>
                        <tr>
                            <td>Saturday</td>
                            <td>
                                <div class="flex compact-gap nowrap" style="align-items: center; align-content: center;">
                                    <span>$</span>
                                    <input type="text" placeholder="" bind:value={
                                        employees[
                                            specialPayEditingEmployeeIndex
                                        ].specialPay.sat.perDay
                                    } style="min-width: 5rem; field-sizing: content;">
                                </div>
                            </td>
                            <td>
                                <div class="flex compact-gap nowrap" style="align-items: center; align-content: center;">
                                    <span>$</span>
                                    <input type="text" placeholder="" bind:value={
                                        employees[
                                            specialPayEditingEmployeeIndex
                                        ].specialPay.sat.perHour
                                    } style="min-width: 5rem; field-sizing: content;">
                                </div>
                            </td>
                        </tr>
                        <tr>
                            <td>Sunday</td>
                            <td>
                                <div class="flex compact-gap nowrap" style="align-items: center; align-content: center;">
                                    <span>$</span>
                                    <input type="text" placeholder="" bind:value={
                                        employees[
                                            specialPayEditingEmployeeIndex
                                        ].specialPay.sun.perDay
                                    } style="min-width: 5rem; field-sizing: content;">
                                </div>
                            </td>
                            <td>
                                <div class="flex compact-gap nowrap" style="align-items: center; align-content: center;">
                                    <span>$</span>
                                    <input type="text" placeholder="" bind:value={
                                        employees[
                                            specialPayEditingEmployeeIndex
                                        ].specialPay.sun.perHour
                                    } style="min-width: 5rem; field-sizing: content;">
                                </div>
                            </td>
                        </tr>
                    </tbody>
                </table>
                <p class="fg0">Leave the ones you don't need blank :)</p>
            {/if}
            <div class="flex">
                <button onclick={() => {
                    if (employees[specialPayEditingEmployeeIndex].specialPay) {
                        [
                            "mon",
                            "tue",
                            "wed",
                            "thu",
                            "fri",
                            "sat",
                            "sun"
                        ].forEach(function (day) {
                        /* parse strings as floats when done is pressed */
                            const dayObj = employees[
                                specialPayEditingEmployeeIndex
                            ].specialPay[day]; /* reference to the object,
                            so the og array & obj gets updated correctly */
                            if (isNaN(parseFloat(dayObj.perDay))) {
                                /* isNaN(null) returns false
                                isNaN(parseFloat(null)) returns true */
                                dayObj.perDay = null;
                            } else {
                                if (typeof dayObj.perDay === "string") {
                                    /* remove any user-inputted commas for thousands */
                                    dayObj.perDay = dayObj.perDay.replaceAll(",", "");
                                }
                                dayObj.perDay = parseFloat(dayObj.perDay);
                            }
                            if (isNaN(parseFloat(dayObj.perHour))) {
                                dayObj.perHour = null;
                            } else {
                                if (typeof dayObj.perHour === "string") {
                                    /* remove any user-inputted commas for thousands */
                                    dayObj.perHour = dayObj.perHour.replaceAll(",", "");
                                }
                                dayObj.perHour = parseFloat(dayObj.perHour);
                            }
                        })
                    }

                    showSpecialPayEditing = false;
                    saveEmployee(employees[specialPayEditingEmployeeIndex]);
                }}><CheckmarkIcon></CheckmarkIcon> Done</button>
            </div>
        </div>
    </div>
{/if}
