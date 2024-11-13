using Microsoft.EntityFrameworkCore;
using ParkingManager.Configurations;
using ParkingManager.Data;
using ParkingManager.Hubs;
using ParkingManager.RabbitMq;
using ParkingManager.Services.Interfaces;

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddControllers();
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

builder.Services.Configure<ParkingSettings>(
    builder.Configuration.GetSection(ParkingSettings.SectionName)
);

builder.Services.Configure<RabbitMqSettings>(
    builder.Configuration.GetSection(RabbitMqSettings.SectionName)
);

builder.Services.Configure<MinioSettings>(
    builder.Configuration.GetSection(MinioSettings.SectionName)
);

builder.Services.AddDbContext<ParkingContext>(options =>
    options.UseNpgsql(builder.Configuration.GetConnectionString("Postgresql"))
);

builder.Services.AddScoped<IParkingManager, ParkingManager.Services.ParkingManager>();
builder.Services.AddSingleton<RpcClient>();

builder.Services.AddHostedService<MessageConsumer>();

builder.Services.AddSignalR();

var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

using (var scope = app.Services.CreateScope())
{
    await scope.ServiceProvider.GetRequiredService<ParkingContext>().Database.MigrateAsync();
    await scope.ServiceProvider.GetRequiredService<RpcClient>().StartAsync();
}

app.UseHttpsRedirection();

string[] origins = [builder.Configuration["Cors:Frontend"]!];

app.UseCors(cors => 
    cors.WithOrigins(origins)
        .AllowAnyHeader()
        .AllowAnyMethod()
        .AllowCredentials()
);

app.UseAuthorization();

app.MapControllers();
app.MapHub<ParkingLotHub>("/parkingLotHub");

app.Run();