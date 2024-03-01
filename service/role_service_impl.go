package service

import (
	"context"
	"first_task/exception"
	"first_task/helper"
	"first_task/model/web"
	"first_task/repository"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func NewRoleService(roleRepository repository.RoleRepository, db *gorm.DB, validate *validator.Validate) RoleService {
	return &RoleServiceImpl{
		RoleRepository: roleRepository,
		DB:             db,
		Validate:       validate,
	}
}

type RoleServiceImpl struct {
	RoleRepository repository.RoleRepository
	DB             *gorm.DB
	Validate       *validator.Validate
}

func (service RoleServiceImpl) CreateRole(ctx context.Context, request web.RoleCreateRequest) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer tx.Rollback()

	err = service.RoleRepository.CreateRole(ctx, tx, request)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	tx.Commit()
}

func (service RoleServiceImpl) UpdateRole(ctx context.Context, request web.RoleUpdateRequest) web.RoleResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer tx.Rollback()

	role, err := service.RoleRepository.UpdateRole(ctx, tx, request)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	tx.Commit()

	return helper.ToRoleResponse(role)
}

func (service RoleServiceImpl) DeleteRole(ctx context.Context, categoryId int) {
	tx := service.DB.Begin()
	defer tx.Rollback()

	err := service.RoleRepository.DeleteRole(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	tx.Commit()
}

func (service RoleServiceImpl) FindByIdRole(ctx context.Context, categoryId int) web.RoleResponse {
	tx := service.DB.Begin()
	defer tx.Rollback()

	role, err := service.RoleRepository.FindRoleById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	tx.Commit()

	return helper.ToRoleResponse(role)
}

func (service RoleServiceImpl) FindAllRole(ctx context.Context) []web.RoleResponse {
	tx := service.DB.Begin()
	defer tx.Rollback()

	roles, err := service.RoleRepository.FindAllRole(ctx, tx)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	tx.Commit()

	webRoleResponses := []web.RoleResponse{}
	for _, role := range roles {
		roleResponse := web.RoleResponse{
			RoleId:   role.RoleId,
			RoleName: role.RoleName,
		}
		webRoleResponses = append(webRoleResponses, roleResponse)
	}
	return webRoleResponses
}
