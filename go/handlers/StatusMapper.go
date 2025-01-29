package handlers

import (
	"fmt"
	d "liaison_go/domain"
	pb "liaison_go/generated/service"
)

// Map Original to Protobuf
func ToProtoShipmentStatus(status *d.ShipmentStatus) pb.ShipmentStatus {
	switch *status {
	case d.Pending:
		return pb.ShipmentStatus_SHIPMENT_STATUS_REGISTERED
	case d.InProgress:
		return pb.ShipmentStatus_SHIPMENT_STATUS_ONROUTE
	case d.Completed:
		return pb.ShipmentStatus_SHIPMENT_STATUS_SHIPPED
	case d.Failed:
		return pb.ShipmentStatus_SHIPMENT_STATUS_NOT_SHIPPED
	default:
		return pb.ShipmentStatus_SHIPMENT_STATUS_UNSPECIFIED
	}
}

// Map Protobuf to Original
func ToDomainShipmentStatus(protoStatus *pb.ShipmentStatus) d.ShipmentStatus {
	switch *protoStatus {
	case pb.ShipmentStatus_SHIPMENT_STATUS_REGISTERED:
		return d.Pending
	case pb.ShipmentStatus_SHIPMENT_STATUS_ONROUTE:
		return d.InProgress
	case pb.ShipmentStatus_SHIPMENT_STATUS_SHIPPED:
		return d.Completed
	case pb.ShipmentStatus_SHIPMENT_STATUS_NOT_SHIPPED:
		return d.Failed
	default:
		panic(fmt.Sprintf("invalid protoStatus: %v", protoStatus))
	}
}
