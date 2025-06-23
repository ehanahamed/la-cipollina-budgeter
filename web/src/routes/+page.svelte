<script>
    import { onMount } from "svelte";
    import { base } from "$app/paths";
    let { data } = $props();
    let authed = $state(false);
    let enteredUsername = $state("");
    let enteredPassword = $state("");
    onMount(function () {
        if (window.localStorage) {
            if (localStorage.getItem("auth")) {
                authed = true;
            }
        }
    })
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
    <div class="flex center" style="flex-direction: column;">
        <div class="combo-buttons">
            <a class="button left" href="{base}/edit-budget">Edit Budget</a>
            <a class="button mid" href="{base}/run-day">Run Day</a>
            <a class="button right" href="{base}/edit-day">Edit Day</a>
        </div>
        <a href="{base}/reports" class="button">Reports</a>
        <a href="{base}/settings" class="button">Settings</a>
    </div>
</div>
{:else}
<div style="margin-top: 2rem;">
    <div class="grid thin-centered">
        <div class="content">
            <div class="separator">Login</div>
            <div>
            <input type="text" placeholder="Username" bind:value={enteredUsername}>
            </div>
            <div>
            <input type="password" placeholder="Password" bind:value={enteredPassword}>
            </div>
            <div class="flex">
                <button onclick={async function () {
                    try {
                        const res = await (
                            await fetch(data.PUBLIC_API_URL + "/log-in", {
                               method: "POST",
                               body: JSON.stringify({
                                  username: enteredUsername,
                                  password: enteredPassword
                               })
                           })
                        ).json();
                        if (res?.token) {
                            localStorage.setItem("auth", res.token);
                            authed = true;
                        } else {
                            alert("wrong password?");
                        }
                    } catch (err) {
                        console.err("Error in logging in req: ", err);
                        alert("There was an error while logging in :( try again mabye");
                    }
                }}>Continue</button>
            </div>
        </div>
    </div>
</div>
{/if}

