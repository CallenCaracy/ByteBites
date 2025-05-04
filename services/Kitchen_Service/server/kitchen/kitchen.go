package kitchen

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/CallenCaracy/ByteBites/services/Kitchen_Service/pb"
)

type KitchenServiceServer struct {
	pb.UnimplementedKitchenServiceServer
}

const gqlEndpoint = "http://localhost:8080/query"

func (s *KitchenServiceServer) CheckStock(ctx context.Context, req *pb.CheckStockRequest) (*pb.CheckStockResponse, error) {
	query := `
	query GetStock($id: ID!) {
		inventory(id: $id) {
			availableServings
		}
	}`

	payload := map[string]interface{}{
		"query": query,
		"variables": map[string]interface{}{
			"id": req.MenuItemId,
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	res, err := http.Post(gqlEndpoint, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response struct {
		Data struct {
			Inventory struct {
				AvailableServings int `json:"availableServings"`
			} `json:"inventory"`
		} `json:"data"`
	}

	if err := json.Unmarshal(data, &response); err != nil {
		return nil, err
	}

	if response.Data.Inventory.AvailableServings == 0 {
		return &pb.CheckStockResponse{
			Available:         false,
			AvailableQuantity: 0,
			Message:           "No inventory found or stock is zero",
		}, nil
	}

	currentStock := response.Data.Inventory.AvailableServings
	return &pb.CheckStockResponse{
		Available:         currentStock >= int(req.Quantity),
		AvailableQuantity: int32(currentStock),
		Message:           "Check stock successful",
	}, nil
}

func (s *KitchenServiceServer) DeductStock(ctx context.Context, req *pb.DeductStockRequest) (*pb.DeductStockResponse, error) {
	checkResp, err := s.CheckStock(ctx, &pb.CheckStockRequest{
		MenuItemId: req.MenuItemId,
		Quantity:   req.Quantity,
	})
	if err != nil {
		return nil, err
	}
	if !checkResp.Available {
		return &pb.DeductStockResponse{
			Success: false,
			Message: "Not enough stock",
		}, nil
	}

	newQuantity := int(checkResp.AvailableQuantity) - int(req.Quantity)

	mutation := `
	mutation UpdateInventory($id: ID!, $quantity: Int!) {
		updateInventory(id: $id, availableServings: $quantity) {
			id
			availableServings
		}
	}`

	payload := map[string]interface{}{
		"query": mutation,
		"variables": map[string]interface{}{
			"id":       req.MenuItemId,
			"quantity": newQuantity,
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	res, err := http.Post(gqlEndpoint, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var gqlResp struct {
		Data struct {
			UpdateInventory struct {
				ID                string `json:"id"`
				AvailableServings int    `json:"availableServings"`
			} `json:"updateInventory"`
		} `json:"data"`
	}

	if err := json.Unmarshal(respBody, &gqlResp); err != nil {
		return nil, err
	}

	return &pb.DeductStockResponse{
		Success:           true,
		Message:           "Stock deducted successfully",
		RemainingQuantity: int32(gqlResp.Data.UpdateInventory.AvailableServings),
	}, nil
}
