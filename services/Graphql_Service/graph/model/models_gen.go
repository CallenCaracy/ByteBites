// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"

	"github.com/google/uuid"
)

type AuthResponse struct {
	AccessToken  string  `json:"accessToken"`
	RefreshToken string  `json:"refreshToken"`
	Error        *string `json:"error,omitempty"`
}

type MenuItem struct {
	ID                 string  `json:"id"`
	Name               string  `json:"name"`
	Description        *string `json:"description,omitempty"`
	Price              float64 `json:"price"`
	Category           *string `json:"category,omitempty"`
	AvailabilityStatus bool    `json:"availability_status"`
	ImageURL           *string `json:"image_url,omitempty"`
	CreatedAt          string  `json:"created_at"`
	UpdatedAt          *string `json:"updated_at,omitempty"`
}

type Mutation struct {
}

type NewMenuItem struct {
	Name               string  `json:"name"`
	Description        *string `json:"description,omitempty"`
	Price              float64 `json:"price"`
	Category           *string `json:"category,omitempty"`
	AvailabilityStatus bool    `json:"availability_status"`
	ImageURL           *string `json:"image_url,omitempty"`
}

type NewReceiptInput struct {
	TransactionID uuid.UUID `json:"transactionID"`
	UserID        uuid.UUID `json:"userID"`
	Amount        float64   `json:"amount"`
	PaymentMethod string    `json:"paymentMethod"`
}

type NewTransactionInput struct {
	OrderID       uuid.UUID `json:"orderID"`
	UserID        uuid.UUID `json:"userID"`
	AmountPaid    float64   `json:"amountPaid"`
	PaymentMethod string    `json:"paymentMethod"`
}

type Query struct {
}

type Receipt struct {
	ID            uuid.UUID    `json:"id"`
	Transaction   *Transaction `json:"transaction"`
	UserID        uuid.UUID    `json:"userID"`
	Amount        float64      `json:"amount"`
	PaymentMethod string       `json:"paymentMethod"`
	Timestamp     time.Time    `json:"timestamp"`
}

type SignInEmployeeInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpInput struct {
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Role      string  `json:"role"`
	Address   *string `json:"address,omitempty"`
	Phone     *string `json:"phone,omitempty"`
}

type Transaction struct {
	ID                   uuid.UUID `json:"id"`
	OrderID              uuid.UUID `json:"orderID"`
	UserID               uuid.UUID `json:"userID"`
	AmountPaid           float64   `json:"amountPaid"`
	PaymentMethod        string    `json:"paymentMethod"`
	TransactionStatus    string    `json:"transactionStatus"`
	TransactionTimestamp time.Time `json:"transactionTimestamp"`
}

type UpdateMenuItem struct {
	Name               *string  `json:"name,omitempty"`
	Description        *string  `json:"description,omitempty"`
	Price              *float64 `json:"price,omitempty"`
	Category           *string  `json:"category,omitempty"`
	AvailabilityStatus *bool    `json:"availability_status,omitempty"`
	ImageURL           *string  `json:"image_url,omitempty"`
}

type UpdateUserInput struct {
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
	Role      *string `json:"role,omitempty"`
	Address   *string `json:"address,omitempty"`
	Phone     *string `json:"phone,omitempty"`
	IsActive  *bool   `json:"isActive,omitempty"`
}

type User struct {
	ID        string  `json:"id"`
	Email     string  `json:"email"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Role      string  `json:"role"`
	Address   *string `json:"address,omitempty"`
	Phone     *string `json:"phone,omitempty"`
	IsActive  string  `json:"isActive"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}
