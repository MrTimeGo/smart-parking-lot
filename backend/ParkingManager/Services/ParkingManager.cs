using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Options;
using ParkingManager.Configurations;
using ParkingManager.Data;
using ParkingManager.Dtos;
using ParkingManager.Entities;
using ParkingManager.Services.Interfaces;

namespace ParkingManager.Services;

public class ParkingManager(
    ParkingContext context,
    IOptions<ParkingSettings> parkingSettings
    ) : IParkingManager
{
    public async Task<List<ParkingDto>> GetParkingStatus()
    {
        return await context.ActionLogs
            .GroupBy(l => l.Place)
            .Where(g => g.Select(l => l.Action).FirstOrDefault() == ActionType.Enter)
            .Select(g => new ParkingDto()
            {
                Place = g.Key,
                PlateNumber = g.OrderByDescending(l => l.At).Select(l => l.PlateNumber).First()
            })
            .ToListAsync();
    }

    public async Task<List<ActionLogDto>> GetLastActionLogs()
    {
        return await context.ActionLogs
            .OrderByDescending(l => l.At)
            .Select(l => new ActionLogDto()
            {
                Image = l.Image,
                Place = l.Place,
                PlateNumber = l.PlateNumber,
                Cost = l.Cost,
                At = l.At,
                Action = l.Action,
            })
            .Take(50)
            .ToListAsync();
    }

    public async Task ParkCar(string plateImage)
    {
        var plateNumber = await GetPlateNumberByImageAsync(plateImage);

        var freePlace = await GetFreePlaceNumber();

        if (freePlace == 0)
        {
            return;
        }
        
        var actionLog = new ActionLog()
        {
            Action = ActionType.Enter,
            At = DateTime.UtcNow,
            Image = plateImage,
            PlateNumber = plateNumber,
            Place = freePlace
        };

        context.Add(actionLog);
        await context.SaveChangesAsync();
    }

    private async Task<int> GetFreePlaceNumber()
    {
        var busyPlaces = await context.ActionLogs
            .GroupBy(l => l.Place)
            .Where(g =>
                g.Count(l => l.Action == ActionType.Enter) != g.Count(l => l.Action == ActionType.Exit)
            )
            .Select(l => l.Key)
            .ToListAsync();

        var maxPlaces = parkingSettings.Value.ParkingPlaces;

        var freePlaces = Enumerable.Range(1, 20).Where(i => !busyPlaces.Contains(i)).ToArray();

        return freePlaces.Length == 0 ? 0 : Random.Shared.GetItems(freePlaces, 1)[0];
    }

    public async Task UnparkCar(string plateImage)
    {
        var plateNumber = await GetPlateNumberByImageAsync(plateImage);

        var enteringActionLog = await context.ActionLogs
            .OrderByDescending(l => l.At)
            .FirstOrDefaultAsync(l => l.PlateNumber == plateNumber);

        if (enteringActionLog is null || enteringActionLog.Action != ActionType.Enter)
        {
            return;
        }
        
        var actionLog = new ActionLog()
        {
            Action = ActionType.Exit,
            At = DateTime.UtcNow,
            Image = plateImage,
            PlateNumber = plateNumber,
            Place = enteringActionLog.Place,
            Cost = (DateTime.UtcNow - enteringActionLog.At).Minutes * parkingSettings.Value.CostPerMinute
        };

        context.Add(actionLog);
        await context.SaveChangesAsync();
    }

    private static async Task<string> GetPlateNumberByImageAsync(string plateImage)
    {
        await Task.Delay(500);
        return "AA1234AA";
    }
}