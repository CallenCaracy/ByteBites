package graph

import (
	"Graphql_Service/graph/model"
	"context"
)

// CreateTransactionRecords - Insert a new transaction record
func (r *mutationResolver) CreateTransactionRecords(ctx context.Context, amountPaid float64, paymentMethod string, transactionStatus string, userID string, orderID string) (*model.Transaction, error) {
	query := `INSERT INTO public.transactions (amount_paid, payment_method, transaction_status, user_id, order_id, timestamp) 
			  VALUES ($1, $2, $3, $4, $5, NOW()) RETURNING transaction_id, amount_paid, payment_method, transaction_status, user_id, order_id, timestamp`

	item := &model.Transaction{
		AmountPaid:        amountPaid,
		PaymentMethod:     paymentMethod,
		TransactionStatus: transactionStatus,
		UserID:            userID,
		OrderID:           orderID,
	}

	err := r.Resolver.DB3.QueryRow(query,
		item.AmountPaid, item.PaymentMethod, item.TransactionStatus,
		item.UserID, item.OrderID,
	).Scan(
		&item.TransactionID, &item.AmountPaid, &item.PaymentMethod,
		&item.TransactionStatus, &item.UserID, &item.OrderID, &item.Timestamp,
	)

	if err != nil {
		return nil, err
	}

	return item, nil
}

// GetTransactionRecordsByUserID - Fetch transaction records by user ID
func (r *queryResolver) GetTransactionRecordsByUserID(ctx context.Context, id string) ([]*model.Transaction, error) {
	query := `SELECT transaction_id, amount_paid, payment_method, transaction_status, user_id, order_id, timestamp
			  FROM public.transactions WHERE user_id = $1`

	rows, err := r.Resolver.DB3.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []*model.Transaction
	for rows.Next() {
		var item model.Transaction
		err := rows.Scan(
			&item.TransactionID, &item.AmountPaid, &item.PaymentMethod,
			&item.TransactionStatus, &item.UserID, &item.OrderID, &item.Timestamp,
		)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, &item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}

// GetTransactionByID - Fetch a single transaction by ID
func (r *queryResolver) GetTransactionByID(ctx context.Context, id string) (*model.Transaction, error) {
	query := `SELECT transaction_id, amount_paid, payment_method, transaction_status, user_id, order_id, timestamp
			  FROM public.transactions WHERE transaction_id = $1`

	var item model.Transaction
	err := r.Resolver.DB3.QueryRow(query, id).Scan(
		&item.TransactionID, &item.AmountPaid, &item.PaymentMethod,
		&item.TransactionStatus, &item.UserID, &item.OrderID, &item.Timestamp,
	)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

// UpdateTransactionStatus - Update transaction status by ID
func (r *mutationResolver) UpdateTransactionStatus(ctx context.Context, transactionID string, newStatus string) (*model.Transaction, error) {
	query := `UPDATE public.transactions SET transaction_status = $1 WHERE transaction_id = $2 
			  RETURNING transaction_id, amount_paid, payment_method, transaction_status, user_id, order_id, timestamp`

	var item model.Transaction
	err := r.Resolver.DB3.QueryRow(query, newStatus, transactionID).Scan(
		&item.TransactionID, &item.AmountPaid, &item.PaymentMethod,
		&item.TransactionStatus, &item.UserID, &item.OrderID, &item.Timestamp,
	)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

// DeleteTransactionByID - Delete a transaction by ID
func (r *mutationResolver) DeleteTransactionByID(ctx context.Context, transactionID string) (bool, error) {
	query := `DELETE FROM public.transactions WHERE transaction_id = $1`

	_, err := r.Resolver.DB3.Exec(query, transactionID)
	if err != nil {
		return false, err
	}

	return true, nil
}
