using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;

namespace ParkingManager.Controllers;

[Route("api/[controller]")]
[ApiController]
public class ParkingLotController : ControllerBase
{
    [HttpGet]
    public async Task<IActionResult> GetParkingLotStatus()
    {
        return Ok();
    }

    [HttpGet("action-logs")]
    public async Task<IActionResult> GetParkingLotActionLogs()
    {
        return Ok();
    }
}

