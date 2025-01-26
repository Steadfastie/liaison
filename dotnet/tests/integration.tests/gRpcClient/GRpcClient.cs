using Grpc.Core;
using Grpc.Net.Client;
using Grpc.Net.Client.Configuration;
using Microsoft.Extensions.Configuration;

namespace integration.tests.gRpcClient;

public class GRpcClient : IDisposable
{
    public readonly GrpcChannel Channel;
    public readonly Liaison.V1.OrderService.OrderServiceClient Client;
    public GRpcClient()
    {
        var config = new ConfigurationBuilder()
            .SetBasePath(Directory.GetCurrentDirectory())
            .AddJsonFile("appsettings.json", false, true)
            .Build();

        var settings = config
             .GetRequiredSection(nameof(ClientSettings))
             .Get<ClientSettings>()!;

        var defaultMethodConfig = new MethodConfig
        {
            Names = { MethodName.Default },
            RetryPolicy = new RetryPolicy
            {
                MaxAttempts = 3,
                InitialBackoff = TimeSpan.FromSeconds(1),
                MaxBackoff = TimeSpan.FromSeconds(7),
                BackoffMultiplier = 2,
                RetryableStatusCodes = { StatusCode.Unavailable }
            }
        };
        var channelOptions = new GrpcChannelOptions
        {
            ServiceConfig = new ServiceConfig
            {
                MethodConfigs = { defaultMethodConfig }
            }
        };
        Channel = GrpcChannel.ForAddress(settings.Url, channelOptions);
        Client = new Liaison.V1.OrderService.OrderServiceClient(Channel);
    }

    public void Dispose()
    {
        Channel.Dispose();
    }
}
