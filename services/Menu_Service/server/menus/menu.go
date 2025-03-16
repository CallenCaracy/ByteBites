package menu

import (
	"context"
	"fmt"

	"Menu_Service/pb"

	"github.com/jackc/pgx/v5"
)

type MenuServiceServer struct {
	pb.UnimplementedMenuServiceServer
	DB *pgx.Conn
}

func (s *MenuServiceServer) GetMenuItems(ctx context.Context, req *pb.GetMenuItemsRequest) (*pb.GetMenuItemsResponse, error) {
	rows, err := s.DB.Query(ctx, `SELECT id, name, image_url, description, price FROM menu_items`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var menuItems []*pb.MenuItem
	for rows.Next() {
		var menuItem pb.MenuItem
		err := rows.Scan(&menuItem.Id, &menuItem.Name, &menuItem.ImageUrl, &menuItem.Description, &menuItem.Price)
		if err != nil {
			return nil, err
		}
		menuItems = append(menuItems, &menuItem)
	}
	return &pb.GetMenuItemsResponse{MenuItems: menuItems}, nil
}

func (s *MenuServiceServer) CreateMenuItem(ctx context.Context, req *pb.CreateMenuItemRequest) (*pb.CreateMenuItemResponse, error) {
	query := `INSERT INTO menu_items (name, image_url, description, price, item_status, created_at) VALUES ($1, $2, $3, $4, $5, NOW()) RETURNING id`
	var menuItemID string
	err := s.DB.QueryRow(ctx, query, req.Name, req.ImageUrl, req.Description, req.Price, req.ItemStatus).Scan(&menuItemID)
	if err != nil {
		return nil, err
	}
	return &pb.CreateMenuItemResponse{Status: "Created"}, nil
}

func (s *MenuServiceServer) DeleteMenuItem(ctx context.Context, req *pb.DeleteMenuItemRequest) (*pb.DeleteMenuItemResponse, error) {
	query := `DELETE FROM menu_items WHERE id=$1`
	_, err := s.DB.Exec(ctx, query, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteMenuItemResponse{Status: "Deleted"}, nil
}

func (s *MenuServiceServer) UpdateMenuItem(ctx context.Context, req *pb.UpdateMenuItemRequest) (*pb.UpdateMenuItemResponse, error) {
	query := `UPDATE menu_items SET `
	args := []interface{}{}
	argIdx := 1

	if req.Name != nil {
		query += fmt.Sprintf("name=$%d, ", argIdx)
		args = append(args, req.Name)
		argIdx++
	}
	if req.ImageUrl != nil {
		query += fmt.Sprintf("image_url=$%d, ", argIdx)
		args = append(args, req.ImageUrl)
		argIdx++
	}
	if req.Description != nil {
		query += fmt.Sprintf("description=$%d, ", argIdx)
		args = append(args, req.Description)
		argIdx++
	}
	if req.Price != nil {
		query += fmt.Sprintf("price=$%d, ", argIdx)
		args = append(args, req.Price)
		argIdx++
	}
	if req.ItemStatus != nil {
		query += fmt.Sprintf("item_status=$%d, ", argIdx)
		args = append(args, req.ItemStatus)
		argIdx++
	}

	query += fmt.Sprintf("updated_at=NOW() WHERE id=$%d", argIdx)
	args = append(args, req.Id)

	_, err := s.DB.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateMenuItemResponse{
		Id:     req.Id,
		Status: "Updated",
	}, nil
}
