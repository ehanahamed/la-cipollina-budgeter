<script>
import BackArrowIcon from "$lib/icons/BackArrow.svelte";
import PencilIcon from "$lib/icons/Pencil.svelte";
import MoreDotsIcon from "$lib/icons/MoreDotsVertical.svelte";
import CheckmarkIcon from "$lib/icons/Checkmark.svelte";
import TrashIcon from "$lib/icons/Trash.svelte";
import IconPlus from "$lib/icons/Plus.svelte";
import IconSettings from "$lib/icons/SettingsCogGear.svelte";
import LockIcon from "$lib/icons/Lock.svelte";
import ExitIcon from "$lib/icons/Exit.svelte";
import { base } from "$app/paths";
import { goto } from "$app/navigation";
let { data } = $props();
let users = $state(data.users);
let showDeleteUserConfirmation = $state(false);
let userToDeleteIndex = $state(-1);
let showSetNewPasswordMenu = $state(false);
let setNewPasswordUserIndex = $state(-1);
let setNewPasswordPassword;
let showAddNewUserMenu = $state(false);
let newUserUsername;
let newUserPassword;
async function createUser(username, password) {
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
            users.push({
                id: res.id,
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
            await fetch(data.PUBLIC_API_URL + "/users/" + users[userToDeleteIndex].id, {
                method: "DELETE",
                headers: {
                    "Authorization": "Bearer " + localStorage.getItem("budgeter:auth")
                }
            })
        ).json();
        if (res?.success) {
            users.splice(userToDeleteIndex, 1);
            users = users; /* to update state/trigger reactivity after splice-ing */
        } else {
            console.log(res);
            alert("Error, couldn't delete ???");
        }
    } catch (err) {
        console.error("Error while deleting user: ", err);
        alert("Error deleting user :( ");
    } finally {
        showDeleteUserConfirmation = false;
        userToDeleteIndex = -1;
    }
}
</script>
<div class="grid page" style="margin-top: 4rem; margin-bottom: 10rem;">
    <div class="content">
                <div style="margin-bottom: 1rem;"><a href={base} class="button faint">
                    <BackArrowIcon></BackArrowIcon>
                    Back
                </a></div>
<div class="box">
    <p>Currently logged in as "{data?.authedUser?.username}"</p>
    <div class="flex">
        <button class="ohno alt" onclick={() => {
            localStorage.removeItem("budgeter:auth");
            goto(base);
        }}>
            <ExitIcon></ExitIcon>
            Sign Out
        </button>
    </div>
</div>
<p>Add/Edit Users</p>
<table style="min-width: 17rem;">
    <tbody>
        {#each users as user, userIndex}
            <tr>
                {#if user.editUsername}
                    <td><div class="flex">
                        <input type="text" placeholder="Username" bind:value={user.username} style="width: 14rem;">
                        <button onclick={async function () {
                            try {
                                const res = await (
                                    await fetch(data.PUBLIC_API_URL + "/users/" + user.id, {
                                        method: "PATCH",
                                        headers: {
                                            "Authorization": "Bearer " + localStorage.getItem("budgeter:auth"),
                                            "Content-Type": "application/json"
                                        },
                                        body: JSON.stringify({
                                            username: user.username
                                        })
                                    })
                                ).json();
                                if (res.error) {
                                    console.log(res);
                                    alert("Error, couldn't update ???");
                                }
                            } catch (err) {
                                console.error("Error while updating username: ", err);
                                alert("Error updating username :( ");
                            } finally {
                                user.editUsername = false;
                            }
                        }}>
                            <CheckmarkIcon></CheckmarkIcon>
                            Save
                        </button>
                    </div></td>
                {:else}
                    <td><div class="flex" style="justify-content: space-between; align-items: center; align-content: center;">
                        <span>{user.username}</span>
                        <div class="dropdown">
                            <button class="icon-only-button">
                                <MoreDotsIcon></MoreDotsIcon>
                            </button>
                            <div class="content">
                                <button onclick={() => {
                                    user.editUsername = true;
                                }}>
                                    <PencilIcon></PencilIcon>
                                    Edit Username
                                </button>
                                <button onclick={() => {
                                    setNewPasswordUserIndex = userIndex;
                                    showSetNewPasswordMenu = true;
                                }}>
                                    <LockIcon></LockIcon>
                                    Edit Password
                                </button>
                                <button class="ohno" onclick={() => {
                                    userToDeleteIndex = userIndex;
                                    showDeleteUserConfirmation = true;
                                }}>
                                    <TrashIcon></TrashIcon>
                                    Delete
                                </button>
                            </div>
                        </div>
                    </div></td>
                {/if}
            </tr>
        {/each}
        {#if showAddNewUserMenu}
            <tr>
                <td><div class="flex col">
                    <input type="text" placeholder="Username" bind:value={newUserUsername} style="width: 14rem;">
                    <input type="password" placeholder="Password" bind:value={newUserPassword} style="width: 14rem;">
                    <button style="width: 14rem;" onclick={() => {
                        if (newUserUsername == "" ||
                        newUserPassword == "") {
                            return;
                        }
                        createUser(
                            newUserUsername,
                            newUserPassword
                        );
                        showAddNewUserMenu = false;
                        newUserUsername = "";
                        newUserPassword = "";
                    }}>
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
{#if showSetNewPasswordMenu}
    <div class="modal">
        <div class="content">
            <p>New password for "{users[setNewPasswordUserIndex].username}"</p>
            <input type="password" placeholder="New Password" bind:value={setNewPasswordPassword}>
            <div class="flex">
                <button onclick={async function () {
                    try {
                        const res = await (
                            await fetch(data.PUBLIC_API_URL + "/users/" + users[setNewPasswordUserIndex].id, {
                                method: "PATCH",
                                headers: {
                                    "Authorization": "Bearer " + localStorage.getItem("budgeter:auth"),
                                    "Content-Type": "application/json"
                                },
                                body: JSON.stringify({
                                    newPassword: setNewPasswordPassword
                                })
                            })
                        ).json();
                        if (res.error) {
                            console.log(res);
                            alert("Error, couldn't update ???");
                        }
                    } catch (err) {
                        console.error("Error while updating user password: ", err);
                        alert("Error updating user password :( ");
                    } finally {
                        if (users[setNewPasswordUserIndex].id == data?.authedUser?.id) {
                            /* if user just edited their own password, bring them to the sign in screen,
                            because the server just signed them out from all devices, including the one they're on rn*/
                            goto(base);
                        }
                        showSetNewPasswordMenu = false;
                        setNewPasswordUserIndex = -1;
                        setNewPasswordPassword = "";
                    }
                }}>
                    <CheckmarkIcon></CheckmarkIcon>
                    Save
                </button>
                <button class="alt" onclick={() => {
                    showSetNewPasswordMenu = false;
                    setNewPasswordUserIndex = -1;
                    setNewPasswordPassword = "";
                }}>Cancel</button>
            </div>
            <p class="fg0"><LockIcon></LockIcon> Changing this user's password will also sign them out (on all devices)</p>
        </div>
    </div>
{/if}
