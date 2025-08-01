export function calculateDay(
    foodBudgetStart,
    kitchenBudgetStart,
    floorBudgetStart,
    day,
    weeklyPay
) {
    let foodBudget = foodBudgetStart;
    let kitchenBudget = kitchenBudgetStart;
    let floorBudget = floorBudgetStart;
    foodBudget += day.foodBudgetIncrease ?? 0;
    kitchenBudget += day.kitchenBudgetIncrease ?? 0;
    floorBudget += day.floorBudgetIncrease ?? 0;

    const date = (() => {
        const [y, m, d] = day.date.split("-");
        return new Date(y, m - 1, d);
    })();
    const weekDayKey = [
        "sun", "mon", "tue", "wed",
        "thu", "fri", "sat",
    ][date.getDay()];

    let foodCostsTotal = 0;
    let kitchenHourlyArray = [];
    let kitchenSpecialArray = [];
    let floorHourlyArray = [];
    let floorSpecialArray = [];

    day?.foodCosts?.forEach((row) => {
        foodCostsTotal += row.cost;
    })

    day?.hoursWorked?.forEach((row) => {
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
    day?.workedToday?.forEach((row) => {
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
    let totalKitchenHourlyEarned = 0;
    let totalKitchenSpecialEarned = 0;
    let totalFloorHourlyEarned = 0;
    let totalFloorSpecialEarned = 0;
    kitchenHourlyArray?.forEach((row) => {
        totalKitchenHourlyEarned += (
            row.hours * row.employee.wage
        );
    })
    floorHourlyArray?.forEach((row) => {
        totalFloorHourlyEarned += (
            row.hours * row.employee.wage
        );
    })
    for (
        let index = 0;
        index < kitchenSpecialArray.length;
        index++
    ) {
        const row = kitchenSpecialArray[index];
        let earned = 0;
        if (
            row.workedToday &&
            row?.employee?.specialPay?.[
                weekDayKey
            ]?.perDay != null
        ) {
            earned += row.employee.specialPay[
                weekDayKey
            ].perDay;
        }
        if (
            row.hours != null &&
            row?.employee?.specialPay?.[
                weekDayKey
            ]?.perHour != null
        ) {
            earned += row.hours * row.employee.specialPay[
                weekDayKey
            ].perHour;
        }
        kitchenSpecialArray[index].earned = earned;
        totalKitchenSpecialEarned += earned;
    }
    for (
        let index = 0;
        index < floorSpecialArray.length;
        index++
    ) {
        const row = floorSpecialArray[index];
        let earned = 0;
        if (
            row.workedToday &&
            row?.employee?.specialPay?.[
                weekDayKey
            ]?.perDay != null
        ) {
            earned += row.employee.specialPay[
                weekDayKey
            ].perDay;
        }
        if (
            row.hours != null &&
            row?.employee?.specialPay?.[
                weekDayKey
            ]?.perHour != null
        ) {
            earned += row.hours * row.employee.specialPay[
                weekDayKey
            ].perHour;
        }
        floorSpecialArray[index].earned = earned;
        totalFloorSpecialEarned += earned;
    }

    let totalKitchenWeeklyEarned = 0;
    let totalFloorWeeklyEarned = 0;
    weeklyPay?.forEach((employee) => {
        if (employee.type == "KITCHEN") {
            totalKitchenWeeklyEarned += employee.weeklyPay
        } else if (employee.type == "FLOOR") {
            totalFloorWeeklyEarned += employee.weeklyPay
        }
    })

    let totalKitchenEarned = totalKitchenHourlyEarned + totalKitchenSpecialEarned + totalKitchenWeeklyEarned;
    let totalFloorEarned = totalFloorHourlyEarned + totalFloorSpecialEarned + totalFloorWeeklyEarned;
    foodBudget -= foodCostsTotal;
    kitchenBudget -= totalKitchenEarned;
    floorBudget -= totalFloorEarned;

    return {
        foodBudgetStart,
        kitchenBudgetStart,
        floorBudgetStart,
        foodBudgetIncrease: day.foodBudgetIncrease ?? 0,
        kitchenBudgetIncrease: day.kitchenBudgetIncrease ?? 0,
        floorBudgetIncrease: day.floorBudgetIncrease ?? 0,
        foodExpenses: foodCostsTotal,
        kitchenExpenses: totalKitchenEarned,
        floorExpenses: totalFloorEarned,
        foodBudgetFinal: foodBudget,
        kitchenBudgetFinal: kitchenBudget,
        floorBudgetFinal: floorBudget,
        kitchenHourly: kitchenHourlyArray,
        kitchenSpecialPay: kitchenSpecialArray,
        floorHourly: floorHourlyArray,
        floorSpecialPay: floorSpecialArray,
        totalKitchenHourlyEarned,
        totalKitchenSpecialEarned,
        totalFloorHourlyEarned,
        totalFloorSpecialEarned,
        totalKitchenWeeklyEarned,
        totalFloorWeeklyEarned,
        date
    };
}

export function remainingBudgetFromDays(
    startFoodBudget,
    startKitchenBudget,
    startFloorBudget,
    days,
    weeklyPay
) {
    let foodBudget = startFoodBudget;
    let kitchenBudget = startKitchenBudget;
    let floorBudget = startFloorBudget;
    let dayNum = 1;
    days.forEach((day) => {
        let result = calculateDay(
            foodBudget,
            kitchenBudget,
            floorBudget,
            day,
            dayNum == 1 && weeklyPay?.length >= 1 ?
                weeklyPay : null
        );
        foodBudget = result.foodBudgetFinal;
        kitchenBudget = result.kitchenBudgetFinal;
        floorBudget = result.floorBudgetFinal;
        dayNum++;
    });
    return {
        foodBudget,
        kitchenBudget,
        floorBudget
    };
}

