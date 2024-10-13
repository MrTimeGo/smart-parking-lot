namespace ParkingManager.Configurations;

public class ParkingSettings
{
    public const string SectionName = "ParkingSettings";
    
    public int ParkingPlaces { get; set; }
    public decimal CostPerMinute { get; set; }
}