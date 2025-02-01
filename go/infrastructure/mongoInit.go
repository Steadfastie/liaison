package infrastructure

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.uber.org/zap"
)

func InitialiseMongoDB(logger *zap.Logger, conf *Config) (*mongo.Database, func()) {
	logger.Info("Setting up Mongo")
	bsonOpts := &options.BSONOptions{
		UseJSONStructTags: true,
	}
	clientOpts := options.Client().
		ApplyURI(conf.MongoSettings.ConnectionString).
		SetBSONOptions(bsonOpts)
	client, err := mongo.Connect(clientOpts)
	if err != nil {
		panic(fmt.Errorf("mongo client connection error %w", err))
	}
	disconnect := func() {
		if err := client.Disconnect(context.Background()); err != nil {
			logger.Error("Mongo client disconnection error", zap.Error(err))
		}
	}
	db := client.Database(conf.MongoSettings.Database)
	return db, disconnect
}
