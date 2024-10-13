namespace ParkingManager.Services.Interfaces;

public interface IRpcClient
{
    Task<string> RecognizeImageAsync(byte[] imageBytes);
}