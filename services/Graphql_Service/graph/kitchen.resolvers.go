package graph

import (
	"Graphql_Service/graph/model"
	"context"
)

// CreateInventory
func (r *mutationResolver) CreateInventory(ctx context.Context, itemName string, quantity int32, unit string, lowStockThreshold *int32, expiryDate *string) (*model.Inventory, error) {
	query := `INSERT INTO public.inventory (item_name, quantity, unit, low_stock_threshold, expiry_date, last_updated)
	          VALUES ($1, $2, $3, COALESCE($4, 5), COALESCE($5::timestamptz, NULL), NOW())
	          RETURNING id, item_name, quantity, unit, low_stock_threshold, expiry_date, last_updated`

	var inv model.Inventory
	err := r.Resolver.DB7.QueryRow(query, itemName, quantity, unit, lowStockThreshold, expiryDate).
		Scan(&inv.ID, &inv.ItemName, &inv.Quantity, &inv.Unit, &inv.LowStockThreshold, &inv.ExpiryDate, &inv.LastUpdated)
	if err != nil {
		return nil, err
	}
	return &inv, nil
}

// UpdateInventory
func (r *mutationResolver) UpdateInventory(ctx context.Context, id string, itemName *string, quantity *int32, unit *string, lowStockThreshold *int32, expiryDate *string) (*model.Inventory, error) {
	query := `UPDATE public.inventory SET 
	          item_name = COALESCE($1, item_name), 
	          quantity = COALESCE($2, quantity),
	          unit = COALESCE($3, unit),
	          low_stock_threshold = COALESCE($4, low_stock_threshold),
	          expiry_date = COALESCE($5::timestamptz, expiry_date),
	          last_updated = NOW()
	          WHERE id = $6
	          RETURNING id, item_name, quantity, unit, low_stock_threshold, expiry_date, last_updated`

	var inv model.Inventory
	err := r.Resolver.DB7.QueryRow(query, itemName, quantity, unit, lowStockThreshold, expiryDate, id).
		Scan(&inv.ID, &inv.ItemName, &inv.Quantity, &inv.Unit, &inv.LowStockThreshold, &inv.ExpiryDate, &inv.LastUpdated)
	if err != nil {
		return nil, err
	}
	return &inv, nil
}

// DeleteInventory
func (r *mutationResolver) DeleteInventory(ctx context.Context, id string) (bool, error) {
	query := `DELETE FROM public.inventory WHERE id = $1`
	_, err := r.Resolver.DB7.Exec(query, id)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Inventories - Get all
func (r *queryResolver) Inventories(ctx context.Context) ([]*model.Inventory, error) {
	query := `SELECT id, item_name, quantity, unit, low_stock_threshold, expiry_date, last_updated FROM public.inventory`
	rows, err := r.Resolver.DB7.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*model.Inventory
	for rows.Next() {
		var inv model.Inventory
		err := rows.Scan(&inv.ID, &inv.ItemName, &inv.Quantity, &inv.Unit, &inv.LowStockThreshold, &inv.ExpiryDate, &inv.LastUpdated)
		if err != nil {
			return nil, err
		}
		result = append(result, &inv)
	}
	return result, nil
}

// Inventory - Get by ID
func (r *queryResolver) Inventory(ctx context.Context, id string) (*model.Inventory, error) {
	query := `SELECT id, item_name, quantity, unit, low_stock_threshold, expiry_date, last_updated FROM public.inventory WHERE id = $1`
	var inv model.Inventory
	err := r.Resolver.DB7.QueryRow(query, id).
		Scan(&inv.ID, &inv.ItemName, &inv.Quantity, &inv.Unit, &inv.LowStockThreshold, &inv.ExpiryDate, &inv.LastUpdated)
	if err != nil {
		return nil, err
	}
	return &inv, nil
}

// CreateOrderQueue
func (r *mutationResolver) CreateOrderQueue(ctx context.Context, orderID string, status *model.OrderStatus, priority *int32) (*model.OrderQueue, error) {
	query := `INSERT INTO public.order_queue (order_id, status, priority, last_updated)
	          VALUES ($1, COALESCE($2, 'preparing'), COALESCE($3, 1), NOW())
	          RETURNING order_id, status, priority, last_updated`

	var order model.OrderQueue
	err := r.Resolver.DB7.QueryRow(query, orderID, status, priority).
		Scan(&order.OrderID, &order.Status, &order.Priority, &order.LastUpdated)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

// UpdateOrderQueue
func (r *mutationResolver) UpdateOrderQueue(ctx context.Context, id string, status *model.OrderStatus, priority *int32) (*model.OrderQueue, error) {
	query := `UPDATE public.order_queue SET 
	          status = COALESCE($1, status),
	          priority = COALESCE($2, priority),
	          last_updated = NOW()
	          WHERE order_id = $3
	          RETURNING order_id, status, priority, last_updated`

	var order model.OrderQueue
	err := r.Resolver.DB7.QueryRow(query, status, priority, id).
		Scan(&order.OrderID, &order.Status, &order.Priority, &order.LastUpdated)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

// DeleteOrderQueue
func (r *mutationResolver) DeleteOrderQueue(ctx context.Context, id string) (bool, error) {
	query := `DELETE FROM public.order_queue WHERE order_id = $1`
	_, err := r.Resolver.DB7.Exec(query, id)
	if err != nil {
		return false, err
	}
	return true, nil
}

// OrderQueues - Get all
func (r *queryResolver) OrderQueues(ctx context.Context) ([]*model.OrderQueue, error) {
	query := `SELECT order_id, status, priority, last_updated FROM public.order_queue`
	rows, err := r.Resolver.DB7.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*model.OrderQueue
	for rows.Next() {
		var o model.OrderQueue
		err := rows.Scan(&o.OrderID, &o.Status, &o.Priority, &o.LastUpdated)
		if err != nil {
			return nil, err
		}
		result = append(result, &o)
	}
	return result, nil
}

// OrderQueue - Get by ID
func (r *queryResolver) OrderQueue(ctx context.Context, id string) (*model.OrderQueue, error) {
	query := `SELECT order_id, status, priority, last_updated FROM public.order_queue WHERE order_id = $1`
	var o model.OrderQueue
	err := r.Resolver.DB7.QueryRow(query, id).
		Scan(&o.OrderID, &o.Status, &o.Priority, &o.LastUpdated)
	if err != nil {
		return nil, err
	}
	return &o, nil
}
