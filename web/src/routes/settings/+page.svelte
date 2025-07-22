<script>
import BackArrowIcon from "$lib/icons/BackArrow.svelte";
import PencilIcon from "$lib/icons/Pencil.svelte";
import MoreDotsIcon from "$lib/icons/MoreDotsVertical.svelte";
import CheckmarkIcon from "$lib/icons/Checkmark.svelte";
import TrashIcon from "$lib/icons/Trash.svelte";
import IconPlus from "$lib/icons/Plus.svelte";
import IconSettings from "$lib/icons/SettingsCogGear.svelte";
import LockIcon from "$lib/icons/Lock.svelte";
import { base } from "$app/paths";
let { data } = $props();
let users = $state(data.users);
let showDeleteUserConfirmation = $state(false);
let userToDeleteIndex = $state(-1);
let showAddNewUserMenu = $state(false);
let newUserUsername;
let newUserPassword;
async function newUser(username, password) {
    try {
        const res = await (
            await fetch(data.PUBLIC_API_URL + "/users", {
                method: "POST",
                headers: {
                    "Authorization": "Bearer " + localStorage.getItem("budgeter:auth"),
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({
                    username: username,
                    newPassword: password,
                })
            })
        ).json();
        if (res?.id) {
            employee.id = res.id;
            users.push({
                id: id,
                username: username
            })
        } else {
            console.log(res);
            console.error("id not returned?")
        }
    } catch (err) {
        console.error("Error adding user: ", err);
        alert("Error while adding user :( ");
    }
}
async function deleteUser() {
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
<p>Add/Edit Users</p>
<table style="min-width: 17rem;">
    <tbody>
        {#each users as user, userIndex}
            <tr>
                <td><div class="flex" style="justify-content: space-between; align-items: center; align-content: center;">
                    <span>{user.username}</span>
                    <div class="dropdown">
                        <button class="icon-only-button">
                            <MoreDotsIcon></MoreDotsIcon>
                        </button>
                        <div class="content">
                            <button>
                                <PencilIcon></PencilIcon>
                                Edit Username
                            </button>
                            <button>
                                <LockIcon></LockIcon>
                                Edit Password
                            </button>
                            <button class="ohno">
                                <TrashIcon></TrashIcon>
                                Delete
                            </button>
                        </div>
                    </div>
                </div></td>
            </tr>
        {/each}
        {#if showAddNewUserMenu}
            <tr>
                <td><div class="flex col">
                    <input type="text" placeholder="Username" bind:value={newUserUsername} style="width: 14rem;">
                    <input type="password" placeholder="Password" bind:value={newUserPassword} style="width: 14rem;">
                    <button style="width: 14rem;">
                        <CheckmarkIcon></CheckmarkIcon>
                        Add
                    </button>
                    <button class="alt" onclick={() => {
                        showAddNewUserMenu = false;
                        newUserUsername = "";
                        newUserPassword = "";
                    }} style="width: 14rem;">
                        Cancel
                    </button>
                </div></td>
            </tr>
        {:else}
        <tr>
            <td>
                <button onclick={function () {
                    newUserUsername = "";
                    newUserPassword = "";
                    showAddNewUserMenu = true;
                }}>
                    <IconPlus></IconPlus>
                    Add
                </button>
            </td>
        </tr>
        {/if}
    </tbody>
</table>
</div>
</div>
{#if showDeleteUserConfirmation}
    <div class="modal">
        <div class="content">
            <p>Are you sure you want to delete "{users[userToDeleteIndex].username}"?</p>
            <div class="flex">
                <button class="ohno" onclick={deleteUser}>
                    <TrashIcon></TrashIcon>
                    Delete
                </button>
                <button class="alt" onclick={() => {
                    showDeleteUserConfirmation = false;
                    userToDeleteIndex = -1;
                }}>Cancel</button>
            </div>
        </div>
    </div>
{/if}
