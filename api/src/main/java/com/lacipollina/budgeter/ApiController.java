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

    ArrayList<Double> dayFinals = new ArrayList<Double>();
            //Holds Initial and day finals for all budget^^^
            ArrayList<Double> weekIncreases = new ArrayList<Double>();
            //Holds all budget Increases(0 if none)^^^
            ArrayList<Double> weekExpenses = new ArrayList<Double>();
        //Holds all expenses from the budget^^^(0 if none)
    public void calculateDay()

    @MutationMapping
    public boolean runDay(
        @Argument List<HoursInput> floorHours,
        @Argument List<HoursInput> kitchenHours,
        @Argument List<float> foodCost
    ) {
        currentWeekRepo.saveData(
            floorHours,
            kitchenHours,
            foodCost
        )
        employeeRepo.getEmployeeByName(
            floorHours.get(0).getName()
        ).getWage() * floorHours.get(0).getHours()
        @Argument float foodCostChange
    ) 
        float totalPay=0.0;

        for(int i=0;i<floorHours.size();i++)
    {
        float floorPay= employeeRepo.getEmployeeByName(floorHours.
        get(i).getName()).getWage()*floorHours.get(i).getHours();
        totalPay+=floorPay;
    }
        for (int j=0;j<kitchenHours.size();j++)
    {
        float kitchenPay=employeeRepo.getEmployeeByName(kitchenHours.
        get(jk).getName()).getWage()*kitchenHours.get(j).getHours();
        totalPay+=floorPay;
    }
        for (int k=0;k<foodCost.size();k++)
    {
        totalPay+=foodCost.get(k);
    }
    
}
