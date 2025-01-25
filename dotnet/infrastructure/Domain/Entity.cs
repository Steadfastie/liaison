using MongoDB.Bson.Serialization.Attributes;

namespace infrastructure.Domain;
public abstract class Entity
{
    [BsonId]
    public string Id { get; set; } = Guid.NewGuid().ToString("D").Split('-').Last();
    public required string CreatedBy { get; set; }
}
