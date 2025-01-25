using MongoDB.Bson;
using MongoDB.Bson.Serialization.Attributes;
using OneOf;

namespace infrastructure.Domain;
public class Order : Entity
{
    [BsonRepresentation(BsonType.String)]
    public OrderStatus Status { get; set; }
    public DateTime ReceivedAt { get; set; }
    public OneOf<TimeSpan, DateTime> ProcessTime { get; set; }
    public List<OrderState> StatesHistory { get; set; } = [];
}
