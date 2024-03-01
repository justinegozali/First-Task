package web

type UserResponse struct {
	UserId    int    `json:"userId"`
	UserName  string `json:"userName"`
	UserEmail string `json:"userEmail"`
	UserRole  string `json:"userRole"`
}
