using ParkingManager.Dtos;

namespace ParkingManager.Services.Interfaces;

public interface IParkingManager
{
    Task<List<ParkingDto>> GetParkingStatus();
    Task<List<ActionLogDto>> GetLastActionLogs();

    Task ParkCar(string plateImage);
    Task UnparkCar(string plateImage);
}