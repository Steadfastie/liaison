using Google.Protobuf.WellKnownTypes;
using infrastructure.Domain;
using Liaison.V1;

namespace service.Mappers;

public static class CreateOrderMapper
{
    public static Response MapFromDomain(this Order order)
    {
        var response = new Response()
        {
            OrderId = order.Id.ToString(),
            Status = order.Status.MapFromDomain(),
            ReceivedAt = Timestamp.FromDateTime(order.ReceivedAt)
        };
        if (order.Duration is not null)
        {
            response.Duration = Duration.FromTimeSpan(order.Duration.Value);
        }
        else if (order.ProcessedAt is not null)
        {
            response.ProcessedAt = Timestamp.FromDateTime(order.ProcessedAt.Value);
        }
        response.StatesHistory.AddRange(order.StatesHistory.Select(i => new State()
        {
            Timestamp = Timestamp.FromDateTime(i.Timestamp),
            Status = i.Status.MapFromDomain()
        }));
        return response;
    }

    public static Status MapFromDomain(this OrderStatus status) =>
        status switch
        {
            OrderStatus.Received => Status.Received,
            OrderStatus.Processing => Status.Processing,
            OrderStatus.Valid => Status.Valid,
            OrderStatus.Invalid => Status.Invalid,
            _ => throw new ArgumentOutOfRangeException(nameof(status), $"Unknown status: {status}"),
        };
}
