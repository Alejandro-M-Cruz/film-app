package auth

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=2,max=32,alphanum,starts_with_alpha"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required,max=255"`
	Password string `json:"password" validate:"required,max=255"`
}
