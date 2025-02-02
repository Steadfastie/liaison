package main

import (
	"fmt"
	"net"
	"time"

	"liaison_go/business"
	"liaison_go/handlers"
	infra "liaison_go/infrastructure"
	"liaison_go/persistence"

	service_v1 "liaison_go/generated/service"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"

	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	logLevelsFunc grpc_zap.CodeToLevel = func(code codes.Code) zapcore.Level {
		if code == codes.OK {
			return zapcore.InfoLevel
		}
		return zapcore.ErrorLevel
	}
	recoveryFunc = func(logger *zap.Logger) recovery.RecoveryHandlerFunc {
		return func(p interface{}) (err error) {
			logger.DPanic("Something went wrong", zap.Any("panic", p))
			return status.Errorf(codes.Internal, "Something went wrong")
		}
	}
)

const (
	systemName = "liaison" // represents the health of the system
)

func main() {
	// Logger intialization
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// Config intialization
	var conf infra.Config
	conf.Populate(logger)

	// MongoDB intialization
	db, disconnect := infra.InitialiseMongoDB(logger, &conf)
	defer disconnect()

	// Services intialization
	logger.Info("Initializing services")
	shipmentStore := persistence.NewShipmentStore(db, logger)               // 3rd layer
	shipmentTracker := business.NewShipmentTracker(shipmentStore, logger)   // 2nd layer
	trackingHandler := handlers.NewTrackingHandler(shipmentTracker, logger) // 1rd layer

	// gRPC intialization
	logger.Info("Starting gRPC server")

	zapOpts := []grpc_zap.Option{
		grpc_zap.WithLevels(logLevelsFunc),
		grpc_zap.WithDurationField(func(duration time.Duration) zapcore.Field {
			return zap.Int64("grpc.time_ns", duration.Nanoseconds())
		}),
	}
	grpc_zap.ReplaceGrpcLoggerV2(logger)

	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(recoveryFunc(logger)),
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Host.Port))
	if err != nil {
		logger.Fatal("failed to listen: %v", zap.Error(err))
	}
	opt := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.UnaryServerInterceptor(logger, zapOpts...),
			recovery.UnaryServerInterceptor(recoveryOpts...),
		),
	}
	grpcServer := grpc.NewServer(opt...)
	service_v1.RegisterTrackingServiceServer(grpcServer, trackingHandler)

	healthcheck := health.NewServer()
	healthgrpc.RegisterHealthServer(grpcServer, healthcheck)
	go infra.MonitorMongo(db.Client(), healthcheck, logger, systemName)

	grpcServer.Serve(lis)

	logger.Info("Finished serving")
}
