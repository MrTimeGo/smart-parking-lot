using Microsoft.EntityFrameworkCore;
using ParkingManager.Entities;

namespace ParkingManager.Data;

public class ParkingContext(DbContextOptions<ParkingContext> options) : DbContext(options)
{
    public DbSet<ActionLog> ActionLogs { get; set; }
}