using OneOf;

namespace infrastructure.Domain;
public class Order
{
    public Guid Id { get; set; }
    public string CreatedBy { get; set; } = string.Empty;
    public OrderStatus Status { get; set; }
    public DateTime ReceivedAt { get; set; }
    public OneOf<TimeSpan, DateTime> ProcessTime { get; set; }
    public List<OrderState> StatesHistory { get; set; } = [];
}
