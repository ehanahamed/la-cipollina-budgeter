<script>
    import { base } from "$app/paths";
    import { PUBLIC_API_URL } from "$env/static/public";
    let authed = $state(false);
    let enteredPassword = $state("");
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
    @media only screen and (max-width: 800px) {
        .header-text {
            font-size: 0.8rem;
        }
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
            <input type="password" placeholder="Enter password" bind:value={enteredPassword}>
            <div class="flex">
                <button onclick={function () {
                    fetch(PUBLIC_API_URL + "/graphql", {
                        method: "POST",
                        body: JSON.stringify({
                            query: `query signIn($password: String!){
    signIn(password: $password)
}`,
                            variables: {
                                password: enteredPassword
                            }
                        })
                    }).then(function (result) {
                        result.json().then(function (resultjson) {
                            if (resultjson?.data?.signIn) {
                                localStorage.setItem("auth", resultjson.data.signIn)
                            } else {
                                alert("wrong password?")
                            }
                        }).catch(function (error) {
                            console.error("Error in signIn req json", error);

                        })
                    }).catch(function (error) {
                        console.error("Error in signIn req", error);
                    });
                }}>Continue</button>
            </div>
        </div>
    </div>
</div>
{/if}

