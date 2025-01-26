using infrastructure;
using infrastructure.Domain;
using MongoDB.Driver;

namespace integration.tests;
internal class OrdersFixture
{
    private readonly IMongoCollection<Order> _collection;
    internal OrdersFixture(IMongoDatabase db)
    {
        Extensions.RegisterMappings();
        var collectionName = Collections.GetCollectionName<Order>();
        _collection = db.GetCollection<Order>(collectionName);
    }

    internal async Task ClearAll()
    {
        var filter = Builders<Order>.Filter.Empty;
        await _collection.DeleteManyAsync(filter);
    }
}
