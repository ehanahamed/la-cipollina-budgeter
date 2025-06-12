package com.lacipollina.budgeter;

public class Employee 
{
    private String name = "";
    private double wage = 0.0;
    private String type = "";
    
    public Employee (String name, double wage, String type) 
    {
        this.name = name;
        this.wage = wage;
        this.type = type;
    }
    public Employee (String name, String type)
    {
        this.name = name;
        this.type = type;
    }

    public String getName() 
    {
        return name;
    }
    public double getWage() 
    {
        return wage;
    }
    public String getType() 
    {
        return type;
    }
    public void setName(String newName) {
        name = newName;
    }
    public void setWage(double newWage) {
        wage = newWage;
    }
    public void setType(String newType) {
        type = newType;
    }
}
