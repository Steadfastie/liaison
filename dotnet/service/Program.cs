using application.Handlers;
using infrastructure;
using infrastructure.Repos;
using Microsoft.Extensions.Diagnostics.HealthChecks;
using service.Services;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
builder.Services.AddGrpc();

builder.Services.AddSingleton<IOrderRepo, OrderRepo>();

builder.Services.AddMediatR(cfg =>
{
    cfg.RegisterServicesFromAssembly(typeof(CreateOrderHandler).Assembly);
});

builder.Services.AddGrpcHealthChecks()
       .AddCheck(string.Empty, () => HealthCheckResult.Healthy());

var app = builder.Build();

// Configure the HTTP request pipeline.
app.MapGrpcService<OrderService>();
app.MapGrpcHealthChecksService();
app.MapGet("/", () => "Order service");

app.Run();