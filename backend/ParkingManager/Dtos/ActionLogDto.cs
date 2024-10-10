using ParkingManager.Entities;

namespace ParkingManager.Dtos;

public class ActionLogDto
{
    public string Image { get; set; } = string.Empty;
    public string PlateNumber { get; set; } = string.Empty;
    public ActionType Action { get; set; }
    public DateTime At { get; set; }
    public int Place { get; set; }
    public decimal? Cost { get; set; }
}