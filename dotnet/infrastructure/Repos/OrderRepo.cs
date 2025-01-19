using infrastructure.Domain;

namespace infrastructure.Repos;

public class OrderRepo : IOrderRepo
{
    public Task<Order> Create(string orderId, string createdBy, List<OrderItem> items)
    {
        var now = DateTime.UtcNow;
        var order = new Order
        {
            Id = Guid.TryParse(orderId, out var id) ? id : Guid.NewGuid(),
            Status = OrderStatus.Valid,
            ReceivedAt = now,
            ProcessTime = TimeSpan.FromMilliseconds(30),
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
            CreatedBy = createdBy
        };
        var statusHistory = order.StatesHistory.OrderBy(i => i.Timestamp);
        order.ProcessTime = statusHistory.Last().Timestamp - statusHistory.First().Timestamp;

        return Task.FromResult(order);
    }
}
