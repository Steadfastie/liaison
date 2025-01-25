using infrastructure.Domain;

namespace infrastructure;
internal static class Collections
{
    public static string GetCollectionName<T>() where T : Entity =>
        typeof(T) switch
        {
            Type entityType when entityType == typeof(Order) => "orders",
            _ => throw new ArgumentOutOfRangeException($"No collection for: {typeof(T).FullName}")
        };
}
