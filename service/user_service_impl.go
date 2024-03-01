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

func NewUserService(userRepository repository.UserRepository, db *gorm.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
		Validate:       validate,
	}
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *gorm.DB
	Validate       *validator.Validate
}

func (service *UserServiceImpl) Login(ctx context.Context, request web.UserLoginRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer tx.Rollback()

	// check if email exist 
	user, err := service.UserRepository.FindUserByEmail(ctx, tx, request.UserEmail)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// check password match
	check := helper.CheckPasswordHash(request.UserPassword, user.UserPassword)
	if !check {
		panic(exception.NewNotFoundError("record not found"))
	}

	tx.Commit()

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Register(ctx context.Context, request web.UserRegisterRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer tx.Rollback()

	// check if email already exist
	_, err = service.UserRepository.FindUserByEmail(ctx, tx, request.UserEmail)
	if err == nil {
		panic(exception.NewNotUniqueError())
	}

	// hashing password & input user
	hash, err := helper.HashPassword(request.UserPassword)
	helper.PanicIfError(err)
	request.UserPassword = hash
	
	_, err = service.UserRepository.Create(ctx, tx, request)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// return user info
	user, err := service.UserRepository.FindUserByEmail(ctx, tx, request.UserEmail)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	tx.Commit()

	return helper.ToUserResponse(user)
}
