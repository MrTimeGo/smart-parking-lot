using Microsoft.AspNetCore.SignalR;
using ParkingManager.Dtos;

namespace ParkingManager.Hubs;

public class ParkingLotHub : Hub
{
    private const string UpdateLogMessageName = "action_update";
    
    public async Task SendActionLogUpdate(ActionLogDto actionLog)
    {
        await Clients.All.SendAsync(UpdateLogMessageName, actionLog);
    }
}