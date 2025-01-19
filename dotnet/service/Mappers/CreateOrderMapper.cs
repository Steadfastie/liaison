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
        order.ProcessTime.Switch(
            (TimeSpan timeSpan) => response.Duration = Duration.FromTimeSpan(timeSpan),
            (DateTime dateTime) => response.ProcessedAt = Timestamp.FromDateTime(dateTime)
        );
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
