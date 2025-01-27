package handlers

import (
	"context"
	"time"

	"liaison_go/business"

	service_v1 "liaison_go/generated/service"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type TrackingHandler struct {
	service_v1.UnimplementedTrackingServiceServer
	business.Tracker
}

func NewTrackingHandler(tracker business.Tracker) *TrackingHandler {
	return &TrackingHandler{
		Tracker: tracker,
	}
}

func (h *TrackingHandler) List(ctx context.Context, in *service_v1.ListRequest) (*service_v1.ListResponse, error) {
	status := ToDomainShipmentStatus(in.Status)
	from := in.From.AsTime()
	to := in.To.AsTime()

	result, err := h.Tracker.List(ctx, in.ShipmentIds, &status, &from, &to)
	if err != nil {
		return nil, err
	}

	// Find the earliest validUntil date
	validUntil := time.Now().UTC()
	if len(result) > 0 {
		validUntil = result[0].ValidUntil
		for _, shipment := range result[1:] {
			if shipment.ValidUntil.Before(validUntil) {
				validUntil = shipment.ValidUntil
			}
		}
	}

	return &service_v1.ListResponse{
		Shipments:  ToProtoShipments(result),
		ValidUntil: timestamppb.New(validUntil),
	}, nil
}

func (h *TrackingHandler) Place(ctx context.Context, in *service_v1.PlaceRequest) (*emptypb.Empty, error) {

	return nil, nil
}
