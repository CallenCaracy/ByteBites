package model

import "time"

type Transaction struct {
	TransactionID     string    `json:"transactionID"`
	AmountPaid        float64   `json:"amountPaid"`
	PaymentMethod     string    `json:"paymentMethod"`
	TransactionStatus string    `json:"transactionStatus"`
	UserID            string    `json:"userID"`
	OrderID           string    `json:"orderID"`
	Timestamp         time.Time `json:"timestamp"`
}
