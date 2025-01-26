using DnsClient.Internal;
using infrastructure;
using infrastructure.Domain;
using MediatR;
using Microsoft.Extensions.Logging;
using OneOf;
using OneOf.Types;

namespace application.Handlers;

public class CreateOrderRequest : IRequest<OneOf<Order, No>>
{
    public string CreatedBy { get; init; } = string.Empty;
    public List<OrderItem> Items { get; init; } = [];
};
public class CreateOrderHandler : IRequestHandler<CreateOrderRequest, OneOf<Order, No>>
{
    private readonly IOrderRepo _orderRepo;
    private readonly ILogger<CreateOrderHandler> _logger;
    public CreateOrderHandler(IOrderRepo orderRepo, ILogger<CreateOrderHandler> logger)
    {
        _orderRepo = orderRepo;
        _logger = logger;
    }

    public async Task<OneOf<Order, No>> Handle(CreateOrderRequest request, CancellationToken cancellationToken)
    {
        var now = DateTime.UtcNow;
        var order = new Order
        {
            Status = OrderStatus.Valid,
            ReceivedAt = now,
            StatesHistory =
            [
                new() {
                    Timestamp = now,
                    Status = OrderStatus.Received
                },
                new() {
                    Timestamp = now.AddMilliseconds(10),
                    Status = OrderStatus.Processing
                },
                new() {
                    Timestamp = now.AddMilliseconds(20),
                    Status = OrderStatus.Valid
                },
            ],
            CreatedBy = request.CreatedBy
        };
        var statusHistory = order.StatesHistory.OrderBy(i => i.Timestamp);
        order.Duration = statusHistory.Last().Timestamp - statusHistory.First().Timestamp;

        _logger.LogInformation("Creating order with {Id} id", order.Id);
        await _orderRepo.Create(order);

        return order;
    }
}
