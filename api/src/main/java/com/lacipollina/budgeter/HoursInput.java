public class HoursInput {
    private String name;
    private float hours;

    public HoursInput(
        String name,
        float hours
    ) {
        this.name = name;
        this.hours = hours;
    }

    public String getName() {
        return name;
    }
    public float getHours() {
        return hours;
    }

    public void setName(String newName) {
        name = newName;
    }
    public void setHours(String newHours) {
        hours = newHours;
    }
}
