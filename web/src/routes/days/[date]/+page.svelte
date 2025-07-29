<script>
import { base } from "$app/paths";
import BackArrowIcon from "$lib/icons/BackArrow.svelte";
import PlusIcon from "$lib/icons/Plus.svelte";
import CheckmarkIcon from "$lib/icons/Checkmark.svelte";
import XMarkIcon from "$lib/icons/CloseXMark.svelte";
import { onMount } from "svelte";
let { data } = $props();
let floorHourlyArray = $state([]);
let floorSpecialArray = $state([]);
let kitchenHourlyArray = $state([]);
let kitchenSpecialArray = $state([]);
let floorEmployeesCount = $state(0);
let kitchenEmployeesCount = $state(0);
const date = (() => {
    const [year, month, day] = data.date.split("-");
    const d = new Date(year, month - 1, day);
    if (!isNaN(d)) {
        return d;
    } else {
        window.location = `${base}/days/${data.date}/err/invalid-date`
    }
})();
function dateToYMDString(dateObj) {
    return dateObj.getFullYear() + "-" +
        (dateObj.getMonth() + 1).toString().padStart(2, "0") + "-" +
        dateObj.getDate().toString().padStart(2, "0")
}
const dateYMD = dateToYMDString(date);
const weekDayName = [
    "Sunday", "Monday", "Tuesday", "Wednesday",
    "Thrusday", "Friday", "Saturday"
][date.getDay()];
const weekDayKey = [
    "mon", "tue", "wed", "thu",
    "fri", "sat", "sun"
][date.getDay()];
const monthName = [
    "January", "February", "March", "April", "May",
    "June", "July", "August", "September",
    "October", "November", "December"
][date.getDay()];
let foodCostsArray = $state([
    {
        cost: "",
        notes: ""
    }
]);
// let foodCostsTotal = $derived.by(() => {
//     let total = 0;
//     foodCostsArray.forEach((row) => {
//         if (typeof row.cost === "string") {
//             const cost = parseFloat(row.cost.replaceAll(",", ""));
//             if (!isNaN(cost)) {
//                 total += cost;
//             }
//         } else if (typeof row.cost === "number") {
//             total += row.cost;
//         }
//     })
//     return total;
// })
let weekData = $state(null);
let showWeekBudget = $state(false);
let startingNewWeek = $state(false);
let prevDaysArray = $state([]);
let selectedDay = $state(null);
const dateAsNum = parseInt(
    data.date.replaceAll("-", "")
);
data.days.forEach((day) => {
    const iterationDateAsNum = parseInt(
        day.date.replaceAll("-", "")
    );
    if (iterationDateAsNum < dateAsNum) {
        prevDaysArray.push(day);
    } else if (iterationDateAsNum == dateAsNum) {
        selectedDay = day;
    }
})
if (!selectedDay) {
    window.location = `${base}/days/${data.date}/err/dayNotFound`
}
selectedDay.hoursWorked.forEach((row) => {
    if (row.employee.specialPay) {
        if (row.employee.type == "FLOOR") {
            floorSpecialArray.push(row);
        } else if (row.employee.type == "KITCHEN") {
            kitchenSpecialArray.push(row);
        }
    }
    if (row.employee.type == "FLOOR") {
        floorEmployeesCount++;
        floorHourlyArray.push(row);
    } else if (row.employee.type == "KITCHEN") {
        kitchenEmployeesCount++;
        kitchenHourlyArray.push(row);
    }
})
selectedDay.workedToday.forEach((row) => {
    if (row.employee.type == "FLOOR") {
        let found = false;
        for (
            let index = 0;
            index < floorSpecialArray.length;
            index++
        ) {
            if (
                floorSpecialArray[index].employee.id ==
                row.employee.id
            ) {
                found = true;
                floorSpecialArray[
                    index
                ].workedToday = row.workedToday;
            }
        }
        if (!found) {
            floorSpecialArray.push(row);
        }
    } else if (row.employee.type == "KITCHEN") {
        let found = false;
        for (
            let index = 0;
            index < kitchenSpecialArray.length;
            index++
        ) {
            if (
                kitchenSpecialArray[index].employee.id ==
                row.employee.id
            ) {
                found = true;
                kitchenSpecialArray[
                    index
                ].workedToday = row.workedToday;
            }
        }
        if (!found) {
            kitchenSpecialArray.push(row);
        }
    }
})
</script>
<div class="grid page" style="margin-top: 4rem; margin-bottom: 10rem;">
    <div class="content">
        <div class="flex">
            <a class="button faint" href={base}>
                <BackArrowIcon></BackArrowIcon>
                Back
            </a>
        </div>
        <h3 style="margin-bottom: 0px;">{weekDayName}'s Report</h3>
        <p style="margin-top: 0.4rem;">{monthName} {date.getDate()}, {date.getFullYear()}</p>
        <div style="display: inline-block;">
            <div class="flex" style="justify-content: space-between;">
                <p>Kitchen Workers</p>
                <p class="fg0">
                    {kitchenEmployeesCount} employees
                </p>
            </div>
            <table>
                <thead>
                    <tr>
                        <th>Employee</th>
                        <th>Hours Worked</th>
                        <th>Amount Earned</th>
                    </tr>
                </thead>
            </table>
        </div>

        <p style="white-space: pre-wrap;">
            {JSON.stringify(selectedDay, null, 4)}
        </p>
    </div>
</div>
