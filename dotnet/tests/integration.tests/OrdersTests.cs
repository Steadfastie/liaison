namespace integration.tests;

[Collection("integration")]
public class OrdersTests
{
    private DbFixture _dbFixutre;
    public OrdersTests(DbFixture dbFixutre)
    {
        _dbFixutre = dbFixutre;
    }

    [Fact]
    public async Task Create()
    {
        await _dbFixutre.Orders.ClearAll();

    }
}
