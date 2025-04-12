package inventory;

import {
	"context"
	"fmt"

	"Kitchen_Service/pb"

	"github.com/jackc/pgx/v5"
}

type OrderQueueServer struct {
	pb.UnimplementedOrderQueueServiceServer
	DB *pgx.Conn
}

func (s *OrderQueueServer) CreateOrder(ctx context.Context, req *pb.OrderCreateRequest) (*pb.OrderCreateResponse, error) {
	id := uuid.New().String()
	status := req.Status
	if status == pb.Status_STATUS_UNSPECIFIED {
		status = pb.Status_PREPARING
	}
	priority := req.Priority
	if priority == 0 {
		priority = 1
	}

	query := `
		INSERT INTO order_queue (id, order_id, status, priority, last_updated)
		VALUES ($1, $2, $3, $4, NOW())
		RETURNING last_updated
	`
	var lastUpdated time.Time
	err := s.DB.QueryRow(ctx, query, id, req.OrderId, status.String(), priority).Scan(&lastUpdated)
	if err != nil {
		return nil, fmt.Errorf("failed to create order: %v", err)
	}

	order := &pb.Order{
		Id:          id,
		OrderId:     req.OrderId,
		Status:      status,
		Priority:    priority,
		LastUpdated: timestamppb.New(lastUpdated),
	}
	return &pb.OrderCreateResponse{Order: order}, nil
}

func (s *OrderQueueServer) GetOrder(ctx context.Context, req *pb.OrderGetRequest) (*pb.OrderGetResponse, error) {
	query := `
		SELECT order_id, status, priority, last_updated
		FROM order_queue
		WHERE id = $1
	`
	var orderId, status string
	var priority int32
	var lastUpdated time.Time
	err := s.DB.QueryRow(ctx, query, req.Id).Scan(&orderId, &status, &priority, &lastUpdated)
	if err != nil {
		return nil, fmt.Errorf("failed to get order: %v", err)
	}

	order := &pb.Order{
		Id:          req.Id,
		OrderId:     orderId,
		Status:      pb.Status(pb.Status_value[strings.ToUpper(status)]),
		Priority:    priority,
		LastUpdated: timestamppb.New(lastUpdated),
	}
	return &pb.OrderGetResponse{Order: order}, nil
}

func (s *OrderQueueServer) UpdateOrder(ctx context.Context, req *pb.OrderUpdateRequest) (*pb.OrderUpdateResponse, error) {
	order := req.Order
	query := `
		UPDATE order_queue
		SET order_id = $1, status = $2, priority = $3, last_updated = NOW()
		WHERE id = $4
		RETURNING last_updated
	`
	var lastUpdated time.Time
	err := s.DB.QueryRow(ctx, query, order.OrderId, order.Status.String(), order.Priority, order.Id).Scan(&lastUpdated)
	if err != nil {
		return nil, fmt.Errorf("failed to update order: %v", err)
	}

	order.LastUpdated = timestamppb.New(lastUpdated)
	return &pb.OrderUpdateResponse{Order: order}, nil
}

func (s *OrderQueueServer) DeleteOrder(ctx context.Context, req *pb.OrderDeleteRequest) (*pb.OrderDeleteResponse, error) {
	query := `DELETE FROM order_queue WHERE id = $1`
	tag, err := s.DB.Exec(ctx, query, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete order: %v", err)
	}
	if tag.RowsAffected() == 0 {
		return nil, fmt.Errorf("no order found with id: %s", req.Id)
	}
	return &pb.OrderDeleteResponse{}, nil
}

func (s *OrderQueueServer) ListOrders(ctx context.Context, req *pb.OrderListRequest) (*pb.OrderListResponse, error) {
	orders := []*pb.Order{}
	query := `
		SELECT id, order_id, status, priority, last_updated
		FROM order_queue
	`
	args := []interface{}{}
	if req.Status != pb.Status_STATUS_UNSPECIFIED {
		query += " WHERE status = $1"
		args = append(args, req.Status.String())
	}

	rows, err := s.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to list orders: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, orderId, status string
		var priority int32
		var lastUpdated time.Time

		if err := rows.Scan(&id, &orderId, &status, &priority, &lastUpdated); err != nil {
			return nil, fmt.Errorf("error scanning order row: %v", err)
		}

		orders = append(orders, &pb.Order{
			Id:          id,
			OrderId:     orderId,
			Status:      pb.Status(pb.Status_value[strings.ToUpper(status)]),
			Priority:    priority,
			LastUpdated: timestamppb.New(lastUpdated),
		})
	}

	return &pb.OrderListResponse{
		Orders: orders,
		// Pagination not implemented
	}, nil
}