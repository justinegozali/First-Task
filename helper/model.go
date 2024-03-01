package helper

import (
	"first_task/model/domain"
	"first_task/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		UserId:    user.UserId,
		UserName:  user.UserName,
		UserEmail: user.UserEmail,
		UserRole:  user.Role.RoleName,
	}
}

func ToRoleResponse(role domain.Role) web.RoleResponse {
	return web.RoleResponse{
		RoleId:   role.RoleId,
		RoleName: role.RoleName,
	}
}
