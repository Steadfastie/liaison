package persistence

import (
	"context"
	"liaison_go/domain"
	"time"
)

type Store interface {
	GetMany(ctx context.Context, ids *[]string, status *domain.ShipmentStatus, from *time.Time, to *time.Time) ([]domain.Shipment, error)
	Create(ctx context.Context, shipments *[]domain.Shipment) error
}
