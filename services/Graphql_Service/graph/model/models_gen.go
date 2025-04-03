// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AuthResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
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

type Query struct {
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
