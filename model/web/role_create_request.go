package web

type RoleCreateRequest struct {
	RoleName string `json:"roleName" validate:"required" `
}
