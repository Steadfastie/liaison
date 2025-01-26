namespace infrastructure.Domain;
public class Order : Entity
{
    public OrderStatus Status { get; set; }
    public DateTime ReceivedAt { get; set; }
    public DateTime? ProcessedAt { get; set; }
    public TimeSpan? Duration { get; set; }
    public List<OrderState> StatesHistory { get; set; } = [];
}
