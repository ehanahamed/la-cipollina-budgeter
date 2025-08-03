<script>
    import { onMount } from "svelte";
    import { base } from "$app/paths";
    import { dateGetWeekNum } from "$lib/dateGetWeekNum.js";
    import { dateToYMDString } from "$lib/dateToYMDString.js";
    import DocIcon from "$lib/icons/DocFilePage.svelte";
    import SettingsIcon from "$lib/icons/SettingsCogGear.svelte";
    import PersonIcon from "$lib/icons/Person.svelte";
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
    const today = new Date();
    const todayYMD = dateToYMDString(today);
    const todayWeekNum = dateGetWeekNum(today);
    let newDay = $state(true);
    $effect(() => {
        if (authed) {
            (async () => {
            try {
                const res = await fetch(`${data.PUBLIC_API_URL}/days/${todayYMD}`, {
                    method: "GET",
                    headers: {
                        "Authorization": "Bearer " + localStorage.getItem("budgeter:auth")
                    }
                })
                const resJson = await res.json();
                if (resJson?.id != null) {
                    newDay = false;
                }
            } catch (err) {
                console.error(err);
            }
            })();
        }
    });
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
        <div class="flex center no-child-clickable-effect" style="flex-direction: column; min-width: 16rem;">
        {#if newDay}
            <a class="button" href="{base}/record-day">
                Record Day
            </a>
        {:else}
            <a class="button" href="{base}/days/{todayYMD}">
                View/Edit Day
            </a>
        {/if}
        <a class="button alt" href="{base}/select-day">
            View/Edit Another Day
        </a>
        <div class="separator">or</div>
        <!-- <div class="combo-buttons"> -->
        <!--     <a class="button left" href="{base}/edit-budget">Edit Budget</a> -->
        <!--     <a class="button mid" href="{base}/record-day">Record Day</a> -->
        <!--     <a class="button right" href="{base}/edit-day">Edit Day</a> -->
        <!-- </div> -->
        <a href="{base}/reports" class="button">
            <DocIcon></DocIcon>
            Reports
        </a>
        <a href="{base}/employees" class="button">
            <PersonIcon></PersonIcon>
            Employees
        </a>
        <a href="{base}/settings" class="button">
            <SettingsIcon></SettingsIcon>
            Settings
        </a>
    </div>
</div>
<div style="width: 100%; margin-top: 4rem; margin-bottom: 8rem;">
    <p class="center fg0">
        Ehan Ahamed, Carson Fusco, Cristian Hernandez <br>
        2025 · <a href="https://github.com/ehanahamed/la-cipollina-budgeter" class="text fg0" style="text-decoration: underline;">Source Code</a>
    </p>
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

