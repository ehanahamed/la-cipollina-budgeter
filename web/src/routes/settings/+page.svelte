<script>
let employees = $state([
    {
        name: "Carson",
        type: "floor",
        wage: 200
    },
    {
        name: "Cristian",
        type: "floor",
        wage: 200
    },
    {
        name: "Example",
        type: "kitchen",
        wage: 200
    }
]);
</script>
Add/Edit Employees
<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Wage</th>
        </tr>
    </thead>
    <tbody>
        {#each employees as employee}
            <tr>
                <td>{employee.name}</td>
                <td>{employee.type}</td>
            </tr>
        {/each}
        <tr>
        </tr>
    </tbody>
</table>

