package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/CallenCaracy/ByteBites/services/Payment_Service/pb"
)

type PaymentServiceServer struct {
	pb.UnimplementedPaymentServiceServer
}

const gqlEndpoint = "http://localhost:8080/query"

func (s *PaymentServiceServer) CreateTransaction(ctx context.Context, req *pb.CreateTransactionRequest) (*pb.TransactionResponse, error) {
	query := `
	mutation CreateTransaction(
		$amount_paid: Float!,
		$payment_method: String!,
		$transaction_status: String!,
		$user_id: ID!,
		$order_id: ID!
	) {
		createTransactionRecords(
			amount_paid: $amount_paid,
			payment_method: $payment_method,
			transaction_status: $transaction_status,
			user_id: $user_id,
			order_id: $order_id
		) {
			transaction_id
			amount_paid
			payment_method
			transaction_status
			user_id
			order_id
			timestamp
		}
	}`

	variables := map[string]interface{}{
		"amount_paid":        req.AmountPaid,
		"payment_method":     req.PaymentMethod,
		"transaction_status": "success", // still hardcoded, still victorious
		"user_id":            req.UserID,
		"order_id":           req.OrderID,
	}

	payload := map[string]interface{}{
		"query":     query,
		"variables": variables,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	res, err := http.Post(gqlEndpoint, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("GraphQL request failed: %w", err)
	}
	defer res.Body.Close()

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read GraphQL response: %w", err)
	}

	var gqlResp struct {
		Data struct {
			CreateTransactionRecords struct {
				TransactionID     string  `json:"transaction_id"`
				UserID            string  `json:"user_id"`
				OrderID           string  `json:"order_id"`
				AmountPaid        float64 `json:"amount_paid"`
				PaymentMethod     string  `json:"payment_method"`
				TransactionStatus string  `json:"transaction_status"`
				Timestamp         string  `json:"timestamp"`
			} `json:"createTransactionRecords"`
		} `json:"data"`
		Errors []struct {
			Message string `json:"message"`
		} `json:"errors"`
	}

	if err := json.Unmarshal(respBody, &gqlResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal GraphQL response: %w", err)
	}

	if len(gqlResp.Errors) > 0 {
		return nil, fmt.Errorf("GraphQL error: %s", gqlResp.Errors[0].Message)
	}

	tx := gqlResp.Data.CreateTransactionRecords

	return &pb.TransactionResponse{
		TransactionID:     tx.TransactionID,
		UserID:            tx.UserID,
		OrderID:           tx.OrderID,
		AmountPaid:        float32(tx.AmountPaid),
		PaymentMethod:     tx.PaymentMethod,
		TransactionStatus: tx.TransactionStatus,
		Timestamp:         tx.Timestamp,
	}, nil
}
