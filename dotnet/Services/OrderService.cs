using Google.Protobuf.WellKnownTypes;
using Grpc.Core;
using Liaison.Proto;
using Status = Liaison.Proto.Status;
using static Liaison.Proto.OrderService;

namespace dotnet.Services;

public class OrderService : OrderServiceBase
{
    public override Task<Response> CreateOrder(Request request, ServerCallContext context)
    {
        var resp = new Response
        {
            OrderId = Guid.NewGuid().ToString(),
            Status = Status.Valid,
            Message = "test",
            ReceivedAt = Timestamp.FromDateTime(DateTime.UtcNow),
            Duration = Duration.FromTimeSpan(TimeSpan.FromMilliseconds(30))
        };
        resp.StatesHistory.AddRange([
            new State
            {
                Timestamp = Timestamp.FromDateTime(DateTime.UtcNow),
                Status = Status.Received
            },
            new State
            {
                Timestamp = Timestamp.FromDateTime(DateTime.UtcNow.AddMilliseconds(10)),
                Status = Status.Received
            },
            new State
            {
                Timestamp = Timestamp.FromDateTime(DateTime.UtcNow.AddMilliseconds(20)),
                Status = Status.Received
            },
        ]);

        return Task.FromResult(resp);
    }
}
