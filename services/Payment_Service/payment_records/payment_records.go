package payment

import (
	"context"
	"database/sql"
	"log"
	"time"

	pb "Payment_Service/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// PaymentServer implements the PaymentServiceServer interface
type PaymentServer struct {
	pb.UnimplementedPaymentServiceServer
	db          *sql.DB          // Database connection
	orderClient pb.OrderServiceClient // Client to interact with OrderService
}

// NewPaymentServer creates a new PaymentServer instance
func NewPaymentServer(db *sql.DB, orderClient pb.OrderServiceClient) *PaymentServer {
	return &PaymentServer{
		db:          db,
		orderClient: orderClient,
	}
}

// CreatePayment creates a new payment record
func (s *PaymentServer) CreatePayment(ctx context.Context, req *pb.Payment) (*pb.Payment, error) {
	// Validate payment method
	validMethods := map[string]bool{
		"Cash": true, "Credit Card": true, "Apple Pay": true, "Google Pay": true, "GCash": true,
	}
	if req.PaymentMethod != "" && !validMethods[req.PaymentMethod] {
		return nil, status.Errorf(codes.InvalidArgument, "invalid payment method: %s", req.PaymentMethod)
	}

	// Validate status (default to "Pending" if not provided)
	validStatuses := map[string]bool{
		"Pending": true, "Completed": true, "Failed": true,
	}
	if req.Status == "" {
		req.Status = "Pending"
	} else if !validStatuses[req.Status] {
		return nil, status.Errorf(codes.InvalidArgument, "invalid status: %s", req.Status)
	}

	// Validate amount >= 0
	if req.Amount < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "amount cannot be negative: %v", req.Amount)
	}

	// Get order to validate amount against total_amount
	orderResp, err := s.orderClient.GetOrder(ctx, &pb.OrderRequest{OrderId: req.OrderId})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get order: %v", err)
	}
	if req.Amount != orderResp.TotalAmount {
		return nil, status.Errorf(codes.InvalidArgument, "payment amount (%v) must match order total (%v)", req.Amount, orderResp.TotalAmount)
	}

	// Insert payment into database
	query := `
		INSERT INTO payments (order_id, amount, payment_method, status, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING payment_id, created_at`
	var paymentID int32
	var createdAt time.Time
	err = s.db.QueryRowContext(ctx, query,
		req.OrderId,
		req.Amount,
		req.PaymentMethod,
		req.Status,
		time.Now(),
	).Scan(&paymentID, &createdAt)
	if err != nil {
		log.Printf("Failed to create payment: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to create payment: %v", err)
	}

	// Construct response
	return &pb.Payment{
		PaymentId:     paymentID,
		OrderId:       req.OrderId,
		Amount:        req.Amount,
		PaymentMethod: req.PaymentMethod,
		Status:        req.Status,
		CreatedAt:     &timestamppb.Timestamp{Seconds: createdAt.Unix(), Nanos: int32(createdAt.Nanosecond())},
	}, nil
}

// GetPayment retrieves a payment by ID
func (s *PaymentServer) GetPayment(ctx context.Context, req *pb.PaymentRequest) (*pb.Payment, error) {
	query := `
		SELECT payment_id, order_id, amount, payment_method, status, created_at
		FROM payments
		WHERE payment_id = $1`
	var p pb.Payment
	var createdAt time.Time
	err := s.db.QueryRowContext(ctx, query, req.PaymentId).Scan(
		&p.PaymentId,
		&p.OrderId,
		&p.Amount,
		&p.PaymentMethod,
		&p.Status,
		&createdAt,
	)
	if err == sql.ErrNoRows {
		return nil, status.Errorf(codes.NotFound, "payment not found: %d", req.PaymentId)
	}
	if err != nil {
		log.Printf("Failed to get payment: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get payment: %v", err)
	}

	p.CreatedAt = &timestamppb.Timestamp{Seconds: createdAt.Unix(), Nanos: int32(createdAt.Nanosecond())}
	return &p, nil
}

// ListPaymentsByOrder streams all payments for a given order
func (s *PaymentServer) ListPaymentsByOrder(req *pb.OrderRequest, stream pb.PaymentService_ListPaymentsByOrderServer) error {
	query := `
		SELECT payment_id, order_id, amount, payment_method, status, created_at
		FROM payments
		WHERE order_id = $1`
	rows, err := s.db.QueryContext(stream.Context(), query, req.OrderId)
	if err != nil {
		log.Printf("Failed to list payments: %v", err)
		return status.Errorf(codes.Internal, "failed to list payments: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var p pb.Payment
		var createdAt time.Time
		err = rows.Scan(
			&p.PaymentId,
			&p.OrderId,
			&p.Amount,
			&p.PaymentMethod,
			&p.Status,
			&createdAt,
		)
		if err != nil {
			log.Printf("Failed to scan payment: %v", err)
			return status.Errorf(codes.Internal, "failed to scan payment: %v", err)
		}
		p.CreatedAt = &timestamppb.Timestamp{Seconds: createdAt.Unix(), Nanos: int32(createdAt.Nanosecond())}
		if err := stream.Send(&p); err != nil {
			log.Printf("Failed to send payment: %v", err)
			return err
		}
	}

	return nil
}

// UpdatePayment updates an existing payment
func (s *PaymentServer) UpdatePayment(ctx context.Context, req *pb.Payment) (*pb.Payment, error) {
	// Validate payment method
	validMethods := map[string]bool{
		"Cash": true, "Credit Card": true, "Apple Pay": true, "Google Pay": true, "GCash": true,
	}
	if req.PaymentMethod != "" && !validMethods[req.PaymentMethod] {
		return nil, status.Errorf(codes.InvalidArgument, "invalid payment method: %s", req.PaymentMethod)
	}

	// Validate status
	validStatuses := map[string]bool{
		"Pending": true, "Completed": true, "Failed": true,
	}
	if !validStatuses[req.Status] {
		return nil, status.Errorf(codes.InvalidArgument, "invalid status: %s", req.Status)
	}

	// Validate amount >= 0
	if req.Amount < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "amount cannot be negative: %v", req.Amount)
	}

	// Validate amount against order total
	orderResp, err := s.orderClient.GetOrder(ctx, &pb.OrderRequest{OrderId: req.OrderId})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get order: %v", err)
	}
	if req.Amount != orderResp.TotalAmount {
		return nil, status.Errorf(codes.InvalidArgument, "payment amount (%v) must match order total (%v)", req.Amount, orderResp.TotalAmount)
	}

	// Update payment in database
	query := `
		UPDATE payments
		SET order_id = $1, amount = $2, payment_method = $3, status = $4
		WHERE payment_id = $5
		RETURNING created_at`
	var createdAt time.Time
	err = s.db.QueryRowContext(ctx, query,
		req.OrderId,
		req.Amount,
		req.PaymentMethod,
		req.Status,
		req.PaymentId,
	).Scan(&createdAt)
	if err == sql.ErrNoRows {
		return nil, status.Errorf(codes.NotFound, "payment not found: %d", req.PaymentId)
	}
	if err != nil {
		log.Printf("Failed to update payment: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to update payment: %v", err)
	}

	// Construct response
	req.CreatedAt = &timestamppb.Timestamp{Seconds: createdAt.Unix(), Nanos: int32(createdAt.Nanosecond())}
	return req, nil
}

// DeletePayment deletes a payment
func (s *PaymentServer) DeletePayment(ctx context.Context, req *pb.PaymentRequest) (*emptypb.Empty, error) {
	query := `DELETE FROM payments WHERE payment_id = $1`
	result, err := s.db.ExecContext(ctx, query, req.PaymentId)
	if err != nil {
		log.Printf("Failed to delete payment: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to delete payment: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to check deletion: %v", err)
	}
	if rowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "payment not found: %d", req.PaymentId)
	}

	return &emptypb.Empty{}, nil
}