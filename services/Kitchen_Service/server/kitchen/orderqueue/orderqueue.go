package orderqueue

import (
	"context"
	"fmt"

	"kitchen-Service/pb"

	"github.com/jackc/pgx/v5"
)


// OrderQueueServiceServer implements the OrderQueueService defined in the proto.
type OrderQueueServiceServer struct {
	pb.UnimplementedOrderQueueServiceServer
	DB *pgx.Conn
}

// CreateOrderQueue creates a new order queue record.
func (s *OrderQueueServiceServer) CreateOrderQueue(ctx context.Context, req *pb.CreateOrderQueueRequest) (*pb.CreateOrderQueueResponse, error) {
	query := `
		INSERT INTO order_queue (order_id, status, priority, last_updated)
		VALUES ($1, $2, $3, NOW())
		RETURNING id
	`
	var id string
	if err := s.DB.QueryRow(ctx, query, req.OrderId, req.Status, req.Priority).Scan(&id); err != nil {
		return nil, fmt.Errorf("failed to create order queue entry: %w", err)
	}
	return &pb.CreateOrderQueueResponse{Id: id}, nil
}

// GetOrderQueue retrieves a single order queue record by its ID.
func (s *OrderQueueServiceServer) GetOrderQueue(ctx context.Context, req *pb.GetOrderQueueRequest) (*pb.GetOrderQueueResponse, error) {
	query := `
		SELECT id, order_id, status, priority, last_updated
		FROM order_queue
		WHERE id = $1
	`
	row := s.DB.QueryRow(ctx, query, req.Id)
	var oq pb.OrderQueue
	var lastUpdated time.Time
	if err := row.Scan(&oq.Id, &oq.OrderId, &oq.Status, &oq.Priority, &lastUpdated); err != nil {
		return nil, fmt.Errorf("failed to retrieve order queue entry: %w", err)
	}
	oq.LastUpdated = lastUpdated.Format(time.RFC3339)
	return &pb.GetOrderQueueResponse{OrderQueue: &oq}, nil
}

// UpdateOrderQueue updates the status and priority of an order queue record.
func (s *OrderQueueServiceServer) UpdateOrderQueue(ctx context.Context, req *pb.UpdateOrderQueueRequest) (*pb.UpdateOrderQueueResponse, error) {
	query := `
		UPDATE order_queue
		SET status = $1,
		    priority = $2,
		    last_updated = NOW()
		WHERE id = $3
	`
	res, err := s.DB.Exec(ctx, query, req.Status, req.Priority, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to update order queue entry: %w", err)
	}
	if res.RowsAffected() == 0 {
		return &pb.UpdateOrderQueueResponse{Success: false}, nil
	}
	return &pb.UpdateOrderQueueResponse{Success: true}, nil
}

// DeleteOrderQueue deletes an order queue record by its ID.
func (s *OrderQueueServiceServer) DeleteOrderQueue(ctx context.Context, req *pb.DeleteOrderQueueRequest) (*pb.DeleteOrderQueueResponse, error) {
	query := `
		DELETE FROM order_queue
		WHERE id = $1
	`
	res, err := s.DB.Exec(ctx, query, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete order queue entry: %w", err)
	}
	if res.RowsAffected() == 0 {
		return &pb.DeleteOrderQueueResponse{Success: false}, nil
	}
	return &pb.DeleteOrderQueueResponse{Success: true}, nil
}

// ListOrderQueue retrieves all order queue records.
func (s *OrderQueueServiceServer) ListOrderQueue(ctx context.Context, req *pb.ListOrderQueueRequest) (*pb.ListOrderQueueResponse, error) {
	query := `
		SELECT id, order_id, status, priority, last_updated
		FROM order_queue
	`
	rows, err := s.DB.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to list order queue entries: %w", err)
	}
	defer rows.Close()

	var orderQueues []*pb.OrderQueue
	for rows.Next() {
		var oq pb.OrderQueue
		var lastUpdated time.Time
		if err := rows.Scan(&oq.Id, &oq.OrderId, &oq.Status, &oq.Priority, &lastUpdated); err != nil {
			return nil, fmt.Errorf("failed to scan order queue row: %w", err)
		}
		oq.LastUpdated = lastUpdated.Format(time.RFC3339)
		orderQueues = append(orderQueues, &oq)
	}
	return &pb.ListOrderQueueResponse{OrderQueues: orderQueues}, nil
}

