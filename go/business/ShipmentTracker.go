package business

import (
	"context"
	"liaison_go/domain"
	"liaison_go/persistence"
	"time"
)

type ShipmentTracker struct {
	persistence.Store
}

func NewShipmentTracker(store persistence.Store) *ShipmentTracker {
	return &ShipmentTracker{
		Store: store,
	}
}

func (tracker *ShipmentTracker) List(
	ctx context.Context,
	ids []string,
	status *domain.ShipmentStatus,
	from *time.Time,
	to *time.Time,
) ([]domain.Shipment, error) {
	result, err := tracker.Store.GetMany(ctx, ids, status, from, to)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (tracker *ShipmentTracker) Place(ctx context.Context, shipments []domain.Shipment, validUntil time.Time) error {
	// Set identical validUntil for all shipments
	for i := range shipments {
		shipments[i].ValidUntil = validUntil
	}
	err := tracker.Store.Create(ctx, shipments)
	return err
}
