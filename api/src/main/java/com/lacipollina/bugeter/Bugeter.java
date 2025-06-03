import java.util.Scanner;
import java.util.ArrayList;
public class Budgeter
{
    private double foodCost;
    private double kitchenPay;
    private double floorPay;
    private double valsPay=1500.0;
    
    private Person[] floorWorkers = {new Person("Franco"), 
    new Person("Benji",20),new Person("Antonio",20), 
    new Person("Carson",15), new Person("Jared",15),
    new Person("Caitlyn",15), new Person("Sophia",15),
    new Person("Emily",15), new Person("Faith",15),
    new Person("Liz",14), new Person("Amina",13),
    new Person("Valente",13), new Person("Eva",12), 
    new Person("Nemo",12), new Person("Dahlia",11), 
    new Person("Maisy",11), new Person("Evan",11), 
    new Person("Debbey",10), new Person("Naomi",10), 
    new Person("Jojo",10), new Person("Ryder",10), 
    new Person("Abby",10), new Person("Artem",10), 
    new Person("Hayden",10), new Person("Hailey",10)};
    
    private Person[] kitchenWorkers = {new Person("Manuel",16), 
    new Person("Renato",19), new Person("Dishwasher Total",15)};
    
    public Budgeter(double f, double k, double p)
    {
        foodCost = f;
        kitchenPay = k;
        floorPay = p;
    }
    public void useFoodCost(double f)
    {
        this.foodCost -= f;
    }
    public void useKitchenPay(double k)
    {
        this.kitchenPay -= k;
    }
    public void useFloorPay(double p)
    {
        this.floorPay -= p;
    }
    public void addFoodCost(double f)
    {
        this.foodCost +=f;
    }
    public void addKitchenPay(double k)
    {
        this.kitchenPay +=k;
    }
    public void addFloorPay(double p)
    {
        this.floorPay +=p;
    }
    public void setFoodCost(double f)
    {
        this.foodCost=f;
    }
    public void setKitchenPay(double k)
    {
        this.kitchenPay=k;
    }
    public void setFloorPay(double p)
    {
        this.floorPay=p;
    }
    public double getFoodCost() 
    {
      return this.foodCost;  
    }
    public double getKitchenPay()
    {
       return this.kitchenPay; 
    }
    public double getFloorPay()
    {
        return this.floorPay;
    }
    public Person[] getFloorWorkers() 
    {
        return floorWorkers;
    }
    public Person[] getKitchenWorkers()
    {
        return kitchenWorkers;
    }
    public double getValsPay()
    {
        return this.valsPay;
    }
   
    
    public static void main(String[] args)
    {
        Scanner scanner= new Scanner (System.in);
        double fdInput = 0.0;
        double kpInput = 0.0;
        double fpInput = 0.0;
        double fInc = 0.0;
        double kpInc = 0.0;
        double fpInc = 0.0;
        boolean frWork = false;
        String frInput = "";
        String reset = "";
        String foodExpInput = "";
        String change = "";
        String startDay = "";
        int startDayVal;
        double foodExpense=0.0;
        double kpExp=0.0;
        double fpExp=0.0;
        ArrayList<Double> dayFinals = new ArrayList<Double>();
        //Holds Initial and day finals for all budget^^^
        ArrayList<Double> weekIncreases = new ArrayList<Double>();
        //Holds all budget Increases(0 if none)^^^
        ArrayList<Double> weekExpenses = new ArrayList<Double>();
        //Holds all expenses from the budget^^^(0 if none)
        System.out.println("LA CIPOLLINA BUDGET CALC:");
//Ask if open Monday Tuesday so we can auto skip
//------------------------------------------------------
        System.out.print("What is the budget's start day? (mon,tue,wed): ");
        startDay = scanner.next();
        if(startDay.equals("wed"))
        {
            startDayVal = 3;
            for(int w = 0; w < 6; w++)
            {
                weekIncreases.add(0.0);
                weekExpenses.add(0.0);
                dayFinals.add(0.0);
            }
        }
        else if(startDay.equals("tue"))
        {
            startDayVal = 2;
            for(int q = 0; q < 3; q++)
            {
                weekIncreases.add(0.0);
                weekExpenses.add(0.0);
                dayFinals.add(0.0);
            }
        }
        else
        {
            startDayVal = 1;
        }
        System.out.print("Enter Food Cost Budget: ");
        fdInput= scanner.nextDouble();
        System.out.print("Enter Kitchen Pay: ");
        kpInput= scanner.nextDouble();
        System.out.print("Enter Floor Pay: ");
        fpInput= scanner.nextDouble();
        Budgeter thisWeek = new Budgeter(fdInput, kpInput, fpInput);
        dayFinals.add(fdInput);
        dayFinals.add(kpInput);
        dayFinals.add(fpInput);
        for(int i = startDayVal; i <= 7; i++)
        {
            double tempFoodCost = thisWeek.getFoodCost();
            double tempKitchenPay = thisWeek.getKitchenPay();
            double tempFloorPay = thisWeek.getFloorPay();
            
//DAY HEADERS
//------------------------------------------------------
            if(i == 1)
            {
                System.out.println("===========================");
                System.out.println("   _    _   ___            ");
                System.out.println("   |\\  /|  /   \\ |   |     ");
                System.out.println("   | \\/ |  |   | |\\  |     ");
                System.out.println("   |    |  |   | | \\ |     ");
                System.out.println("   |    |  |   | |  \\|     ");
                System.out.println("   |    |  \\___/ |   |     ");
                System.out.println("===========================");
            }
            if(i == 2)
            {
                System.out.println("===========================");
                System.out.println("   _____       _____       ");
                System.out.println("     |   |   | |           ");
                System.out.println("     |   |   | |           ");
                System.out.println("     |   |   | |---        ");
                System.out.println("     |   |   | |           ");
                System.out.println("     |   \\__/  |____       ");
                System.out.println("===========================");
            }
            if(i == 3)
            {
                System.out.println("===========================");
                System.out.println("            ____ ___       ");
                System.out.println("    |     | |    |  \\      ");
                System.out.println("    |     | |    |   |     ");
                System.out.println("    |  ^  | |--  |   |     ");
                System.out.println("    |  |  | |    |   |     ");
                System.out.println("    \\_/ \\_/ |___ |__/      ");
                System.out.println("===========================");
            }
            if(i == 4)
            {
                System.out.println("===========================");
                System.out.println("    _____                  ");
                System.out.println("      |   |   | |   |      ");
                System.out.println("      |   |   | |   |      ");
                System.out.println("      |   |---| |   |      ");
                System.out.println("      |   |   | |   |      ");
                System.out.println("      |   |   | \\__/       ");
                System.out.println("===========================");
            }
            if(i == 5)
            {
                System.out.println("===========================");
                System.out.println("    _____ ____  _____      ");
                System.out.println("    |     |   \\   |        ");
                System.out.println("    |     |   /   |        ");
                System.out.println("    |--   |__/    |        ");
                System.out.println("    |     |  \\    |        ");
                System.out.println("    |     |   | __|__      ");
                System.out.println("===========================");
            }
            if(i == 6)
            {
                System.out.println("===========================");
                System.out.println("       _    _   _____      ");
                System.out.println("      / \\  / \\    |        ");
                System.out.println("      \\   |   |   |        ");
                System.out.println("       \\  |___|   |        ");
                System.out.println("        \\ |   |   |        ");
                System.out.println("      \\_/ |   |   |        ");
                System.out.println("===========================");
            }
            if(i == 7)
            {
                System.out.println("===========================");
                System.out.println("       _                   ");
                System.out.println("      / \\ |   | |   |      ");
                System.out.println("      \\   |   | |\\  |      ");
                System.out.println("       \\  |   | | \\ |      ");
                System.out.println("        \\ |   | |  \\|      ");
                System.out.println("      \\_/ \\__/  |   |      ");
                System.out.println("===========================");
            }
//INCREASE BUDGET
//------------------------------------------------------
            System.out.print("Do you want to increase any budget? (y/n): ");
            change= scanner.next();
            if (change.equals("y")||change.equals("yes"))
            {
                System.out.print("Increase for Food Cost Budget: ");
                fInc=scanner.nextDouble();
                
                System.out.print("Increase for Kitchen Pay: ");
                kpInc=scanner.nextDouble();
                
                System.out.print("Increase for Floor Pay: ");
                fpInc=scanner.nextDouble();
                
                thisWeek.addFoodCost(fInc);
                thisWeek.addKitchenPay(kpInc);
                thisWeek.addFloorPay(fpInc);
            }
            else 
            {
                fInc=0.0;
                kpInc=0.0;
                fpInc=0.0;
            }
//FOOD COST CALC
//------------------------------------------------------
            System.out.println();
            System.out.println("---------------------");
            System.out.println("------FOOD COST------");
            System.out.println("---------------------");
            System.out.println();
            System.out.print("Any Food Expenses Today(y/n): ");
            foodExpInput = scanner.next();
            if (foodExpInput.equals("y")||foodExpInput.equals("yes"))
            {
                System.out.print("Enter Total Used On Food: ");
                foodExpense = scanner.nextDouble();
                thisWeek.useFoodCost(foodExpense);
                System.out.println("$" + foodExpense +" Taken Away From Food Budget.");
            }
            else 
            {
                System.out.println("NO EXPENSES TODAY");
            }
//KITCHEN PAY CALC
//------------------------------------------------------
            System.out.println();
            System.out.println("---------------------");
            System.out.println("-----KITCHEN PAY-----");
            System.out.println("---------------------");
            System.out.println();
            for(int k = 0; k < thisWeek.getKitchenWorkers().length; k++)
            {
                if(i == startDayVal && k == 0)
                {
                    System.out.println("Valentino Got Payed $" + thisWeek.getValsPay() + " for this week");
                    thisWeek.useKitchenPay(thisWeek.getValsPay());
                    kpExp+=thisWeek.getValsPay();
                }
                Person[] tem2=thisWeek.getKitchenWorkers();
                Person obj =tem2[k];
                System.out.print(obj.getName() + ": ");
                double timeInput = scanner.nextDouble();
                thisWeek.useKitchenPay(timeInput * obj.getWage());
                if (timeInput!=0)
                { 
                    System.out.println(obj.getName()+" is getting paid $"+obj.getWage()*timeInput+".");
                    kpExp+=obj.getWage()*timeInput;
                }
                else 
                {
                    System.out.println("NO PAY.");
                }
            }
//FLOOR PAY CALC         
//------------------------------------------------------
            System.out.println();
            System.out.println("---------------------");
            System.out.println("------FLOOR PAY------");
            System.out.println("---------------------");
            System.out.println();
            System.out.print("Did Franco work today?(y/n): ");
            frInput= scanner.next();
            if (frInput.equals("y")||frInput.equals("yes")) 
            {
                frWork = true;
            }
            else 
            {
                frWork=false;
                System.out.println("NO PAY.");
            }
            if(frWork)
            {
                if(i <= 4)
                {
                    thisWeek.useFloorPay(75);
                    fpExp+=75;
                    System.out.println("Franco got paid $75 for today.");
                }
                else if(i == 5)
                {
                    thisWeek.useFloorPay(125);
                    fpExp+=125;
                    System.out.println("Franco got paid $125 for today.");
                }
                else if(i == 6)
                {
                    thisWeek.useFloorPay(250);
                    fpExp+=250;
                    System.out.println("Franco got paid $250 for today.");
                }
                else if (i == 7)
                {
                    thisWeek.useFloorPay(125);
                    fpExp+=125;
                    System.out.println("Franco got paid $125 for today.");
                }
                else
                {
                    System.out.println("You messed up something because you're on the 8th day of the week. RESTART by pressing --Run--!");
                }
            }
            for(int k = 1; k < thisWeek.getFloorWorkers().length; k++)
            {
                Person[] tem=thisWeek.getFloorWorkers();
                Person obj =tem[k];
                System.out.print(obj.getName() + ": ");
                double timeInput = scanner.nextDouble();
                thisWeek.useFloorPay(timeInput * obj.getWage());
                if (timeInput!=0)
                { 
                    System.out.println(obj.getName()+" is getting paid $"+obj.getWage()*timeInput+".");
                    fpExp+=obj.getWage()*timeInput;
                }
                else 
                {
                    System.out.println("NO PAY.");
                }
            }
//VALUE RECORDING FOR WEEK REPORT
//------------------------------------------------------
            weekExpenses.add(foodExpense);
            weekExpenses.add(kpExp);
            weekExpenses.add(fpExp);
            weekIncreases.add(fInc);
            weekIncreases.add(kpInc);
            weekIncreases.add(fpInc);
            dayFinals.add(thisWeek.getFoodCost());
            dayFinals.add(thisWeek.getKitchenPay());
            dayFinals.add(thisWeek.getFloorPay());
//DAY REPORT
//------------------------------------------------------
            
            System.out.println();
            System.out.println("Todays Cycle is Done!");
            System.out.println();
            //FULL REPORT
            System.out.println("PLEASE LEAVE THIS TAB OPEN ON YOUR PHONE IF TODAY ISN'T SUNDAY!");
            System.out.println("WHEN YOU COME BACK DO NOT PRESS RUN AGAIN! JUST CONTINUE WITH THIS!");
            System.out.println();
//DAY REPORT
//------------------------------------------------------
            if (i==1)
            {
                System.out.println("------------MONDAY'S REPORT----------------");
            }
            else if (i==2)
            {
                System.out.println("-----------TUESDAY'S REPORT---------------");
            }
            else if (i==3)
            {
                System.out.println("----------WEDNESDAY'S REPORT--------------");
            }
            else if (i==4)
            {
                System.out.println("-----------THURSDAY'S REPORT---------------");
            }
            else if (i==5)
            {
                System.out.println("------------FRIDAY'S REPORT----------------");
            }
            else if (i==6)
            {
                System.out.println("-----------SATURDAY'S REPORT---------------");
            }
            else if (i==7)
            {
                System.out.println("------------SUNDAY'S REPORT----------------");
            }
            System.out.printf("%-10s %-3s %-10s %-10s %-10s%n"," "," |","Food Cost","Ktch Budg","Flr Budg");
            System.out.printf("%-10s %-4s %-10s %-10s %-10s%n","Day Start"," |"," $"+tempFoodCost," $"+tempKitchenPay," $"+tempFloorPay);
            System.out.printf("%-10s %-4s %-10s %-10s %-10s%n","Increases"," |","+$"+fInc,"+$"+kpInc,"+$"+fpInc);
            System.out.printf("%-10s %-4s %-10s %-10s %-10s%n","Expenses"," |","-$"+foodExpense,"-$"+kpExp,"-$"+fpExp);
            System.out.printf("%-10s %-4s %-10s %-10s %-10s%n","Day Final"," |"," $"+thisWeek.getFoodCost()," $"+thisWeek.getKitchenPay()," $"+thisWeek.getFloorPay());
            foodExpense=0.0;
            kpExp=0.0;
            fpExp=0.0;
//Ask if want to restart day
//------------------------------------------------------
            System.out.println();
            System.out.print("Did you mess up and want to restart the day?(y/n): ");
            reset = scanner.next();
            if(reset.equals("y") || reset.equals("yes"))
            {
                i--;
                thisWeek.setFoodCost(tempFoodCost);
                thisWeek.setKitchenPay(tempKitchenPay);
                thisWeek.setFloorPay(tempFloorPay);
                System.out.println("DAY RESETTING:");
                System.out.println("______________________");
                for(int y = 1; y <= 3; y++)
                {
                    weekExpenses.remove(weekExpenses.size()-1);
                    weekIncreases.remove(weekIncreases.size()-1);
                    dayFinals.remove(dayFinals.size()-1);
                }
            }
        }       
//FULL WEEK REPORT TO SEND TO ANTHONY
//------------------------------------------------------
        System.out.println("-------------------WEEK REPORT-------------------");
        System.out.printf("%-10s %-1s %-10s %-10s %-10s\n", "   Day:   ", "|", "Food Cost", "Kitch Budg", "Floor Budg");
        System.out.println("-------------------------------------------------");
        System.out.printf("%-10s %-1s %-10s %-10s %-10s\n", "Monday    ", "|", " $" + dayFinals.get(0), " $" + dayFinals.get(1), " $" + dayFinals.get(2));
        System.out.printf("%-10s %-1s %-10s %-10s %-10s\n", "-Increases", "|", "+$" + weekIncreases.get(0), "+$" + weekIncreases.get(1), "+$" + weekIncreases.get(2));
        System.out.printf("%-10s %-1s %-10s %-10s %-10s\n", "-Expenses ", "|", "-$" + weekExpenses.get(0), "-$" + weekExpenses.get(1), "-$" + weekExpenses.get(2));
        System.out.println();
        System.out.printf("%-10s %-1s %-10s %-10s %-10s\n", "Tuesday   ", "|", " $" + dayFinals.get(3), " $" + dayFinals.get(4), " $" + dayFinals.get(5));
        System.out.printf("%-10s %-1s %-10s %-10s %-10s\n", "-Increases", "|", "+$" + weekIncreases.get(3), "+$" + weekIncreases.get(4), "+$" + weekIncreases.get(5));
        System.out.printf("%-10s %-1s %-10s %-10s %-10s\n", "-Expenses ", "|", "-$" + weekExpenses.get(3), "-$" + weekExpenses.get(4), "-$" + weekExpenses.get(5));
        System.out.println();
        System.out.printf("%-10s %-1s %-10s %-10s %-10s\n", "Wednesday ", "|", " $" + dayFinals.get(6), " $" + dayFinals.get(7), " $" + dayFinals.get(8));
        System.out.printf("%-10s %-1s %-10s %-10s %-10s\n", "-Increases", "|", "+$" + weekIncreases.get(6), "+$" + weekIncreases.get(7), "+$" + weekIncreases.get(8));
        System.out.printf("%-10s %-1s %-10s %-10s %-10s\n", "-Expenses ", "|", "-$" + weekExpenses.get(6), "-$" + weekExpenses.get(7), "-$" + weekExpenses.get(8));
        System.out.println();
        System.out.printf("%-10s %-1s %-10s %-10s %-10s\n", "Thursday  ", "|", " $" + dayFinals.get(9), " $" + dayFinals.get(10), " $" + dayFinals.get(11));
        System.out.printf("%-10s %-1s %-10s %-10s %-10s\n", "-Increases", "|", "+$" + weekIncreases.get(9), "+$" + weekIncreases.get(10), "+$" + weekIncreases.get(11));
        System.out.printf("%-10s %-1s %-10s %-10s %-10s\n", "-Expenses ", "|", "-$" + weekExpenses.get(9), "-$" + weekExpenses.get(10), "-$" + weekExpenses.get(11));
        System.out.println();
        System.out.printf("%-10s %-1s %-10s %-10s %-10s\n", "Friday    ", "|", " $" + dayFinals.get(12), " $" + dayFinals.get(13), " $" + dayFinals.get(14));
        System.out.printf("%-10s %-1s %-10s %-10s %-10s\n", "-Increases", "|", "+$" + weekIncreases.get(12), "+$" + weekIncreases.get(13), "+$" + weekIncreases.get(14));
        System.out.printf("%-10s %-1s %-10s %-10s %-10s\n", "-Expenses ", "|", "-$" + weekExpenses.get(12), "-$" + weekExpenses.get(13), "-$" + weekExpenses.get(14));
        System.out.println();
        System.out.printf("%-10s %-1s %-10s %-10s %-10s\n", "Saturday  ", "|", " $" + dayFinals.get(15), " $" + dayFinals.get(16), " $" + dayFinals.get(17));
        System.out.printf("%-10s %-1s %-10s %-10s %-10s\n", "-Increases", "|", "+$" + weekIncreases.get(15), "+$" + weekIncreases.get(16), "+$" + weekIncreases.get(17));
        System.out.printf("%-10s %-1s %-10s %-10s %-10s\n", "-Expenses ", "|", "-$" + weekExpenses.get(15), "-$" + weekExpenses.get(16), "-$" + weekExpenses.get(17));
        System.out.println();
        System.out.printf("%-10s %-1s %-10s %-10s %-10s\n", "Sunday    ", "|", " $" + dayFinals.get(18), " $" + dayFinals.get(19), " $" + dayFinals.get(20));
        System.out.printf("%-10s %-1s %-10s %-10s %-10s\n", "-Increases", "|", "+$" + weekIncreases.get(18), "+$" + weekIncreases.get(19), "+$" + weekIncreases.get(20));
        System.out.printf("%-10s %-1s %-10s %-10s %-10s\n", "-Expenses ", "|", "-$" + weekExpenses.get(18), "-$" + weekExpenses.get(19), "-$" + weekExpenses.get(20));
        System.out.println();
        System.out.println("-------------------------------------------------");
        System.out.printf("%-10s %-1s %-10s %-10s %-10s\n", "WEEK FINAL", "|", " $" + dayFinals.get(21), " $" + dayFinals.get(22), " $" + dayFinals.get(23));
    }
}