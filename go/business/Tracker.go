package business

import (
	"context"
	"time"

	"liaison_go/domain"
)

type Tracker interface {
	List(ctx context.Context, ids []string, status *domain.ShipmentStatus, from *time.Time, to *time.Time) ([]domain.Shipment, error)
	Place(ctx context.Context, shipments []domain.Shipment, validUntil time.Time) error
}
