using infrastructure.Domain;

namespace infrastructure;
public interface IOrderRepo
{
    public Task Create(Order order);
}
