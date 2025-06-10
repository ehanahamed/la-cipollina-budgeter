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
    <img src="/logo.png" alt="La Cipollina">
</div>
<p class="h4 header-text">Budgeter</p>
<div class="flex center">
    <div class="flex center" style="flex-direction: column;">
        <div class="combo-buttons">
            <a class="button left" href="/edit-budget">Edit Budget</a>
            <a class="button mid" href="/run-day">Run Day</a>
            <button class="right">Edit Day</button>
        </div>
        <a href="/reports" class="button">Reports</a>
        <a href="/settings" class="button">Settings</a>
    </div>
</div>

