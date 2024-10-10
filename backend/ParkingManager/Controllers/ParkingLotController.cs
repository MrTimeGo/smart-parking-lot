using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using ParkingManager.Dtos;
using ParkingManager.Services.Interfaces;

namespace ParkingManager.Controllers;

[Route("api/[controller]")]
[ApiController]
public class ParkingLotController(IParkingManager parkingManager) : ControllerBase
{
    [HttpGet]
    public async Task<ActionResult<List<ParkingDto>>> GetParkingLotStatus()
    {
        return Ok(await parkingManager.GetParkingStatus());
    }

    [HttpGet("action-logs")]
    public async Task<ActionResult<List<ActionLogDto>>> GetParkingLotActionLogs()
    {
        return Ok(await parkingManager.GetLastActionLogs());
    }
}

