<script>
import BackArrowIcon from "$lib/icons/BackArrow.svelte";
import { base } from "$app/paths";
let { data } = $props();
let employees = $state(data.employees);
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
                        <select bind:value={employee.type}>
                            <option value="floor">Floor</option>
                            <option value="kitchen">Kitchen</option>
                        </select>
                    </div>
                </td>
                <td>
                    <input type="text" placeholder="Name" bind:value={employee.wage} style="width: 8rem;">
                </td>
                <td><button class="alt" onclick={function () {
                    employee.edit = false
                }}>Save</button></td>
                {:else}
                <td>{employee.name}</td>
                <td>{employee.type}</td>
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
