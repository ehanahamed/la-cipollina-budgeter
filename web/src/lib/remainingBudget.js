export function remainingBudgetFromDays(
    startFoodBudget,
    startKitchenBudget,
    startFloorBudget,
    days
) {
    let foodBudget = startFoodBudget;
    let kitchenBudget = startKitchenBudget;
    let floorBudget = startFloorBudget;
    days.forEach((day) => {
        foodBudget += day.foodBudgetIncrease;
        kitchenBudget += day.kitchenBudgetIncrease;
        floorBudget += day.floorBudgetIncrease;
    
        const date = (() => {
            const [year, month, day] = day.date.split("-");
            return new Date(year, month - 1, day);
        })();
        const weekDayKey = [
            "mon", "tue", "wed", "thu",
            "fri", "sat", "sun"
        ][date.getDay()];
    
        let foodCostsTotal = 0;
        let kitchenHourlyArray = [];
        let kitchenSpecialArray = [];
        let floorHourlyArray = [];
        let floorSpecialArray = [];
    
        day.foodCosts.forEach((row) => {
            foodCostsTotal += row.cost;
        })
    
        day.hoursWorked.forEach((row) => {
            if (row.employee.specialPay) {
                if (row.employee.type == "FLOOR") {
                    floorSpecialArray.push(row);
                } else if (row.employee.type == "KITCHEN") {
                    kitchenSpecialArray.push(row);
                }
            }
            if (row.employee.type == "FLOOR") {
                floorHourlyArray.push(row);
            } else if (row.employee.type == "KITCHEN") {
                kitchenHourlyArray.push(row);
            }
        })
        day.workedToday.forEach((row) => {
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
        let totalKitchenEarned = 0;
        let totalFloorEarned = 0;
        kitchenHourlyArray.forEach((row) => {
            totalKitchenEarned += (
                row.hours * row.employee.wage
            );
        })
        floorHourlyArray.forEach((row) => {
            totalFloorEarned += (
                row.hours * row.employee.wage
            );
        })
        for (
            let index = 0;
            index < kitchenSpecialArray.length;
            index++
        ) {
            const row = kitchenSpecialArray[index];
            if (
                row.workedToday &&
                row?.employee?.specialPay?.[
                    weekDayKey
                ]?.perDay != null
            ) {
                totalKitchenEarned += row.employee.specialPay[
                    weekDayKey
                ].perDay;
            }
            if (
                row.hours != null &&
                row?.employee?.specialPay?.[
                    weekDayKey
                ]?.perHour != null
            ) {
                totalKitchenEarned += hours * row.employee.specialPay[
                    weekDayKey
                ].perDay;
            }
        }
        for (
            let index = 0;
            index < floorSpecialArray.length;
            index++
        ) {
            const row = floorSpecialArray[index];
            if (
                row.workedToday &&
                row?.employee?.specialPay?.[
                    weekDayKey
                ]?.perDay != null
            ) {
                totalFloorEarned += row.employee.specialPay[
                    weekDayKey
                ].perDay;
            }
            if (
                row.hours != null &&
                row?.employee?.specialPay?.[
                    weekDayKey
                ]?.perHour != null
            ) {
                totalFloorEarned += hours * row.employee.specialPay[
                    weekDayKey
                ].perDay;
            }
        }
    
        foodBudget -= foodCostsTotal;
        kitchenBudget -= totalKitchenEarned;
        floorBudget -= totalFloorEarned;
    })

    return {
        foodBudget,
        kitchenBudget,
        floorBudget
    };
}
