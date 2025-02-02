using application.Handlers;
using infrastructure;
using Microsoft.Extensions.Diagnostics.HealthChecks;
using Serilog;
using Serilog.Formatting.Compact;
using service;
using service.Services;

Log.Logger = new LoggerConfiguration()
    .Enrich.FromLogContext()
    .WriteTo.Console(new RenderedCompactJsonFormatter())
    .CreateLogger();

try
{
    Log.Information("Starting up");

    var builder = WebApplication.CreateBuilder(args);

    builder.Services.AddSerilog();
    builder.Services.AddGrpc();

    var mongoSettings = builder.Configuration.GetSection(nameof(MongoSettings)).Get<MongoSettings>()!;
    builder.Services.AddPersistence(mongoSettings);

    builder.Services.AddMediatR(cfg =>
    {
        cfg.RegisterServicesFromAssembly(typeof(CreateOrderHandler).Assembly);
    });

    builder.Services.AddSingleton<IHealthCheck, MongoHealthCheck>();
    builder.Services.AddGrpcHealthChecks()
         .AddCheck<MongoHealthCheck>(string.Empty);

    var app = builder.Build();

    // Configure the HTTP request pipeline.
    app.MapGrpcService<OrderService>();
    app.MapGrpcHealthChecksService();
    app.MapGet("/", () => "Order service");

    app.Run();

    Log.Information("Stopped cleanly");
}
catch (Exception ex)
{
    Log.Fatal(ex, "Application terminated unexpectedly");
}
finally
{
    Log.CloseAndFlush();
}