<script>
import BackArrowIcon from "$lib/icons/BackArrow.svelte";
import { base } from "$app/paths";
let { data } = $props();
let employees = $state(data.employees ?? []);
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
        {#each employees as employee}
            <tr>
                {#if employee.edit}
                <td><input type="text" placeholder="Name" bind:value={employee.name} style="width: 8rem;"></td>
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
                    <input type="text" placeholder="Name" bind:value={employee.wage} style="width: 8rem;">
                </td>
                <td><button class="alt" onclick={async function () {
                    employee.edit = false
                    if (employee.id == null) {
                        try {
                            const res = await (
                                await fetch(data.PUBLIC_API_URL + "/employees", {
                                    method: "POST",
                                    headers: {
                                        "Authorization": "Bearer " + localStorage.getItem("auth"),
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
                                        "Authorization": "Bearer " + localStorage.getItem("auth"),
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
                }}>Save</button></td>
                {:else}
                <td>{employee.name}</td>
                <td>{employee.type.toLowerCase()}</td>
                <td>${employee.wage}/hr</td>
                <td><button class="alt" onclick={function () {
                    employee.edit = true
                }}>Edit</button></td>
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
