// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Query struct {
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
