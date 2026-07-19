package dto

type RegisterRequest struct {
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=8"`
	Name        string `json:"name" validate:"required"`
	PhoneNumber string `json:"phone_number"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh" validate:"required"`
}

type TokenResponse struct {
	Access  string      `json:"access"`
	Refresh string      `json:"refresh"`
	User    interface{} `json:"user,omitempty"`
}
