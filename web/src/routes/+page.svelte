<script>
    import { onMount } from "svelte";
    import { base } from "$app/paths";
    import { dateGetWeekNum } from "$lib/dateGetWeekNum.js";
    let { data } = $props();
    let authed = $state(data.authed);
    let enteredUsername = $state("");
    let enteredPassword = $state("");
    async function loginButton() {
        try {
            const res = await (
                await fetch(data.PUBLIC_API_URL + "/log-in", {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json"
                    },
                    body: JSON.stringify({
                       username: enteredUsername,
                       password: enteredPassword
                    })
               })
            ).json();
            if (res?.token) {
                localStorage.setItem("budgeter:auth", res.token);
                authed = true;
            } else {
                console.log(res);
                alert("wrong password?");
            }
        } catch (err) {
            console.error("Error in log-in req: ", err);
            alert("There was an error while logging in :( try again mabye");
        }
    }
    let passwordTextbox;
    let today = new Date();
    let todayWeekNum = dateGetWeekNum(today);
</script>
<style>
    .header-text {
        text-align: center;
        margin-top: 0.2rem;
        font-weight: normal;
        font-size: 1.4rem;
    }
    .img-container {
        display: flex;
        margin-top: 4rem;
        width: 100%;
        justify-items: center;
        justify-content: center;
    }
    .img-container img {
        height: 4rem;
    }
</style>
<div class="img-container">
    <img src="{base}/logo.png" alt="La Cipollina">
</div>
<p class="header-text">Budgeter</p>
{#if authed}
<div class="flex center underheadertext" style="flex-direction: column; margin-top: 4rem; gap: 0px;">
    <span style="margin-bottom: 0px;">Today is</span>
    <span class="h3" style="margin-top: 0px; margin-bottom: 0px;">{["Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"][today.getDay()]}</span>
    <span>{["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"][today.getMonth()]} {
        today.getDate()
    }, {today.getFullYear()}{
        localStorage.getItem("budgeter:showWeekNums") == "true" ?
            `, W${todayWeekNum}` : ""
    }</span>
</div>
<div class="flex center">
    <div class="flex center no-child-clickable-effect" style="flex-direction: column;">
        <div class="combo-buttons">
            <a class="button left" href="{base}/edit-budget">Edit Budget</a>
            <a class="button mid" href="{base}/record-day">Record Day</a>
            <a class="button right" href="{base}/edit-day">Edit Day</a>
        </div>
        <a href="{base}/reports" class="button">Reports</a>
        <a href="{base}/employees" class="button">Employees</a>
        <a href="{base}/settings" class="button">Settings</a>
    </div>
</div>
{:else}
<div style="margin-top: 2rem;">
    <div class="grid thin-centered">
        <div class="content">
            <div class="separator">Login</div>
            <div>
            <input type="text" placeholder="Username" bind:value={enteredUsername} onkeydown={function (event) {
                if (event.key === "Enter") {
                    passwordTextbox.focus();
                }
            }}>
            </div>
            <div>
            <input type="password" placeholder="Password" bind:this={passwordTextbox} bind:value={enteredPassword} onkeydown={function (event) {
                if (event.key === "Enter") {
                    /* hitting enter in the password textbox makes it try to log in just like hitting the login button */
                    loginButton();
                }
            }}>
            </div>
            <div class="flex">
                <button onclick={loginButton}>Continue</button>
            </div>
        </div>
    </div>
</div>
{/if}

