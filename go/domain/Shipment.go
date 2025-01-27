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

type Shipments []Shipment

func (shipments Shipments) EarliestValidUntil() *time.Time {
	if len(shipments) == 0 {
		now := time.Now().UTC()
		return &now
	}

	earliest := &shipments[0].ValidUntil
	for _, shipment := range shipments {
		if shipment.ValidUntil.Before(*earliest) {
			earliest = &shipment.ValidUntil
		}
	}
	return earliest
}
