namespace ParkingManager.Configurations;

public class MinioSettings
{
    public const string SectionName = "Minio";
    
    public string InnerLink { get; set; } = string.Empty;
    public string OuterLink { get; set; } = string.Empty;
}