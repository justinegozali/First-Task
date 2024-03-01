package repository

import (
	"context"
	"first_task/model/domain"
	"first_task/model/web"

	"gorm.io/gorm"
)

type RoleRepository interface {
	CreateRole(ctx context.Context, tx *gorm.DB, request web.RoleCreateRequest) error
	UpdateRole(ctx context.Context, tx *gorm.DB, request web.RoleUpdateRequest) (domain.Role, error)
	DeleteRole(ctx context.Context, tx *gorm.DB, roleId int) error
	FindRoleById(ctx context.Context, tx *gorm.DB, roleId int) (domain.Role, error)
	FindAllRole(ctx context.Context, tx *gorm.DB) ([]domain.Role, error)
}
