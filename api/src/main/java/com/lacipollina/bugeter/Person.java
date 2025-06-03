public class Person 
{
    private String name="";
    private double wage=0;
    
    public Person (String inName, double inWage) 
    {
        name=inName;
        wage=inWage;
    }
    
    //FRANCO PAY CALC
    public Person (String inName)
    {
        name=inName;
    }
    public String getName() 
    {
        return name;
    }
    public double getWage() 
    {
        return wage;
    }
}
