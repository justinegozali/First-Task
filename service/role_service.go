package service

import (

	"context"
	"first_task/model/web"
)

type RoleService interface {
	CreateRole(ctx context.Context, request web.RoleCreateRequest) 
	UpdateRole(ctx context.Context, request web.RoleUpdateRequest) web.RoleResponse
	DeleteRole(ctx context.Context, categoryId int)
	FindByIdRole(ctx context.Context, categoryId int) web.RoleResponse
	FindAllRole(ctx context.Context) []web.RoleResponse
}
