using System.Text;
using Microsoft.Extensions.Options;
using ParkingManager.Configurations;
using ParkingManager.Services.Interfaces;
using RabbitMQ.Client;
using RabbitMQ.Client.Events;

namespace ParkingManager.RabbitMq;

public class MessageConsumer(
    IOptions<RabbitMqSettings> settings,
    IServiceScopeFactory serviceScopeFactory
    ): IAsyncDisposable, IHostedService
{
    private readonly IConnectionFactory _connectionFactory = new ConnectionFactory
    {
        HostName = settings.Value.HostName,
        UserName = settings.Value.UserName,
        Password = settings.Value.Password
    };
    
    IConnection? _connection;
    IChannel? _channel;

    private async Task StartListening()
    {
        _connection = await _connectionFactory.CreateConnectionAsync();
        _channel = await _connection.CreateChannelAsync();
        
        // await _channel.ExchangeDeclareAsync(exchange: "",
        //     type: ExchangeType.Direct);

        await ListenToEnteredCarQueue();
        await ListenToExitedCarQueue();
    }

    private async Task ListenToExitedCarQueue()
    {
        var queueName = settings.Value.ExitedCarQueueName;
        QueueDeclareOk queueDeclareResult = await _channel!.QueueDeclareAsync(queue: queueName);
        //await _channel!.QueueBindAsync(queue: queueName, exchange: "", routingKey: string.Empty);
   
        Console.WriteLine(" [*] Waiting for logs.");
   
        var consumer = new AsyncEventingBasicConsumer(_channel);
        consumer.ReceivedAsync += async (model, ea) =>
        {
            byte[] body = ea.Body.ToArray();
            var message = Encoding.UTF8.GetString(body);
            Console.WriteLine($" [x] Received from exited car queue: {message}");
            using var scope = serviceScopeFactory.CreateScope();
            var parkingManager =
                scope.ServiceProvider.GetRequiredService<IParkingManager>();

            await parkingManager.UnparkCar(message);
        };
   
        await _channel.BasicConsumeAsync(queueName, autoAck: true, consumer: consumer);
    }

    private async Task ListenToEnteredCarQueue()
    {
        var queueName = settings.Value.EnteredCarQueueName;
        QueueDeclareOk queueDeclareResult = await _channel!.QueueDeclareAsync(queue: queueName);
        //await _channel!.QueueBindAsync(queue: queueName, exchange: "", routingKey: string.Empty);
   
        Console.WriteLine(" [*] Waiting for logs.");
   
        var consumer = new AsyncEventingBasicConsumer(_channel);
        consumer.ReceivedAsync += async (model, ea) =>
        {
            byte[] body = ea.Body.ToArray();
            var message = Encoding.UTF8.GetString(body);
            Console.WriteLine($" [x] Received from entered car queue: {message}");
            using var scope = serviceScopeFactory.CreateScope();
            var parkingManager =
                scope.ServiceProvider.GetRequiredService<IParkingManager>();

            await parkingManager.ParkCar(message);
        };
   
        await _channel.BasicConsumeAsync(queueName, autoAck: true, consumer: consumer);
    }
    
    
    public async ValueTask DisposeAsync()
    {
        if (_connection is not null)
        {
            await _connection.CloseAsync();
        }
        
        if (_channel is not null)
        {
            await _channel.CloseAsync();
        }
    }

    public async Task StartAsync(CancellationToken cancellationToken)
    {
        await StartListening();
    }

    public Task StopAsync(CancellationToken cancellationToken)
    {
        return Task.CompletedTask;
    }
}