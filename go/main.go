package main

import (
	"log"
	"net"

	"liaison_go/handlers"

	service_v1 "liaison_go/generated/service"

	"google.golang.org/grpc"
)

var (
	serverAddr = "localhost:5001"
)

func main() {
	lis, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	service_v1.RegisterTrackingServiceServer(grpcServer, handlers.NewTrackingHandler())
	grpcServer.Serve(lis)
}
