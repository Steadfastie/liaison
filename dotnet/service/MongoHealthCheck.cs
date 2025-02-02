using Microsoft.Extensions.Diagnostics.HealthChecks;
using MongoDB.Bson;
using MongoDB.Driver;

namespace service;
public class MongoHealthCheck : IHealthCheck
{
    private readonly IMongoDatabase _database;
    public MongoHealthCheck(IMongoDatabase database)
    {
        _database = database;
    }

    public async Task<HealthCheckResult> CheckHealthAsync(HealthCheckContext context, CancellationToken cancellationToken = default)
    {
        try
        {
            var command = new BsonDocument("ping", 1);
            await _database.RunCommandAsync<BsonDocument>(command, cancellationToken: cancellationToken);
            return HealthCheckResult.Healthy();
        }
        catch
        {
            return HealthCheckResult.Unhealthy();
        }

    }
}
