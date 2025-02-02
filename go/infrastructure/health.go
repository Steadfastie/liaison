package infrastructure

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"go.uber.org/zap"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func MonitorMongo(client *mongo.Client, healthServer *health.Server, logger *zap.Logger, system string) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	for {
		<-ticker.C
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		err := client.Ping(ctx, readpref.Primary())
		cancel()
		if err != nil {
			healthServer.SetServingStatus(system, healthpb.HealthCheckResponse_NOT_SERVING)
			log.Printf("mongo ping failed: %v", err)
		} else {
			healthServer.SetServingStatus(system, healthpb.HealthCheckResponse_SERVING)
			log.Printf("mongo ping succeeded")
		}
	}
}
