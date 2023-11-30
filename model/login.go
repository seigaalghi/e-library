package model

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token     string `json:"token"`
	ExpiredAt int64  `json:"expired_at"`
}
