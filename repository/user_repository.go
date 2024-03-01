package repository

import (
	"context"
	"first_task/model/domain"
	"first_task/model/web"

	"gorm.io/gorm"
)

type UserRepository interface {
	// Login(ctx context.Context, tx *gorm.DB, data web.UserLoginRequest) (domain.User, error)
	Create(ctx context.Context, tx *gorm.DB, data web.UserRegisterRequest) (domain.User, error)
	FindUserByEmail(ctx context.Context, tx *gorm.DB, data string) (domain.User, error)
}
