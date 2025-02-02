package tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	service_v1 "liaison_go/generated/service"
	infra "liaison_go/infrastructure"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/stretchr/testify/require"

	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// https://pkg.go.dev/google.golang.org/grpc/examples/features/retry#section-readme
const gRpcClientConfig = `{
  "methodConfig": [{
    "name": [{}],
    "retryPolicy": {
      "MaxAttempts": 4,
      "InitialBackoff": "0.1s",
      "MaxBackoff": "1s",
      "BackoffMultiplier": 2,
      "RetryableStatusCodes": ["UNAVAILABLE"]
    }
  }]
}`

func TestShipmentService(t *testing.T) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	var conf infra.Config
	conf.Populate(logger)

	db, disconnect := infra.InitialiseMongoDB(logger, &conf)
	defer disconnect()
	dbFixture := NewShipmentStoreFixutre(db, logger)

	var opts []grpc.DialOption

	opts = append(opts, grpc.WithDefaultServiceConfig(gRpcClientConfig))
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient(fmt.Sprintf("localhost:%d", conf.Host.Port), opts...)
	if err != nil {
		panic(fmt.Errorf("could not connect to the gRPC server %w", err))
	}
	defer conn.Close()
	client := service_v1.NewTrackingServiceClient(conn)

	t.Run("Create_shipments", func(t *testing.T) {
		// Arrange
		dbFixture.ClearAll()
		now := time.Now().UTC()

		req := &service_v1.PlaceRequest{
			ValidUntil: timestamppb.New(now.Add(24 * time.Hour)),
			Shipments: []*service_v1.Shipment{
				{
					ShipmentId:  uuid.New().String(),
					Status:      service_v1.ShipmentStatus_SHIPMENT_STATUS_REGISTERED,
					LastUpdated: timestamppb.New(now),
					Location:    "Bangladesh",
				},
				{
					ShipmentId:  uuid.New().String(),
					Status:      service_v1.ShipmentStatus_SHIPMENT_STATUS_SHIPPED,
					LastUpdated: timestamppb.New(now),
					Location:    "China",
				},
			},
		}

		// Act
		_, err := client.Place(context.Background(), req)
		dbData, dbErr := dbFixture.GetAll()

		// Assert
		require.NoError(t, err)
		require.NoError(t, dbErr)

		require.Len(t, dbData, 2)
		require.Equal(t, req.Shipments[0].ShipmentId, dbData[0].ShipmentId)
		require.Equal(t, "Pending", dbData[0].Status)
		require.WithinDuration(t, req.Shipments[0].LastUpdated.AsTime(), dbData[0].LastUpdated, time.Millisecond)
		require.Equal(t, req.Shipments[0].Location, dbData[0].Location)
		require.WithinDuration(t, now.Add(24*time.Hour), dbData[0].ValidUntil, time.Millisecond)

		require.Equal(t, req.Shipments[1].ShipmentId, dbData[1].ShipmentId)
		require.Equal(t, "Completed", dbData[1].Status)
		require.WithinDuration(t, req.Shipments[1].LastUpdated.AsTime(), dbData[1].LastUpdated, time.Millisecond)
		require.Equal(t, req.Shipments[1].Location, dbData[1].Location)
		require.WithinDuration(t, now.Add(24*time.Hour), dbData[1].ValidUntil, time.Millisecond)
	})

	t.Run("List_shipments_by_ids", func(t *testing.T) {
		// Arrange
		dbFixture.ClearAll()
		now := time.Now().UTC()

		createReq := &service_v1.PlaceRequest{
			ValidUntil: timestamppb.New(now.Add(24 * time.Hour)),
			Shipments: []*service_v1.Shipment{
				{
					ShipmentId:  uuid.New().String(),
					Status:      service_v1.ShipmentStatus_SHIPMENT_STATUS_REGISTERED,
					LastUpdated: timestamppb.New(now),
					Location:    "Bangladesh",
				},
				{
					ShipmentId:  uuid.New().String(),
					Status:      service_v1.ShipmentStatus_SHIPMENT_STATUS_SHIPPED,
					LastUpdated: timestamppb.New(now),
					Location:    "China",
				},
			},
		}
		client.Place(context.Background(), createReq)

		req := &service_v1.ListRequest{
			ShipmentIds: []string{createReq.Shipments[0].ShipmentId, createReq.Shipments[1].ShipmentId},
		}

		// Act
		res, err := client.List(context.Background(), req)

		// Assert
		require.NoError(t, err)

		require.Len(t, res.Shipments, 2)
		idList := []string{createReq.Shipments[0].ShipmentId, createReq.Shipments[1].ShipmentId}
		require.Contains(t, idList, res.Shipments[0].ShipmentId)
		require.Contains(t, idList, res.Shipments[1].ShipmentId)
		require.WithinDuration(t, res.ValidUntil.AsTime(), createReq.ValidUntil.AsTime(), time.Millisecond)
	})
}
