package com.lacipollina.budgeter;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.graphql.data.method.annotation.Argument;
import org.springframework.graphql.data.method.annotation.QueryMapping;
import org.springframework.graphql.data.method.annotation.MutationMapping;
import org.springframework.graphql.data.method.annotation.SchemaMapping;
import org.springframework.stereotype.Controller;

import graphql.schema.DataFetchingEnvironment;

import java.util.List;
import java.util.ArrayList;
import java.time.LocalDate;

import com.lacipollina.budgeter.repos.EmployeeRepo;
import com.lacipollina.budgeter.repos.DayInputRepo;

@Controller
public class ApiController {
    @Autowired
    private EmployeeRepo employeeRepo;
    @Autowired
    private DayInputRepo dayInputRepo;
    @Autowired
    private DayReportRepo dayReportRepo;

    @QueryMapping
    public List<Employee> allEmployees() {
        return employeRepo.getAllEmployees();
    }

    @MutationMapping
    public boolean runDay(
        @Argument LocalDate id,
        @Argument List<HoursInput> floorHours,
        @Argument List<HoursInput> kitchenHours,
        @Argument List<Double> foodCosts
    ) {
        dayInputRepo.create(
            id,
            floorHours,
            kitchenHours,
            foodCost
        );
        double floorBudget = 0.0;
        double kitchenBudget = 0.0;
        double kitchenBudget = 0.0;
        for (HoursInput floorHoursInput : floorHours) {
            floorBudget += floorHoursInput.getHours() * employeeRepo.getEmployeeByName(
                floorHoursInput.getName()
            ).getWage();
        }
        for (HoursInput kitchenHoursInput : kitchenHours) {
            kitchenBudget += kitchenHoursInput.getHours() * employeeRepo.getEmployeeByName(
                kitchenHoursInput.getName()
            ).getWage();
        }
    }
}
