package domain

// ShipmentStatus represents the status of a shipment.
type ShipmentStatus int

const (
	Pending ShipmentStatus = iota
	InProgress
	Completed
	Failed
)
