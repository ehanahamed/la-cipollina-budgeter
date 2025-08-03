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
import XMarkIcon from "$lib/icons/CloseXMark.svelte";
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
<h4 style="margin-bottom: 0px; text-align: center;">Settings</h4>
                <div style="margin-bottom: 2rem;"><a href={base} class="button faint">
                    <BackArrowIcon></BackArrowIcon>
                    Back
                </a></div>
<p>Show week numbers?</p>
<div class="combo-select" style="margin-bottom: 4rem">
    <button class="left {
        localStorage.getItem("budgeter:showWeekNums") == "true" ?
            "" : "selected"
    }" onclick={() => {
        localStorage.setItem("budgeter:showWeekNums", "false");
        window.location.reload();
    }}>
        <CheckmarkIcon class="combo-selected-icon"></CheckmarkIcon>
        Hide
    </button>
    <button class="right {
        localStorage.getItem("budgeter:showWeekNums") == "true" ?
            "selected" : ""
    }" onclick={() => {
        localStorage.setItem("budgeter:showWeekNums", "true");
        window.location.reload();
    }}>
        <CheckmarkIcon class="combo-selected-icon"></CheckmarkIcon>
        Show
    </button>
</div>
<div style="margin-bottom: 2rem;">
    <p>Currently logged in as "{data?.authedUser?.username}"</p>
    <p class="fg0" style="margin-top: 0.4rem;">Role: {
        data?.authedUser?.admin ?
            "admin user" : "normal user"
    }</p>
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
{#if data.authedUser.admin}
<p>Add/Edit Users</p>
{:else}
<p>Users:</p>
{/if}
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
                        {#if data.authedUser.admin || user.id == data.authedUser.id}
                        <div class="dropdown" tabindex="0">
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
                                {#if data.authedUser.admin && user.admin &&
                                user.id != data.authedUser.id}
                                <button class="ohno" onclick={async () => {
                                    try {
                                        const res = await fetch(
                                            `${data.PUBLIC_API_URL}/users/${user.id}`,
                                            {
                                                method: "PATCH",
                                                headers: {
                                                    "Authorization": `Bearer ${
                                                        localStorage.getItem("budgeter:auth")
                                                    }`,
                                                    "Content-Type": "application/json"
                                                },
                                                body: JSON.stringify({
                                                    admin: false
                                                })
                                            }
                                        );
                                        const resJson = await res.json();
                                        if (resJson.id) {
                                            window.location.reload();
                                        } else {
                                            console.log(resJson);
                                            alert("it seems to have not worked")
                                        }
                                    } catch (err) {
                                        console.error(err);
                                        alert("something didn't work")
                                    }
                                }}>
                                    <XMarkIcon></XMarkIcon>
                                    Remove Admin
                                </button>
                                {:else if data.authedUser.admin &&
                                user.id != data.authedUser.id}
                                <button onclick={async () => {
                                    try {
                                        const res = await fetch(
                                            `${data.PUBLIC_API_URL}/users/${user.id}`,
                                            {
                                                method: "PATCH",
                                                headers: {
                                                    "Authorization": `Bearer ${
                                                        localStorage.getItem("budgeter:auth")
                                                    }`,
                                                    "Content-Type": "application/json"
                                                },
                                                body: JSON.stringify({
                                                    admin: true
                                                })
                                            }
                                        );
                                        const resJson = await res.json();
                                        if (resJson.id) {
                                            window.location.reload();
                                        } else {
                                            console.log(resJson);
                                            alert("it seems to have not worked")
                                        }
                                    } catch (err) {
                                        console.error(err);
                                        alert("something didn't work")
                                    }
                                }}>
                                    <CheckmarkIcon></CheckmarkIcon>
                                    Give Admin
                                </button>
                                {/if}
                                {#if data.authedUser.admin}
                                <button class="ohno" onclick={() => {
                                    userToDeleteIndex = userIndex;
                                    showDeleteUserConfirmation = true;
                                }}>
                                    <TrashIcon></TrashIcon>
                                    Delete
                                </button>
                                {/if}
                            </div>
                        </div>
                        {/if}
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
        {:else if data.authedUser.admin}
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
{#if data.authedUser.admin}
<p class="fg0">
    Admin users can add/edit other users.<br>
    Normal users can only change their own username and password.
</p>
{:else}
<p class="fg0">
    Only admin users can add/edit other users.<br>
    Normal users can only change their own username and password.
</p>
{/if}
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
            {#if users[setNewPasswordUserIndex].id == data.authedUser.id}
            <p class="fg0"><LockIcon></LockIcon> Changing your password will also sign you out (on all devices)</p>
            {:else}
            <p class="fg0"><LockIcon></LockIcon> Changing this user's password will also sign them out (on all devices)</p>
            {/if}
        </div>
    </div>
{/if}
