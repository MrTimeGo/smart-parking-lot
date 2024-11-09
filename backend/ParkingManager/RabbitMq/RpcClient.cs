using System.Collections.Concurrent;
using System.Text;
using Microsoft.Extensions.Options;
using ParkingManager.Configurations;
using RabbitMQ.Client;
using RabbitMQ.Client.Events;

namespace ParkingManager.RabbitMq;

public class RpcClient(IOptions<RabbitMqSettings> settings) : IAsyncDisposable
{
    private readonly string _queueName = settings.Value.RpcQueueName;

    private readonly IConnectionFactory _connectionFactory = new ConnectionFactory
    {
        HostName = settings.Value.HostName,
        UserName = settings.Value.UserName,
        Password = settings.Value.Password
    };

    private readonly ConcurrentDictionary<string, TaskCompletionSource<string>> _callbackMapper
        = new();

    private IConnection? _connection;
    private IChannel? _channel;
    private string? _replyQueueName;

    public async Task StartAsync()
    {
        _connection = await _connectionFactory.CreateConnectionAsync();
        _channel = await _connection.CreateChannelAsync();

        // declare a server-named queue
        QueueDeclareOk queueDeclareResult = await _channel.QueueDeclareAsync();
        _replyQueueName = queueDeclareResult.QueueName;
        var consumer = new AsyncEventingBasicConsumer(_channel);

        consumer.ReceivedAsync += (model, ea) =>
        {
            string? correlationId = ea.BasicProperties.CorrelationId;

            if (false == string.IsNullOrEmpty(correlationId))
            {
                if (_callbackMapper.TryRemove(correlationId, out var tcs))
                {
                    var body = ea.Body.ToArray();
                    var response = Encoding.UTF8.GetString(body);
                    tcs.TrySetResult(response);
                }
            }

            return Task.CompletedTask;
        };

        await _channel.BasicConsumeAsync(_replyQueueName, true, consumer);
    }

    public async Task<string> CallAsync(string message,
        CancellationToken cancellationToken = default)
    {
        var messageBytes = Encoding.UTF8.GetBytes(message);
        return await CallAsync(messageBytes, cancellationToken);
    }

    public async Task<string> CallAsync(byte[] messageBytes,
        CancellationToken cancellationToken = default)
    {
        if (_channel is null)
        {
            throw new InvalidOperationException();
        }

        string correlationId = Guid.NewGuid().ToString();
        var props = new BasicProperties
        {
            CorrelationId = correlationId,
            ReplyTo = _replyQueueName
        };

        var tcs = new TaskCompletionSource<string>(
            TaskCreationOptions.RunContinuationsAsynchronously);
        _callbackMapper.TryAdd(correlationId, tcs);

        await _channel.BasicPublishAsync(exchange: string.Empty, routingKey: _queueName,
            mandatory: true, basicProperties: props, body: messageBytes);

        using CancellationTokenRegistration ctr =
            cancellationToken.Register(() =>
            {
                _callbackMapper.TryRemove(correlationId, out _);
                tcs.SetCanceled();
            });

        return await tcs.Task;
    }

    public async ValueTask DisposeAsync()
    {
        if (_channel is not null)
        {
            await _channel.CloseAsync();
        }

        if (_connection is not null)
        {
            await _connection.CloseAsync();
        }
    }
}