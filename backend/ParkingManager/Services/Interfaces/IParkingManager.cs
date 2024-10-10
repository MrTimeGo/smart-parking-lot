namespace ParkingManager.Services.Interfaces;

public interface IParkingManager
{
    Task<object> GetParkingStatus();
    
}