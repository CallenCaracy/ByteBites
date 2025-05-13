package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"payment-service/pb"
)

type PaymentServiceServer struct {
	pb.UnimplementedPaymentServiceServer
}

const gqlEndpoint = "http://localhost:8080/query"

func (s *PaymentServiceServer) CreateTransaction(ctx context.Context, req *pb.CreateTransactionRequest) (*pb.TransactionResponse, error) {
	query := `
	mutation CreateTransaction(
		$amountPaid: Float!,
		$paymentMethod: String!,
		$transactionStatus: String!,
		$userID: String!,
		$orderID: String!
	) {
		createTransactionRecords(
			amountPaid: $amountPaid,
			paymentMethod: $paymentMethod,
			transactionStatus: $transactionStatus,
			userID: $userID,
			orderID: $orderID
		) {
			transactionID
			amountPaid
			paymentMethod
			transactionStatus
			userID
			orderID
			timestamp
		}
	}`

	variables := map[string]interface{}{
		"amountPaid":        req.AmountPaid,
		"paymentMethod":     req.PaymentMethod,
		"transactionStatus": "success", // hardcoded
		"userID":            req.UserID,
		"orderID":           req.OrderID,
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
				TransactionID     string  `json:"transactionID"`
				UserID            string  `json:"userID"`
				OrderID           string  `json:"orderID"`
				AmountPaid        float64 `json:"amountPaid"`
				PaymentMethod     string  `json:"paymentMethod"`
				TransactionStatus string  `json:"transactionStatus"`
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
