<script>
    import { onMount } from "svelte";
    import { base } from "$app/paths";
    let { data } = $props();
    let authed = $state(false);
    let enteredUsername = $state("");
    let enteredPassword = $state("");
    onMount(function () {
        if (window.localStorage) {
            if (localStorage.getItem("budgeter:auth")) {
                authed = true;
            }
        }
    })
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
</script>
<style>
    .header-text {
        text-align: center;
        margin-top: 1rem;
    }
    .img-container {
        display: flex;
        margin-top: 4rem;
        width: 100%;
        justify-items: center;
        justify-content: center;
    }
</style>
<div class="img-container">
    <img src="{base}/logo.png" alt="La Cipollina">
</div>
<p class="h4 header-text">Budgeter</p>
{#if authed}
<div class="flex center">
    <div class="flex center no-child-clickable-effect" style="flex-direction: column;">
        <div class="combo-buttons">
            <a class="button left" href="{base}/edit-budget">Edit Budget</a>
            <a class="button mid" href="{base}/run-day">Run Day</a>
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

