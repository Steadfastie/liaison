using infrastructure.Domain;
using infrastructure.Repos;
using Microsoft.Extensions.DependencyInjection;
using MongoDB.Bson;
using MongoDB.Bson.Serialization;
using MongoDB.Bson.Serialization.Serializers;
using MongoDB.Driver;

namespace infrastructure;
public static class Extensions
{
    public static void AddPersistence(this IServiceCollection services, MongoSettings settings)
    {
        services.AddSingleton(provider =>
        {
            var client = new MongoClient(settings.ConnectionString);
            return client.GetDatabase(settings.Database);
        });
        services.AddScoped<IOrderRepo, OrderRepo>();

        RegisterMappings();
    }

    public static void RegisterMappings()
    {
        BsonSerializer.RegisterSerializer(new EnumSerializer<OrderStatus>(BsonType.String));

        BsonClassMap.RegisterClassMap<Order>();
        BsonClassMap.RegisterClassMap<OrderState>();
    }
}
