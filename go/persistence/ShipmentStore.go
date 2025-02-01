package persistence

import (
	"context"
	"fmt"
	"liaison_go/domain"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.uber.org/zap"
)

const (
	CollectioName = "shipments"
)

type ShipmentStore struct {
	collection *mongo.Collection
	logger     *zap.Logger
}

func NewShipmentStore(db *mongo.Database, logger *zap.Logger) *ShipmentStore {
	collection := db.Collection(CollectioName)
	return &ShipmentStore{
		collection: collection,
		logger:     logger.Named("ShipmentStore"),
	}
}

type Shipment struct {
	ShipmentId  string    `bson:"_id"`
	Status      string    `bson:"status"`
	LastUpdated time.Time `bson:"lastUpdated"`
	Location    string    `bson:"location"`
	ValidUntil  time.Time `bson:"validUntil"`
}

func (store *ShipmentStore) GetMany(
	ctx context.Context,
	ids *[]string,
	status *domain.ShipmentStatus,
	from *time.Time,
	to *time.Time,
) ([]domain.Shipment, error) {
	filter := bson.M{}

	if len(*ids) > 0 {
		filter["_id"] = bson.M{"$in": ids}
	}

	if status != nil {
		filter["status"] = StatusToString(*status)
	}

	dateFilter := bson.M{}
	if from != nil {
		dateFilter["$gte"] = *from
	}
	if to != nil {
		dateFilter["$lte"] = *to
	}
	if len(dateFilter) > 0 {
		filter["lastUpdated"] = dateFilter
	}

	opts := options.Find().
		SetSort(bson.D{{Key: "lastUpdated", Value: 1}, {Key: "id", Value: 1}})

	cursor, err := store.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var docs []Shipment
	if err = cursor.All(ctx, &docs); err != nil {
		return nil, err
	}

	var results []domain.Shipment
	for _, doc := range docs {
		shipment := domain.Shipment{
			ShipmentId:  doc.ShipmentId,
			Status:      StringToStatus(doc.Status),
			LastUpdated: doc.LastUpdated,
			Location:    doc.Location,
			ValidUntil:  doc.ValidUntil,
		}
		results = append(results, shipment)
	}

	return results, nil
}

func (store *ShipmentStore) Create(ctx context.Context, shipments *[]domain.Shipment) error {
	// Convert domain to BSON
	var docs []Shipment
	for _, s := range *shipments {
		doc := Shipment{
			ShipmentId:  s.ShipmentId,
			Status:      StatusToString(s.Status),
			LastUpdated: s.LastUpdated,
			Location:    s.Location,
			ValidUntil:  s.ValidUntil,
		}
		docs = append(docs, doc)
	}
	_, err := store.collection.InsertMany(ctx, docs)
	return err
}

func StatusToString(s domain.ShipmentStatus) string {
	switch s {
	case domain.Pending:
		return "Pending"
	case domain.InProgress:
		return "InProgress"
	case domain.Completed:
		return "Completed"
	case domain.Failed:
		return "Failed"
	default:
		panic(fmt.Sprintf("invalid shipment status: %d", s))
	}
}

func StringToStatus(status string) domain.ShipmentStatus {
	switch status {
	case "Pending":
		return domain.Pending
	case "InProgress":
		return domain.InProgress
	case "Completed":
		return domain.Completed
	case "Failed":
		return domain.Failed
	default:
		panic(fmt.Sprintf("invalid shipment status: %v", status))
	}
}
