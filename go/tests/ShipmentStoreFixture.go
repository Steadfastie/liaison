package tests

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"liaison_go/persistence"

	"go.uber.org/zap"
)

type ShipmentStoreFixture struct {
	collection *mongo.Collection
	logger     *zap.Logger
}

func NewShipmentStoreFixutre(db *mongo.Database, logger *zap.Logger) *ShipmentStoreFixture {
	collection := db.Collection(persistence.CollectioName)
	return &ShipmentStoreFixture{
		collection: collection,
		logger:     logger.Named("ShipmentStore"),
	}
}

func (fixture *ShipmentStoreFixture) GetAll() ([]persistence.Shipment, error) {
	ctx := context.Background()
	cursor, err := fixture.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var docs []persistence.Shipment
	if err = cursor.All(ctx, &docs); err != nil {
		return nil, err
	}
	return docs, nil
}

func (fixture *ShipmentStoreFixture) ClearAll() error {
	_, err := fixture.collection.DeleteMany(context.Background(), bson.D{})
	return err
}
