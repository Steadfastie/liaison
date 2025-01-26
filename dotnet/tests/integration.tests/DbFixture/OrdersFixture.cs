using infrastructure;
using infrastructure.Domain;
using MongoDB.Driver;

namespace integration.tests.DbFixture;
public class OrdersFixture
{
    private readonly IMongoCollection<Order> _collection;
    public OrdersFixture(IMongoDatabase db)
    {
        Extensions.RegisterMappings();
        var collectionName = Collections.GetCollectionName<Order>();
        _collection = db.GetCollection<Order>(collectionName);
    }

    public async Task ClearAll()
    {
        var filter = Builders<Order>.Filter.Empty;
        await _collection.DeleteManyAsync(filter);
    }
}
