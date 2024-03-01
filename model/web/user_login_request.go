package web

type UserLoginRequest struct {
	UserEmail    string `json:"userEmail" validate:"required,email"`
	UserPassword string `json:"userPassword" validate:"required"`
}
