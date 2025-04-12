package graph

import (
    "context"
    "database/sql"
    "fmt"
    "time"
    "payment-service/graph/model"
)

type Resolver struct {
    DB *sql.DB
}

// NewResolver initializes the resolver with a shared DB connection
func NewResolver(db *sql.DB) *Resolver {
    return &Resolver{
        DB: db,
    }
}

// Query resolvers
func (r *Resolver) GetCart(ctx context.Context, userID string) (*model.Cart, error) {
    rows, err := r.DB.Query(`SELECT id, food_item, quantity, price FROM cart_items WHERE user_id = $1`, userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var items []*model.CartItem
    var totalAmount float64
    for rows.Next() {
        var item model.CartItem
        if err := rows.Scan(&item.ID, &item.FoodItem, &item.Quantity, &item.Price); err != nil {
            return nil, err
        }
        totalAmount += float64(item.Quantity) * item.Price
        items = append(items, &item)
    }

    return &model.Cart{
        UserID:      userID,
        Items:       items,
        TotalAmount: totalAmount,
    }, nil
}

func (r *Resolver) GetPayment(ctx context.Context, orderID string) (*model.Payment, error) {
    var p model.Payment
    var paymentInfo sql.NullString
    err := r.DB.QueryRow(`SELECT order_id, payment_method, payment_info, amount_paid, transaction_id, transaction_status, timestamp FROM payments WHERE order_id = $1`, orderID).
        Scan(&p.OrderID, &p.PaymentMethod, &paymentInfo, &p.AmountPaid, &p.TransactionID, &p.TransactionStatus, &p.Timestamp)
    if err != nil {
        return nil, err
    }
    if paymentInfo.Valid {
        p.PaymentInfo = &paymentInfo.String
    }
    return &p, nil
}

func (r *Resolver) GetReceipt(ctx context.Context, transactionID string) (*model.Receipt, error) {
    var rc model.Receipt
    err := r.DB.QueryRow(`SELECT transaction_id, user_id, order_id, amount_paid, payment_method, timestamp, transaction_status FROM receipts WHERE transaction_id = $1`, transactionID).
        Scan(&rc.TransactionID, &rc.UserID, &rc.OrderID, &rc.AmountPaid, &rc.PaymentMethod, &rc.Timestamp, &rc.TransactionStatus)
    if err != nil {
        return nil, err
    }
    return &rc, nil
}

// Mutation resolvers
func (r *Resolver) AddToCart(ctx context.Context, userID string, items []model.CartItemInput) (*model.Cart, error) {
    tx, err := r.DB.Begin()
    if err != nil {
        return nil, err
    }
    defer tx.Rollback()

    var totalAmount float64
    for _, item := range items {
        _, err := tx.Exec(`INSERT INTO cart_items (user_id, food_item, quantity, price) VALUES ($1, $2, $3, $4)`,
            userID, item.FoodItem, item.Quantity, item.Price)
        if err != nil {
            return nil, err
        }
        totalAmount += float64(item.Quantity) * item.Price
    }

    if err := tx.Commit(); err != nil {
        return nil, err
    }

    return r.GetCart(ctx, userID)
}

func (r *Resolver) Checkout(ctx context.Context, userID string) (*model.Cart, error) {
    return r.GetCart(ctx, userID)
}

func (r *Resolver) ProcessPayment(ctx context.Context, orderID string, paymentMethod string, paymentInfo *string) (*model.Payment, error) {
    if paymentMethod != "cash" && paymentMethod != "online" {
        return nil, fmt.Errorf("invalid payment method: %s", paymentMethod)
    }

    transactionID := fmt.Sprintf("tx-%d", time.Now().Unix())
    timestamp := time.Now().Format(time.RFC3339)

    _, err := r.DB.Exec(`
        INSERT INTO payments (order_id, payment_method, payment_info, amount_paid, transaction_id, transaction_status, timestamp)
        VALUES ($1, $2, $3, $4, $5, $6, $7)`,
        orderID, paymentMethod, paymentInfo, 0.0, transactionID, "completed", timestamp)
    if err != nil {
        return nil, err
    }

    return r.GetPayment(ctx, orderID)
}

func (r *Resolver) GenerateReceipt(ctx context.Context, amountPaid float64, paymentMethod string, timestamp string, transactionStatus string, transactionID string, userID string, orderID string) (*model.Receipt, error) {
    _, err := r.DB.Exec(`
        INSERT INTO receipts (transaction_id, user_id, order_id, amount_paid, payment_method, timestamp, transaction_status)
        VALUES ($1, $2, $3, $4, $5, $6, $7)`,
        transactionID, userID, orderID, amountPaid, paymentMethod, timestamp, transactionStatus)
    if err != nil {
        return nil, err
    }

    return &model.Receipt{
        TransactionID:    transactionID,
        UserID:           userID,
        OrderID:          orderID,
        AmountPaid:       amountPaid,
        PaymentMethod:    paymentMethod,
        Timestamp:        timestamp,
        TransactionStatus: transactionStatus,
    }, nil
}
