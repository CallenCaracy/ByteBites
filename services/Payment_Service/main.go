package main

import (
	"net"
	"os"

	"payment-service/pb"
	payment "payment-service/server/payments"
	"payment-service/utils"

	"google.golang.org/grpc"
)

func main() {
	log, err := utils.NewLogger()
	if err != nil {
		os.Exit(1)
	}

	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatal("Failed to listen: %v", err)
		os.Exit(1)
	}

	grpcServer := grpc.NewServer()

	paymentService := &payment.PaymentServiceServer{}

	pb.RegisterPaymentServiceServer(grpcServer, paymentService)

	log.Info("Order Service running on port 50054...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to serve: %v", err)
		os.Exit(1)
	}
}
