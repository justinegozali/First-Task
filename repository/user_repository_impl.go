package repository

import (
	"context"
	"first_task/model/domain"
	"first_task/model/web"

	"gorm.io/gorm"
)

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

type UserRepositoryImpl struct {
}

// func (repository *UserRepositoryImpl) Login(ctx context.Context, tx *gorm.DB, data web.UserLoginRequest) (domain.User, error) {
// 	user := domain.User{}
// 	err := tx.Where("user_email = ?", data.UserEmail).Where("user_password = ?", data.UserPassword).Take(&user).Error
// 	if err != nil  {
// 		return user, err
// 	}
// 	return user, nil
// }

func (repository *UserRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, data web.UserRegisterRequest) (domain.User, error) {
	user := domain.User{
		UserName: data.UserName,
		UserEmail: data.UserEmail,
		UserPassword: data.UserPassword,
		UserRole: data.UserRole,
	}
	err := tx.Create(&user).Error
	if err != nil  {
		return user, err
	}
	return user, nil
}

func (repository *UserRepositoryImpl) FindUserByEmail(ctx context.Context, tx *gorm.DB, userEmail string) (domain.User, error) {
	user := domain.User{}
	err := tx.Where("user_email = ?", userEmail).Joins("Role").Take(&user).Error
	if err != nil  {
		return user, err
	}
	return user, nil
}
