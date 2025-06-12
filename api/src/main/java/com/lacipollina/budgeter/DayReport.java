package com.lacipollina.budgeter;

import java.time.LocalDate;

public class DayReport {
    private LocalDate id;
    private double floorBudgetStart;
    private double kitchenBudgetStart;
    private double foodCostStart;
    private double floorBudgetFinal;
    private double kitchenBudgetFinal;
    private double foodCostFinal;

    public DayInput(
        LocalDate id,
        double floorBudgetStart,
        double kitchenBudgetStart,
        double foodCostStart,
        double floorBudgetFinal,
        double kitchenBudgetFinal,
        double foodCostFinal
    ) {
        this.floorBudgetStart = floorBudgetStart;
        this.kitchenBudgetStart = kitchenBudgetStart;
        this.foodCostStart = foodCostStart;
        this.floorBudgetFinal = floorBudgetFinal;
        this.kitchenBudgetFinal = kitchenBudgetFinal;
        this.foodCostFinal = foodCostFinal;
    }

    public LocalDate getId() {
        return id;
    }
    public double getFloorBudgetStart() {
        return getFloorBudgetStart;
    }
    public double getKitchenBudgetStart() {
        return kitchenBudgetStart;
    }
    public double getFoodCostStart() {
        return foodCostStart;
    }
    public double getFloorBudgetFinal() {
        return floorBudgetFinal;
    }
    public double getKitchenBudgetFinal() {
        return kitchenBudgetFinal;
    }
    public double getFoodCostFinal() {
        return foodCostFinal;
    }
}
