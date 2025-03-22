package server

import (
	"context"
	"fmt"

	"order-service/pb"

	"time"

	"github.com/jackc/pgx/v5"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type OrderServiceServer struct {
	pb.UnimplementedOrderServiceServer
	DB                        *pgx.Conn
	GuestSessionServiceServer pb.GuestSessionServiceServer
}

func (s *OrderServiceServer) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	CheckGuestStatusResponse, errReq := s.GuestSessionServiceServer.CheckGuestStatus(ctx, &pb.CheckGuestStatusRequest{
		GuestId: req.GuestSessionId,
	})
	if errReq != nil {
		return nil, errReq
	}

	if CheckGuestStatusResponse.Status != "active" {
		return nil, fmt.Errorf("guest session is not active")
	}

	query := `INSERT INTO orders (guest_session_id, menu_item_id, quantity) VALUES ($1, $2, $3) RETURNING id`
	var orderID string
	err := s.DB.QueryRow(ctx, query, req.GuestSessionId, req.MenuItemId, req.Quantity).Scan(&orderID)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{Status: "Created"}, nil
}

func (s *OrderServiceServer) GetOrders(ctx context.Context, req *pb.GetOrdersRequest) (*pb.GetOrdersResponse, error) {
	rows, err := s.DB.Query(ctx, `SELECT id, guest_session_id, menu_item_id, quantity, order_time FROM orders WHERE guest_session_id=$1`, req.GuestSessionId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*pb.Order
	for rows.Next() {
		var order pb.Order
		var orderTime time.Time
		err := rows.Scan(&order.Id, &order.GuestSessionId, &order.MenuItemId, &order.Quantity, &orderTime)
		if err != nil {
			return nil, err
		}
		order.OrderTime = timestamppb.New(orderTime)
		orders = append(orders, &order)
	}
	return &pb.GetOrdersResponse{Orders: orders}, nil
}

func (s *OrderServiceServer) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
	query := `UPDATE orders SET quantity=$1 WHERE id=$2 AND guest_session_id=$3`
	_, err := s.DB.Exec(ctx, query, req.Quantity, req.OrderId, req.GuestSessionId)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateOrderResponse{Status: "Updated"}, nil
}

func (s *OrderServiceServer) DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error) {
	query := `DELETE FROM orders WHERE id=$1 AND guest_session_id=$2`
	_, err := s.DB.Exec(ctx, query, req.OrderId, req.GuestSessionId)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteOrderResponse{Status: "Deleted"}, nil
}
