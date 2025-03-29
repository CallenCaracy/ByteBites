package transaction_records

import (
	"context"
	"fmt"
	"strings"

	"github.com/CallenCaracy/ByteBites/services/User_Service/pb"

	"github.com/jackc/pgx/v5"
)

type TransactionServiceServer struct {
	pb.UnimplementedTransactionServiceServer
	DB *pgx.Conn
}

// CreateTransaction inserts a new record into transaction_records.
func (s *TransactionServiceServer) CreateTransaction(ctx context.Context, req *pb.CreateTransactionRequest) (*pb.TransactionResponse, error) {
	query := `
		INSERT INTO transaction_records (user_id, transaction_id, amount, status, timestamp)
		VALUES ($1, $2, $3, $4, NOW())
		RETURNING id, timestamp
	`
	var id string
	var timestamp string
	err := s.DB.QueryRow(ctx, query, req.UserId, req.TransactionId, req.Amount, req.Status).Scan(&id, &timestamp)
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction: %v", err)
	}

	transaction := &pb.Transaction{
		Id:            id,
		UserId:        req.UserId,
		TransactionId: req.TransactionId,
		Amount:        req.Amount,
		Status:        req.Status,
		Timestamp:     timestamp,
	}

	return &pb.TransactionResponse{Transaction: transaction}, nil
}

// GetTransaction retrieves a transaction record by its ID.
func (s *TransactionServiceServer) GetTransaction(ctx context.Context, req *pb.GetTransactionRequest) (*pb.TransactionResponse, error) {
	query := `
		SELECT user_id, transaction_id, amount, status, timestamp 
		FROM transaction_records 
		WHERE id = $1
	`
	var userId, transactionId, status, timestamp string
	var amount float64
	err := s.DB.QueryRow(ctx, query, req.Id).Scan(&userId, &transactionId, &amount, &status, &timestamp)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction: %v", err)
	}

	transaction := &pb.Transaction{
		Id:            req.Id,
		UserId:        userId,
		TransactionId: transactionId,
		Amount:        amount,
		Status:        status,
		Timestamp:     timestamp,
	}
	return &pb.TransactionResponse{Transaction: transaction}, nil
}

// UpdateTransaction updates an existing transaction record.
// This implementation builds the query dynamically based on which fields are provided.
func (s *TransactionServiceServer) UpdateTransaction(ctx context.Context, req *pb.UpdateTransactionRequest) (*pb.TransactionResponse, error) {
	updates := []string{}
	args := []interface{}{req.Id} // first parameter is always the transaction id
	argIdx := 2

	// Note: We assume optional fields are pointers.
	if req.UserId != nil {
		updates = append(updates, fmt.Sprintf("user_id = $%d", argIdx))
		args = append(args, *req.UserId)
		argIdx++
	}
	if req.TransactionId != nil {
		updates = append(updates, fmt.Sprintf("transaction_id = $%d", argIdx))
		args = append(args, *req.TransactionId)
		argIdx++
	}
	if req.Amount != nil {
		updates = append(updates, fmt.Sprintf("amount = $%d", argIdx))
		args = append(args, *req.Amount)
		argIdx++
	}
	if req.Status != nil {
		updates = append(updates, fmt.Sprintf("status = $%d", argIdx))
		args = append(args, *req.Status)
		argIdx++
	}

	if len(updates) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	// Build the query. We update the timestamp as well.
	query := fmt.Sprintf("UPDATE transaction_records SET %s, timestamp = NOW() WHERE id = $1 RETURNING user_id, transaction_id, amount, status, timestamp", strings.Join(updates, ", "))

	var userId, transactionId, status, timestamp string
	var amount float64
	err := s.DB.QueryRow(ctx, query, args...).Scan(&userId, &transactionId, &amount, &status, &timestamp)
	if err != nil {
		return nil, fmt.Errorf("failed to update transaction: %v", err)
	}

	transaction := &pb.Transaction{
		Id:            req.Id,
		UserId:        userId,
		TransactionId: transactionId,
		Amount:        amount,
		Status:        status,
		Timestamp:     timestamp,
	}
	return &pb.TransactionResponse{Transaction: transaction}, nil
}

// DeleteTransaction deletes a transaction record by its ID.
func (s *TransactionServiceServer) DeleteTransaction(ctx context.Context, req *pb.DeleteTransactionRequest) (*pb.DeleteTransactionResponse, error) {
	query := `DELETE FROM transaction_records WHERE id = $1`
	tag, err := s.DB.Exec(ctx, query, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete transaction: %v", err)
	}
	if tag.RowsAffected() == 0 {
		return &pb.DeleteTransactionResponse{Success: false}, fmt.Errorf("no transaction found with id: %s", req.Id)
	}
	return &pb.DeleteTransactionResponse{Success: true}, nil
}
