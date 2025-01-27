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

func (t *ShipmentTracker) List(ctx context.Context, ids []string, status *domain.ShipmentStatus, from *time.Time, to *time.Time) ([]domain.Shipment, error) {
	return nil, nil
}

func (t *ShipmentTracker) Place(ctx context.Context) error {
	return nil
}
