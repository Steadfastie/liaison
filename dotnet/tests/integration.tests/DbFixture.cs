using infrastructure;
using Microsoft.Extensions.Configuration;
using MongoDB.Driver;

namespace integration.tests;

[CollectionDefinition("integration")]
public class DbFixture : ICollectionFixture<DbFixture>
{
    internal readonly OrdersFixture Orders;
    public DbFixture()
    {
        var config = new ConfigurationBuilder()
            .SetBasePath(Directory.GetCurrentDirectory())
            .AddJsonFile("appsettings.json", false, true)
            .Build();

        var settings = config
             .GetRequiredSection(nameof(MongoSettings))
             .Get<MongoSettings>()!;

        var client = new MongoClient(settings.ConnectionString);
        var db = client.GetDatabase(settings.Database);

        Orders = new OrdersFixture(db);
    }
}
