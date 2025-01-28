package domain

import (
	"time"
)

// Shipment represents shipment details.
type Shipment struct {
	ShipmentId  string
	Status      ShipmentStatus
	LastUpdated time.Time
	Location    string
	ValidUntil  time.Time
}
