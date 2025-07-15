<script>
import BackArrowIcon from "$lib/icons/BackArrow.svelte";
import PencilIcon from "$lib/icons/Pencil.svelte";
import MoreDotsIcon from "$lib/icons/MoreDotsVertical.svelte";
import CheckmarkIcon from "$lib/icons/Checkmark.svelte";
import TrashIcon from "$lib/icons/Trash.svelte";
import { base } from "$app/paths";
let { data } = $props();
let employees = $state(data.employees ?? []);
let showDeleteEmployeeConfirmation = $state(false);
let employeeToDeleteIndex = $state(-1);
</script>
<div class="grid page">
    <div class="content">
                <div style="margin-top: 1rem; margin-bottom: 1rem;"><a href={base} class="button faint">
                    <BackArrowIcon></BackArrowIcon>
                    Back
                </a></div>
Add/Edit Employees
<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Wage</th>
        </tr>
    </thead>
    <tbody>
        {#each employees as employee, employeeIndex}
            <tr>
                {#if employee.edit}
                    <td><input type="text" placeholder="Name" bind:value={employee.name} style="min-width: 6rem; field-sizing: content;"></td>
                    <td>
                        <div class="select-wrapper" style="width: 8rem;">
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
                        <input type="text" placeholder="Name" bind:value={employee.wage} style="min-width: 6rem; field-sizing: content;">
                    </td>
                    <td><button onclick={async function () {
                        employee.edit = false
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
                                            wage: parseFloat(employee.wage)
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
                                            wage: parseFloat(employee.wage)
                                        })
                                    })
                                ).json();
                            } catch (err) {
                                console.error("Error adding employee: ", err);
                                alert("Error while adding employee :( ");
                            }
                        }
                    }}><CheckmarkIcon></CheckmarkIcon> Save</button></td>
                {:else}
                    <td>{employee.name}</td>
                    <td>{employee.type.toLowerCase()}</td>
                    <td>${employee.wage}/hr</td>
                    <td>
                        <div class="flex" style="flex-wrap: nowrap;">
                            <button class="alt" onclick={function () {
                                employee.edit = true
                            }}><PencilIcon></PencilIcon> Edit</button>
                            <div class="dropdown">
                                <button class="icon-only-button" aria-label="dropdown">
                                    <MoreDotsIcon></MoreDotsIcon>
                                </button>
                                <div class="content">
                                    <button class="ohno" onclick={() => {
                                        showDeleteEmployeeConfirmation = true;
                                        employeeToDeleteIndex = employeeIndex;
                                    }}>
                                        <TrashIcon></TrashIcon> Delete
                                    </button>
                                </div>
                            </div>
                        </div>
                    </td>
                {/if}
            </tr>
        {/each}
        <tr>
            <td>
                <button onclick={function () {
                    employees.push({
                        name: "",
                        type: "",
                        wage: 0,
                        edit: true
                    })
                }}>
                    Add
                </button>
            </td>
        </tr>
    </tbody>
</table>
</div>
</div>
{#if showDeleteEmployeeConfirmation}
    <div class="modal">
        <div class="content">
            <p>Are you sure you want to remove "{employees[employeeToDeleteIndex].name}"?</p>
            <div class="flex">
                <button class="ohno" onclick={async function () {
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
                }}><TrashIcon></TrashIcon> Remove</button>
                <button class="alt" onclick={() => {
                    showDeleteEmployeeConfirmation = false;
                    employeeToDeleteIndex = -1;
                }}>Cancel</button>
            </div>
        </div>
    </div>
{/if}
