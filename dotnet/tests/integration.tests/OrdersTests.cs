using integration.tests.DbFixture;
using integration.tests.gRpcClient;
using Liaison.V1;
using Shouldly;
using static Liaison.V1.Response;

namespace integration.tests;

[Collection("integration")]
public class OrdersTests
{
    private readonly DatabaseFixture _dbFixutre;
    private readonly GRpcClient _grpc;
    public OrdersTests(DatabaseFixture dbFixutre, GRpcClient grpc)
    {
        _dbFixutre = dbFixutre;
        _grpc = grpc;
    }

    [Fact]
    public async Task Create()
    {
        // arrange
        await _dbFixutre.Orders.ClearAll();

        var now = DateTime.UtcNow;
        var request = new Request
        {
            Items = {
                { "one", new OrderItem
                    {
                        Code = "item1",
                        Quantity = 5,
                        Price = 2.2
                    }
                },
                { "two", new OrderItem
                    {
                        Code = "item2",
                        Quantity = 5,
                        Price = 2.4
                    }
                }
            },
            CreatedBy = "test"
        };

        // act
        var response = await _grpc.Client.CreateOrderAsync(request);

        // assert
        response.Status.ShouldBe(Status.Valid);
        response.StatesHistory.ShouldNotBeNull();
        response.ReceivedAt.ToDateTime().ShouldBe(now, TimeSpan.FromSeconds(10));
        if (response.TestOneofCase is TestOneofOneofCase.ProcessedAt)
        {
            response.ProcessedAt.ToDateTime().ShouldBe(now, TimeSpan.FromSeconds(10));
        }
        if (response.TestOneofCase is TestOneofOneofCase.Duration)
        {
            var timespan = TimeSpan.FromSeconds(response.Duration.Seconds)
                   + TimeSpan.FromTicks(response.Duration.Nanos / 100);
            timespan.ShouldBeLessThan(TimeSpan.FromMilliseconds(100));
        }
    }
}
