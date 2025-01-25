using infrastructure.Domain;
using infrastructure.Repos;
using Microsoft.Extensions.DependencyInjection;
using MongoDB.Bson.Serialization;
using MongoDB.Driver;

namespace infrastructure;
public static class Extensions
{
    public static void AddPersistence(this IServiceCollection services, MongoSettings settings)
    {
        services.AddSingleton<IMongoDatabase>(provider =>
        {
            var client = new MongoClient(settings.ConnectionString);
            return client.GetDatabase(settings.Database);
        });
        services.AddScoped<IOrderRepo, OrderRepo>();

        RegisterMappings();
    }

    private static void RegisterMappings()
    {
        BsonSerializer.RegisterSerializer(new OneOfTimeSpanDateTimeSerializer());

        BsonClassMap.RegisterClassMap<Order>();
    }
}
