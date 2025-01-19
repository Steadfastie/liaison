using infrastructure;
using infrastructure.Domain;
using MediatR;
using OneOf;
using OneOf.Types;

namespace application.Handlers;

public class CreateOrderRequest : IRequest<OneOf<Order, No>>
{
    public required string OrderId { get; init; }
    public string CreatedBy { get; init; } = string.Empty;
    public List<OrderItem> Items { get; init; } = [];
};
public class CreateOrderHandler : IRequestHandler<CreateOrderRequest, OneOf<Order, No>>
{
    private readonly IOrderRepo _orderRepo;
    public CreateOrderHandler(IOrderRepo orderRepo)
    {
        _orderRepo = orderRepo;
    }

    public async Task<OneOf<Order, No>> Handle(CreateOrderRequest request, CancellationToken cancellationToken)
    {
        return await _orderRepo.Create(request.OrderId, request.CreatedBy, request.Items);
    }
}
