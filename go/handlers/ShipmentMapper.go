package handlers

import (
	d "liaison_go/domain"
	pb "liaison_go/generated/service"

	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// Map the domain Shipments to protobuf Shipments
func ToProtoShipments(shipments []d.Shipment) []*pb.Shipment {
	var pbShipments []*pb.Shipment
	for _, shipment := range shipments {
		protoShipment := toProtoShipment(&shipment)
		pbShipments = append(pbShipments, protoShipment)
	}
	return pbShipments
}

func toProtoShipment(shipment *d.Shipment) *pb.Shipment {
	return &pb.Shipment{
		ShipmentId:  shipment.ShipmentId,
		Status:      ToProtoShipmentStatus(&shipment.Status),
		LastUpdated: timestamppb.New(shipment.LastUpdated),
		Location:    shipment.Location,
	}
}
