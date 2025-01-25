using infrastructure.Domain;
using MongoDB.Driver;

namespace infrastructure.Repos;
internal abstract class Repo<T> where T : Entity
{
    protected IMongoCollection<T> Collection { get; set; }
    internal Repo(IMongoDatabase db)
    {
        var collectionName = Collections.GetCollectionName<T>();
        Collection = db.GetCollection<T>(collectionName);
    }
}
