package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"order-service/pb"
)

type OrderServiceServer struct {
	pb.UnimplementedOrderServiceServer
}

const gqlEndpoint = "http://localhost:8080/query"

func (s *OrderServiceServer) CreateCart(ctx context.Context, req *pb.CreateCartRequest) (*pb.CartResponse, error) {
	query := `
	mutation CreateCart($userID: String!) {
		createCart(userID: $userID) {
			id
			userID
			createdAt
			updatedAt
		}
	}`

	payload := map[string]interface{}{
		"query": query,
		"variables": map[string]interface{}{
			"userID": req.UserID,
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal GraphQL payload: %v", err)
	}

	res, err := http.Post(gqlEndpoint, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to make GraphQL request: %v", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var response struct {
		Data struct {
			CreateCart struct {
				ID        string `json:"id"`
				UserID    string `json:"userID"`
				CreatedAt string `json:"createdAt"`
				UpdatedAt string `json:"updatedAt"`
			} `json:"createCart"`
		} `json:"data"`
		Errors []struct {
			Message string `json:"message"`
		} `json:"errors"`
	}

	if err := json.Unmarshal(data, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal GraphQL response: %v", err)
	}

	if len(response.Errors) > 0 {
		return nil, fmt.Errorf("GraphQL error: %s", response.Errors[0].Message)
	}

	cart := response.Data.CreateCart
	return &pb.CartResponse{
		Id:        cart.ID,
		UserID:    cart.UserID,
		CreatedAt: cart.CreatedAt,
		UpdatedAt: cart.UpdatedAt,
	}, nil
}
