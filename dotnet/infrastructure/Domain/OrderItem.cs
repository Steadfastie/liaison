namespace infrastructure.Domain;
public record OrderItem
{
    public required string Code { get; init; }
    public int Quantity { get; init; }
    public long Price { get; init; }
}
