package com.lacipollina.budgeter;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.graphql.data.method.annotation.Argument;
import org.springframework.graphql.data.method.annotation.QueryMapping;
import org.springframework.graphql.data.method.annotation.MutationMapping;
import org.springframework.graphql.data.method.annotation.SchemaMapping;
import org.springframework.stereotype.Controller;

import graphql.schema.DataFetchingEnvironment;

import java.util.List;

import com.lacipollina.budgeter.repos.EmployeeRepo;

@Controller
public class ApiController {
    @Autowired
    private EmployeeRepo employeRepo;

    @QueryMapping
    public List<Employee> allEmployees() {
        return employeRepo.getAllEmployees();
    }

    @MutationMapping
    public boolean runDay(
        @Argument List<HoursInput> floorHours,
        @Argument List<HoursInput> kitchenHours,
        @Argument float foodCostChange
    ) {
        
    }
}
