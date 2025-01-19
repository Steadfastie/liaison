namespace infrastructure.Domain;

public record OrderState
{
    public DateTime Timestamp { get; set; }
    public OrderStatus Status { get; set; }
}
