package web

type UserRegisterRequest struct {
	UserName     string `json:"userName" validate:"required"`
	UserEmail    string `json:"userEmail" validate:"required,email"`
	UserPassword string `json:"userPassword" validate:"required"`
	UserRole     int    `json:"userRole" validate:"required,min=1,max=2"`
}