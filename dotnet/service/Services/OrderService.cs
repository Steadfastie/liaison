using application.Handlers;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;
using Liaison.V1;
using MediatR;
using service.Mappers;

namespace service.Services;

public class OrderService : Liaison.V1.OrderService.OrderServiceBase
{
    private readonly IMediator _mediator;
    public OrderService(IMediator mediator)
    {
        _mediator = mediator;
    }

    public async override Task<Response> CreateOrder(Request request, ServerCallContext context)
    {
        var result = await _mediator.Send(new CreateOrderRequest()
        {
            OrderId = request.OrderId,
            CreatedBy = request.CreatedBy,
            Items = request.Items.Select(i => new infrastructure.Domain.OrderItem()
            {
                Code = i.Key,
                Quantity = i.Value.Quantity,
                Price = (long)i.Value.Price
            }).ToList()
        });

        return result.Match(
            order =>
            {
                var response = order.MapFromDomain();
                response.Message = "test";
                response.Details = Any.Pack(new State
                {
                    Status = Liaison.V1.Status.Received,
                    Timestamp = Timestamp.FromDateTime(order.StatesHistory.First().Timestamp)
                });
                return response;
            },
            no =>
            {
                throw new RpcException(new Grpc.Core.Status(StatusCode.Internal, "Something went wrong"));
            }
        );
    }
}
