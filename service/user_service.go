package service

import (
	"context"
	"first_task/model/web"
)

type UserService interface {
	Login(ctx context.Context, request web.UserLoginRequest) web.UserResponse
	Register(ctx context.Context, request web.UserRegisterRequest) web.UserResponse
}