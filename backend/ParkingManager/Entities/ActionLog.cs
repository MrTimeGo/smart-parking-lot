namespace ParkingManager.Entities;

public class ActionLog
{
    public int Id { get; set; }
    public ActionType Action { get; set; }
    public string PlateNumber { get; set; } = string.Empty;
    public string Image { get; set; } = string.Empty;
    public DateTime At { get; set; }
    public int Place { get; set; }
}