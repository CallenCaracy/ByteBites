package graph

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"payment-service/graph/model"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

// InitDB initializes the database connection with proper connection pooling
func InitDB() error {
	var err error
	// In production, use environment variables for connection string
	connectionString := os.Getenv("DATABASE_URL")
	if connectionString == "" {
		connectionString = "postgresql://postgres.jwltksifgtedjisubvxa:QLOVRo32Wv2eiHFp@aws-0-ap-southeast-1.pooler.supabase.com:5432/postgres"
	}

	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		return fmt.Errorf("error opening DB: %v", err)
	}

	// Configure connection pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Test connection with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return fmt.Errorf("error pinging DB: %v", err)
	}
	return nil
}

// checkDB ensures database is initialized before queries with timeout
func checkDB() error {
	if db == nil {
		if err := InitDB(); err != nil {
			return fmt.Errorf("database not initialized: %v", err)
		}
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	
	return db.PingContext(ctx) // Check connection is still alive
}

// AddToCart is the resolver for the addToCart field.
func (r *mutationResolver) AddToCart(ctx context.Context, userID string, items []*model.CartItemInput) (*model.Cart, error) {
	// Ensure database connection
	if err := checkDB(); err != nil {
		return nil, fmt.Errorf("database error: %v", err)
	}

	var total float64
	var cartItems []*model.CartItem

	// Calculate total before inserting the cart
	for _, item := range items {
		total += float64(item.Quantity) * item.Price
	}

	// Begin a transaction with context
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	// Check if a cart already exists for this user
	var existingCartID int
	err = tx.QueryRowContext(ctx, `SELECT id FROM carts WHERE user_id = $1`, userID).Scan(&existingCartID)
	if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("failed to check for existing cart: %v", err)
	}

	var cartID int
	if err == sql.ErrNoRows {
		// No existing cart, create a new one
		err = tx.QueryRowContext(ctx, `INSERT INTO carts (user_id, total_amount) VALUES ($1, $2) RETURNING id`, 
			userID, total).Scan(&cartID)
		if err != nil {
			return nil, fmt.Errorf("failed to insert cart: %v", err)
		}
	} else {
		// Cart exists, update total amount and clear existing items
		cartID = existingCartID
		_, err = tx.ExecContext(ctx, `DELETE FROM cart_items WHERE cart_id = $1`, cartID)
		if err != nil {
			return nil, fmt.Errorf("failed to clear existing cart items: %v", err)
		}
	}

	// Insert items into the database
	for _, item := range items {
		var itemID int
		err = tx.QueryRowContext(ctx, `INSERT INTO cart_items (cart_id, food_item, quantity, price) 
			VALUES ($1, $2, $3, $4) RETURNING id`, 
			cartID, item.FoodItem, item.Quantity, item.Price).Scan(&itemID)
		if err != nil {
			return nil, fmt.Errorf("failed to insert item: %v", err)
		}
		
		cartItem := &model.CartItem{
			ID:       fmt.Sprintf("item-%d", itemID),
			FoodItem: item.FoodItem,
			Quantity: item.Quantity,
			Price:    item.Price,
		}
		cartItems = append(cartItems, cartItem)
	}

	// Update total amount in the cart
	_, err = tx.ExecContext(ctx, "UPDATE carts SET total_amount = $1 WHERE id = $2", total, cartID)
	if err != nil {
		return nil, fmt.Errorf("failed to update total amount: %v", err)
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	// Create a response Cart model
	cart := &model.Cart{
		UserID:      userID,
		TotalAmount: total,
		Items:       cartItems,
	}

	return cart, nil
}

// Checkout is the resolver for the checkout field.
func (r *mutationResolver) Checkout(ctx context.Context, userID string) (*model.Cart, error) {
	// Ensure database connection
	if err := checkDB(); err != nil {
		return nil, fmt.Errorf("database error: %v", err)
	}

	// Begin a transaction with context
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	// Get cart information
	var cart model.Cart
	cart.UserID = userID
	var cartID int
	err = tx.QueryRowContext(ctx, `SELECT id, total_amount FROM carts WHERE user_id = $1`, userID).Scan(&cartID, &cart.TotalAmount)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no cart found for user %s", userID)
		}
		return nil, fmt.Errorf("failed to fetch cart: %v", err)
	}

	// Get cart items
	rows, err := tx.QueryContext(ctx, `SELECT id, food_item, quantity, price FROM cart_items WHERE cart_id = $1`, cartID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch cart items: %v", err)
	}
	defer rows.Close()

	// Parse cart items
	var items []*model.CartItem
	for rows.Next() {
		var item model.CartItem
		var itemID int
		if err := rows.Scan(&itemID, &item.FoodItem, &item.Quantity, &item.Price); err != nil {
			return nil, fmt.Errorf("failed to parse cart item: %v", err)
		}
		item.ID = fmt.Sprintf("item-%d", itemID)
		items = append(items, &item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during cart items iteration: %v", err)
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	cart.Items = items
	return &cart, nil
}

// ProcessPayment is the resolver for the processPayment field.
func (r *mutationResolver) ProcessPayment(ctx context.Context, orderID string, paymentMethod string, paymentInfo *string) (*model.Payment, error) {
	// Ensure database connection
	if err := checkDB(); err != nil {
		return nil, fmt.Errorf("database error: %v", err)
	}

	// Begin a transaction with context
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	// Query the cart for payment amount
	var amount float64
	err = tx.QueryRowContext(ctx, `SELECT total_amount FROM carts WHERE user_id = $1`, orderID).Scan(&amount)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no cart found for order %s", orderID)
		}
		return nil, fmt.Errorf("failed to fetch cart for payment: %v", err)
	}

	// Handle nil paymentInfo
	if paymentInfo == nil {
		paymentInfo = new(string)
		*paymentInfo = "" // Or handle the absence of payment info as needed
	}

	// Generate unique transaction ID
	transactionID := fmt.Sprintf("txn-%s-%d", orderID, time.Now().Unix())
	
	// Insert payment into the database
	insertPaymentQuery := `INSERT INTO payments (order_id, payment_method, payment_info, amount_paid, transaction_id, transaction_status, timestamp)
	VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err = tx.ExecContext(ctx, insertPaymentQuery, orderID, paymentMethod, paymentInfo, amount, transactionID, "SUCCESS", time.Now())
	if err != nil {
		return nil, fmt.Errorf("failed to insert payment: %v", err)
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	// Create Payment model
	timestamp := time.Now().Format(time.RFC3339)
	payment := &model.Payment{
		OrderID:           orderID,
		PaymentMethod:     paymentMethod,
		PaymentInfo:       paymentInfo,
		AmountPaid:        amount,
		TransactionID:     transactionID,
		TransactionStatus: "SUCCESS",
		Timestamp:         timestamp,
	}

	return payment, nil
}

// GenerateReceipt is the resolver for the generateReceipt field.
func (r *mutationResolver) GenerateReceipt(ctx context.Context, amountPaid float64, paymentMethod string, timestamp string, transactionStatus string, transactionID string, userID string, orderID string) (*model.Receipt, error) {
	// Ensure database connection
	if err := checkDB(); err != nil {
		return nil, fmt.Errorf("database error: %v", err)
	}

	// Begin a transaction with context
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	// Check if receipt already exists
	var existingReceiptID string
	err = tx.QueryRowContext(ctx, `SELECT transaction_id FROM receipts WHERE transaction_id = $1`, transactionID).Scan(&existingReceiptID)
	if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("failed to check for existing receipt: %v", err)
	}

	if err == sql.ErrNoRows {
		// Insert receipt into the database
		insertReceiptQuery := `INSERT INTO receipts (transaction_id, user_id, order_id, amount_paid, payment_method, timestamp, transaction_status)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
		_, err = tx.ExecContext(ctx, insertReceiptQuery, transactionID, userID, orderID, amountPaid, paymentMethod, timestamp, transactionStatus)
		if err != nil {
			return nil, fmt.Errorf("failed to generate receipt: %v", err)
		}
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	// Create Receipt model
	receipt := &model.Receipt{
		TransactionID:     transactionID,
		UserID:            userID,
		OrderID:           orderID,
		AmountPaid:        amountPaid,
		PaymentMethod:     paymentMethod,
		Timestamp:         timestamp,
		TransactionStatus: transactionStatus,
	}

	return receipt, nil
}

// GetCart is the resolver for the getCart field.
func (r *queryResolver) GetCart(ctx context.Context, userID string) (*model.Cart, error) {
	// Ensure database connection
	if err := checkDB(); err != nil {
		return nil, fmt.Errorf("database error: %v", err)
	}

	// Begin a transaction with context
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	// Get cart information
	var cart model.Cart
	cart.UserID = userID
	var cartID int
	err = tx.QueryRowContext(ctx, `SELECT id, total_amount FROM carts WHERE user_id = $1`, userID).Scan(&cartID, &cart.TotalAmount)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no cart found for user %s", userID)
		}
		return nil, fmt.Errorf("failed to fetch cart: %v", err)
	}

	// Get cart items
	rows, err := tx.QueryContext(ctx, `SELECT id, food_item, quantity, price FROM cart_items WHERE cart_id = $1`, cartID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch cart items: %v", err)
	}
	defer rows.Close()

	// Parse cart items
	var items []*model.CartItem
	for rows.Next() {
		var item model.CartItem
		var itemID int
		if err := rows.Scan(&itemID, &item.FoodItem, &item.Quantity, &item.Price); err != nil {
			return nil, fmt.Errorf("failed to parse cart item: %v", err)
		}
		item.ID = fmt.Sprintf("item-%d", itemID)
		items = append(items, &item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during cart items iteration: %v", err)
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	cart.Items = items
	return &cart, nil
}

// GetPayment is the resolver for the getPayment field.
func (r *queryResolver) GetPayment(ctx context.Context, orderID string) (*model.Payment, error) {
	// Ensure database connection
	if err := checkDB(); err != nil {
		return nil, fmt.Errorf("database error: %v", err)
	}

	var payment model.Payment
	err := db.QueryRowContext(ctx, `SELECT order_id, payment_method, payment_info, amount_paid, transaction_id, transaction_status, timestamp 
		FROM payments WHERE order_id = $1 ORDER BY timestamp DESC LIMIT 1`, orderID).Scan(
		&payment.OrderID, &payment.PaymentMethod, &payment.PaymentInfo, &payment.AmountPaid, 
		&payment.TransactionID, &payment.TransactionStatus, &payment.Timestamp)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("payment not found for order %s", orderID)
		}
		return nil, fmt.Errorf("failed to fetch payment: %v", err)
	}
	return &payment, nil
}

// GetReceipt is the resolver for the getReceipt field.
func (r *queryResolver) GetReceipt(ctx context.Context, transactionID string) (*model.Receipt, error) {
	// Ensure database connection
	if err := checkDB(); err != nil {
		return nil, fmt.Errorf("database error: %v", err)
	}

	var receipt model.Receipt
	err := db.QueryRowContext(ctx, `SELECT transaction_id, user_id, order_id, amount_paid, payment_method, timestamp, transaction_status 
		FROM receipts WHERE transaction_id = $1`, transactionID).Scan(
		&receipt.TransactionID, &receipt.UserID, &receipt.OrderID, &receipt.AmountPaid, 
		&receipt.PaymentMethod, &receipt.Timestamp, &receipt.TransactionStatus)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("receipt not found for transaction %s", transactionID)
		}
		return nil, fmt.Errorf("failed to fetch receipt: %v", err)
	}
	return &receipt, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }