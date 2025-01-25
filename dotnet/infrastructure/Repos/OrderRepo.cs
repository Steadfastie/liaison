using infrastructure.Domain;
using MongoDB.Driver;

namespace infrastructure.Repos;

internal class OrderRepo : Repo<Order>, IOrderRepo
{
    public OrderRepo(IMongoDatabase db) : base(db)
    {
    }

    public async Task Create(Order order)
    {
        await Collection.InsertOneAsync(order);
    }
}
