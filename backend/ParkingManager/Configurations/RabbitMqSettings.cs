namespace ParkingManager.Configurations;

public class RabbitMqSettings
{
    public const string SectionName = "RabbitMq";
    
    public string HostName { get; set; } = string.Empty;
    public string UserName { get; set; } = string.Empty;
    public string Password { get; set; } = string.Empty;

    public string RpcQueueName { get; set; } = string.Empty;
    
    public string EnteredCarQueueName { get; set; } = string.Empty;
    public string ExitedCarQueueName { get; set; } = string.Empty;
}