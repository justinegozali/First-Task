package web

type RoleUpdateRequest struct {
	RoleId   int    `validate:"required" json:"roleId"`
	RoleName string `validate:"required" json:"roleName"`
}
