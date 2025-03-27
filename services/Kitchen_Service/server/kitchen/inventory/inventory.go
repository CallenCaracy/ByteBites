package inventory

import (
	"context"
	"fmt"
	"time"
	inventory "kitchen-Service/pb/inventory"
	"kitchen-Service/pb" // Update with the actual module path for your generated pb package
	"github.com/jackc/pgx/v5"
)

// InventoryServiceServer implements the InventoryService defined in the proto.
type InventoryServiceServer struct {
	pb.UnimplementedInventoryServiceServer
	DB *pgx.Conn
}

// CreateItem inserts a new inventory item and returns the created item.
func (s *InventoryServiceServer) CreateItem(ctx context.Context, req *pb.CreateItemRequest) (*pb.ItemResponse, error) {
	query := `
		INSERT INTO inventory_items (item_name, quantity, unit, low_stock_threshold, expiry_date, last_updated)
		VALUES ($1, $2, $3, $4, $5, NOW())
		RETURNING id, last_updated
	`
	var id string
	var lastUpdated time.Time
	err := s.DB.QueryRow(ctx, query, req.ItemName, req.Quantity, req.Unit, req.LowStockThreshold, req.ExpiryDate).
		Scan(&id, &lastUpdated)
	if err != nil {
		return nil, fmt.Errorf("failed to create item: %w", err)
	}

	item := &pb.Item{
		Id:               id,
		ItemName:         req.ItemName,
		Quantity:         req.Quantity,
		Unit:             req.Unit,
		LowStockThreshold: req.LowStockThreshold,
		ExpiryDate:       req.ExpiryDate,
		LastUpdated:      lastUpdated.Format(time.RFC3339),
	}

	return &pb.ItemResponse{Item: item}, nil
}

// GetItem retrieves a single inventory item by its ID.
func (s *InventoryServiceServer) GetItem(ctx context.Context, req *pb.GetItemRequest) (*pb.ItemResponse, error) {
	query := `
		SELECT id, item_name, quantity, unit, low_stock_threshold, expiry_date, last_updated
		FROM inventory_items
		WHERE id = $1
	`
	row := s.DB.QueryRow(ctx, query, req.Id)

	var item pb.Item
	var lastUpdated time.Time
	err := row.Scan(&item.Id, &item.ItemName, &item.Quantity, &item.Unit, &item.LowStockThreshold, &item.ExpiryDate, &lastUpdated)
	if err != nil {
		return nil, fmt.Errorf("failed to get item: %w", err)
	}
	item.LastUpdated = lastUpdated.Format(time.RFC3339)

	return &pb.ItemResponse{Item: &item}, nil
}

// UpdateItem updates an existing inventory item and returns the updated item.
func (s *InventoryServiceServer) UpdateItem(ctx context.Context, req *pb.UpdateItemRequest) (*pb.ItemResponse, error) {
	query := `
		UPDATE inventory_items
		SET item_name = $1,
		    quantity = $2,
		    unit = $3,
		    low_stock_threshold = $4,
		    expiry_date = $5,
		    last_updated = NOW()
		WHERE id = $6
		RETURNING last_updated
	`
	var lastUpdated time.Time
	err := s.DB.QueryRow(ctx, query, req.ItemName, req.Quantity, req.Unit, req.LowStockThreshold, req.ExpiryDate, req.Id).
		Scan(&lastUpdated)
	if err != nil {
		return nil, fmt.Errorf("failed to update item: %w", err)
	}

	updatedItem := &pb.Item{
		Id:               req.Id,
		ItemName:         req.ItemName,
		Quantity:         req.Quantity,
		Unit:             req.Unit,
		LowStockThreshold: req.LowStockThreshold,
		ExpiryDate:       req.ExpiryDate,
		LastUpdated:      lastUpdated.Format(time.RFC3339),
	}

	return &pb.ItemResponse{Item: updatedItem}, nil
}

// DeleteItem removes an inventory item by its ID.
func (s *InventoryServiceServer) DeleteItem(ctx context.Context, req *pb.DeleteItemRequest) (*pb.DeleteItemResponse, error) {
	query := `DELETE FROM inventory_items WHERE id = $1`
	ct, err := s.DB.Exec(ctx, query, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete item: %w", err)
	}
	success := ct.RowsAffected() > 0

	return &pb.DeleteItemResponse{Success: success}, nil
}

// ListItems retrieves all inventory items.
func (s *InventoryServiceServer) ListItems(ctx context.Context, req *pb.ListItemsRequest) (*pb.ListItemsResponse, error) {
	query := `
		SELECT id, item_name, quantity, unit, low_stock_threshold, expiry_date, last_updated
		FROM inventory_items
	`
	rows, err := s.DB.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to list items: %w", err)
	}
	defer rows.Close()

	var items []*pb.Item
	for rows.Next() {
		var item pb.Item
		var lastUpdated time.Time
		err := rows.Scan(&item.Id, &item.ItemName, &item.Quantity, &item.Unit, &item.LowStockThreshold, &item.ExpiryDate, &lastUpdated)
		if err != nil {
			return nil, fmt.Errorf("failed to scan item: %w", err)
		}
		item.LastUpdated = lastUpdated.Format(time.RFC3339)
		items = append(items, &item)
	}

	return &pb.ListItemsResponse{Items: items}, nil
}
