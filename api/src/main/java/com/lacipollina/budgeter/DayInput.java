package com.lacipollina.budgeter;

import java.time.LocalDate;

public class DayInput {
    private LocalDate id;
    private List<HoursInput> floorHours;
    private List<HoursInput> kitchenHours;
    private List<Double> foodCosts;

    public DayInput(
        LocalDate id,
        List<HoursInput> floorHours,
        List<HoursInput> kitchenHours,
        List<Double> foodCosts
    ) {
        this.id = id;
        this.floorHours = floorHours;
        this.kitchenHours = kitchenHours;
        this.foodCosts = foodCosts;
    }

    public LocalDate getId() {
        return id;
    }
    public List<HoursInput> getFloorHours() {
        return floorHours;
    }
    public List<HoursInput> getKitchenHours() {
        return kitchenHours;
    }
    public List<Double> getFoodCosts() {
        return foodCosts;
    }
}
