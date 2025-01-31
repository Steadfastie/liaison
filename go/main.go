package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"liaison_go/business"
	"liaison_go/handlers"
	"liaison_go/persistence"

	service_v1 "liaison_go/generated/service"

	"github.com/spf13/viper"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

const (
	prefix = "LIAISON"
)

var (
	customFunc grpc_zap.CodeToLevel = func(code codes.Code) zapcore.Level {
		if code == codes.OK {
			return zapcore.InfoLevel
		}
		return zapcore.ErrorLevel
	}
)

func main() {
	// Logger intialization
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// Config intialization
	logger.Info("Reading configuration")
	viper.SetEnvPrefix(prefix)
	viper.AutomaticEnv()

	// inspired by https://github.com/spf13/viper/issues/761#issuecomment-1578931559
	for _, e := range os.Environ() {
		split := strings.Split(e, "=")
		k := split[0]
		if strings.HasPrefix(k, prefix) {
			name := strings.Join(strings.Split(k, "_")[1:], ".")
			// Explicit Set has the highest priority
			viper.Set(name, split[1])
		}
	}

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Warn("Config file not found, using environment variables")
		} else {
			panic(fmt.Errorf("config file error %w", err))
		}
	}

	var conf Config
	err = viper.Unmarshal(&conf)
	if err != nil {
		panic(fmt.Errorf("config unmarshalling error %w", err))
	}
	logger.Info("Configuration read", zap.Any("config", conf))

	// MongoDB intialization
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
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			panic(fmt.Errorf("mongo client disconnection error %w", err))
		}
	}()
	db := client.Database(conf.MongoSettings.Database)

	// Services intialization
	logger.Info("Initializing services")
	shipmentStore := persistence.NewShipmentStore(db, logger)               // 3rd layer
	shipmentTracker := business.NewShipmentTracker(shipmentStore, logger)   // 2nd layer
	trackingHandler := handlers.NewTrackingHandler(shipmentTracker, logger) // 1rd layer

	// gRPC intialization
	logger.Info("Starting gRPC server")

	opts := []grpc_zap.Option{
		grpc_zap.WithLevels(customFunc),
		grpc_zap.WithDurationField(func(duration time.Duration) zapcore.Field {
			return zap.Int64("grpc.time_ns", duration.Nanoseconds())
		}),
	}
	grpc_zap.ReplaceGrpcLoggerV2(logger)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Host.Port))
	if err != nil {
		logger.Fatal("failed to listen: %v", zap.Error(err))
	}
	opt := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.UnaryServerInterceptor(logger, opts...),
		),
	}
	grpcServer := grpc.NewServer(opt...)
	service_v1.RegisterTrackingServiceServer(grpcServer, trackingHandler)
	grpcServer.Serve(lis)
	logger.Info("Finished serving")
}
