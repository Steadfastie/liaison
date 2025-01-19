using infrastructure.Domain;

namespace infrastructure;
public interface IOrderRepo
{
    public Task<Order> Create(string orderId, string createdBy, List<OrderItem> items);
}
