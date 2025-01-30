package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"liaison_go/business"
	"liaison_go/handlers"
	"liaison_go/persistence"

	service_v1 "liaison_go/generated/service"

	"github.com/spf13/viper"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"google.golang.org/grpc"
)

const ()

func main() {
	// Config intialization
	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("config file error %w", err))
	}

	var conf Config
	err = viper.Unmarshal(&conf)
	if err != nil {
		panic(fmt.Errorf("config unmarshalling error %w", err))
	}

	// MongoDB intialization
	bsonOpts := &options.BSONOptions{
		UseJSONStructTags: true,
	}
	clientOpts := options.Client().
		ApplyURI(conf.MongoSettings.ConnectionString).
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
	db := client.Database(conf.MongoSettings.Database)

	// Services intialization
	shipmentStore := persistence.NewShipmentStore(db)               // 3rd layer
	shipmentTracker := business.NewShipmentTracker(shipmentStore)   // 2nd layer
	trackingHandler := handlers.NewTrackingHandler(shipmentTracker) // 1rd layer

	// gRPC intialization
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", conf.Host.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	service_v1.RegisterTrackingServiceServer(grpcServer, trackingHandler)
	grpcServer.Serve(lis)
}
