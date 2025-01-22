package main

import (
	"context"
	"log"

	client_go "liaison_go/client"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverAddr = "localhost:5002"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := client_go.NewOrderServiceClient(conn)

	var req = &client_go.Request{
		OrderId:   "38DE78BC-614D-44F2-BF2E-130F42224DD4",
		CreatedBy: "John Doe",
		Items: map[string]*client_go.OrderItem{
			"1111": {
				Code:     "1111",
				Quantity: 1,
				Price:    100.00,
			},
			"2222": {
				Code:     "2222",
				Quantity: 2,
				Price:    120.00,
			},
		},
	}

	var resp, orderErr = client.CreateOrder(context.Background(), req)
	if orderErr != nil {
		log.Printf("fail to create order: %v", err)
	}

	log.Printf("got the response: %v", resp)
}
