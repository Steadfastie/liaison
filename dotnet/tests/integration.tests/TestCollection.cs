using integration.tests.DbFixture;
using integration.tests.gRpcClient;

namespace integration.tests;

[CollectionDefinition("integration")]
public class TestCollection :
    ICollectionFixture<GRpcClient>,
    ICollectionFixture<DatabaseFixture>
{
}
