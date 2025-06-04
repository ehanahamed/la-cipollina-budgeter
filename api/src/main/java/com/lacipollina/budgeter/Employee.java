public class Employee 
{
    private String name = "";
    private double wage = 0.0;
    
    public Employee (String name, double wage) 
    {
        this.name = name;
        this.wage = wage;
    }
    public Employee (String name)
    {
        this.name = name;
    }

    public String getName() 
    {
        return name;
    }
    public double getWage() 
    {
        return wage;
    }
    public void setName(String newName) {
        name = newName;
    }
    public void setWage(double newWage) {
        wage = newWage;
    }
}
