<script>
let employees = [
    {
        name: "Carson",
        wage: 200000
    },
    {
        name: "Cristian",
        wage: 200000
    },
    {
        name: "Example",
        wage: 1234
    },
]
</script>
<style>
.layoutthing {
    display: grid;
    gap: 1rem;
    grid-template-rows: auto;
    grid-template-columns: 1fr 1fr 1fr;
    padding: 1rem;
}
.layoutthing table {
    margin-top: 0px;
}
@media only screen and (max-width: 1000px) {
    .layoutthing {
        grid-template-rwos: auto auto auto;
        grid-template-columns: auto;
    }
}
.centerthisplz {
    text-align: center;
}
</style>
<div style="padding: 1rem;">
<h2 class="centerthisplz">Run Day</h2>
<div class="layoutthing">
<table>
    <thead>
        <tr>
            <th>List of People</th>
            <th>Hours Worked</th>
        </tr>
    </thead>
    <tbody>
        {#each employees as employee}
        <tr>
            <td>{employee.name}</td>
            <td>
                <input type="text">
            </td>
        </tr>
        {/each}
    </tbody>
</table>
<table>
    <thead>
        <tr>
            <th>List of People</th>
            <th>Hours Worked</th>
        </tr>
    </thead>
    <tbody>
        {#each employees as employee}
        <tr>
            <td>{employee.name}</td>
            <td>
                <input type="text">
            </td>
        </tr>
        {/each}
    </tbody>
</table>
<table>
    <thead>
        <tr>
            <th>List of People</th>
            <th>Hours Worked</th>
        </tr>
    </thead>
    <tbody>
        {#each employees as employee}
        <tr>
            <td>{employee.name}</td>
            <td>
                <input type="text">
            </td>
        </tr>
        {/each}
    </tbody>
</table>
</div>
<div class="flex">
    <button onclick={function () {
        fetch("/api/graphql", {
            method: "POST",
            body: JSON.stringify({
                query: `{
    runDay(
        floorHours: $floorHoursArray,
        kitchenHours: kitchenHoursArray,
        foodCost: $foodCost
    )
}`,
                variables: {
                    "floorHoursArray": ""
                }
            })
        })
    }}>Save</button>
</div>
</div>

