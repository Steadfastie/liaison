package main

import (
	"context"
	"log"
	"net"

	"liaison_go/business"
	"liaison_go/handlers"
	"liaison_go/persistence"

	service_v1 "liaison_go/generated/service"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"google.golang.org/grpc"
)

var (
	serverAddr = "localhost:5001"
	mongoAddr  = "mongodb://localhost:27017"
	database   = "liaison"
)

const ()

func main() {

	// MongoDB intialization
	bsonOpts := &options.BSONOptions{
		UseJSONStructTags: true,
	}
	clientOpts := options.Client().
		ApplyURI(mongoAddr).
		SetBSONOptions(bsonOpts)
	client, err := mongo.Connect(clientOpts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	db := client.Database(database)

	// Services intialization
	shipmentStore := persistence.NewShipmentStore(db)               // 3rd layer
	shipmentTracker := business.NewShipmentTracker(shipmentStore)   // 2nd layer
	trackingHandler := handlers.NewTrackingHandler(shipmentTracker) // 1rd layer

	// gRPC intialization
	lis, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	service_v1.RegisterTrackingServiceServer(grpcServer, trackingHandler)
	grpcServer.Serve(lis)
}
