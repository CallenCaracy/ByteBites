package menu_list

import (
	"context"
	"fmt"
	"time"

	"Menu_Service/pb"
	"Menu_Service/utils"

	"github.com/jackc/pgx/v5"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MenuServiceServer struct {
	pb.UnimplementedMenuServiceServer
	DB     *pgx.Conn
	Logger *utils.Logger
}

// CreateMenuItem handles adding a new menu item
func (s *MenuServiceServer) CreateMenuItem(ctx context.Context, req *pb.CreateMenuItemRequest) (*pb.CreateMenuItemResponse, error) {
	// ✅ Check for nil database and logger
	if s.DB == nil {
		return nil, fmt.Errorf("database connection is nil")
	}
	if s.Logger == nil {
		return nil, fmt.Errorf("logger is nil")
	}

	// ✅ Validate request payload
	if req == nil {
		s.Logger.Error("Received nil request")
		return nil, fmt.Errorf("invalid request: request is nil")
	}

	// ✅ Log incoming request details
	s.Logger.Info("Creating new menu item: %s", req.Name)

	// SQL Query for inserting data
	query := `
		INSERT INTO menu_list (name, description, price, category, availability_status, image_url, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())
		RETURNING id, name, description, price, category, availability_status, image_url, created_at, updated_at
	`

	// Initialize variables for the returned menu item
	var menuItem pb.MenuItem
	var createdAt, updatedAt time.Time

	// Execute the query
	row := s.DB.QueryRow(ctx, query,
		req.Name,
		req.Description,
		req.Price,
		req.Category,
		req.AvailabilityStatus,
		req.ImageUrl,
	)

	// Scan the returned row into the menu item
	err := row.Scan(
		&menuItem.Id,
		&menuItem.Name,
		&menuItem.Description,
		&menuItem.Price,
		&menuItem.Category,
		&menuItem.AvailabilityStatus,
		&menuItem.ImageUrl,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		s.Logger.Error("Failed to create menu item: %v", err)
		return nil, fmt.Errorf("failed to create menu item: %v", err)
	}

	// Assign timestamps using protobuf Timestamp type
	menuItem.CreatedAt = timestamppb.New(createdAt)
	menuItem.UpdatedAt = timestamppb.New(updatedAt)

	// ✅ Return the response
	return &pb.CreateMenuItemResponse{MenuItem: &menuItem}, nil
}
