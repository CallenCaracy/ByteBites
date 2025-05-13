package main

import (
	"net"
	"os"

	"order-service/pb"
	order "order-service/server/orders"
	"order-service/utils"

	"google.golang.org/grpc"
)

func main() {
	log, err := utils.NewLogger()
	if err != nil {
		os.Exit(1)
	}

	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatal("Failed to listen: %v", err)
		os.Exit(1)
	}

	grpcServer := grpc.NewServer()

	orderService := &order.OrderServiceServer{}

	pb.RegisterOrderServiceServer(grpcServer, orderService)

	log.Info("Order Service running on port 50053...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to serve: %v", err)
		os.Exit(1)
	}
}
