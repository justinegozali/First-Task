package repository

import (
	"context"
	"first_task/model/domain"
	"first_task/model/web"

	"gorm.io/gorm"
)

func NewRoleRepository() RoleRepository {
	return &RoleRepositoryImpl{}
}

type RoleRepositoryImpl struct {
}

func (repository *RoleRepositoryImpl) CreateRole(ctx context.Context, tx *gorm.DB, request web.RoleCreateRequest) error{
	role := domain.Role{
		RoleName: request.RoleName,
	}
	err := tx.Create(&role)
	if err != nil  {
		return err.Error
	}
	return nil
}

func (repository *RoleRepositoryImpl) UpdateRole(ctx context.Context, tx *gorm.DB, request web.RoleUpdateRequest) (domain.Role, error){
	role := domain.Role{}
	err := tx.Where("role_id = ?", request.RoleId).Take(&role)
	if err != nil  {
		return role, err.Error
	}

	role.RoleId = request.RoleId
	role.RoleName = request.RoleName

	err = tx.Save(&role)
	if err != nil  {
		return role, err.Error
	}
	return role, nil
}

func (repository *RoleRepositoryImpl) DeleteRole(ctx context.Context, tx *gorm.DB, roleId int) error {
	err := tx.Where("role_id = ?", roleId).Delete(&domain.Role{})
	if err != nil {
		return err.Error
	}
	return nil
}

func (repository *RoleRepositoryImpl) FindRoleById(ctx context.Context, tx *gorm.DB, roleId int) (domain.Role, error){
	role := domain.Role{}
	err := tx.Where("role_id = ?", roleId).Take(&role)
	if err != nil  {
		return role, err.Error
	}
	return role, nil		
}

func (repository *RoleRepositoryImpl) FindAllRole(ctx context.Context, tx *gorm.DB) ([]domain.Role, error){
	var roles []domain.Role
	err := tx.Find(&roles)
	if err != nil  {
		return roles, err.Error
	}
	return roles, nil	
}