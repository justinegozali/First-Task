package controller

import (
	"first_task/helper"
	"first_task/model/web"
	"first_task/service"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

type UserControllerImpl struct {
	UserService service.UserService
}

func (controller *UserControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userLoginRequest := web.UserLoginRequest{}
	helper.ReadFromRequestBody(request, &userLoginRequest)

	userResponse := controller.UserService.Login(request.Context(), userLoginRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	tokenString, err := helper.CreateToken(userResponse.UserName, userResponse.UserRole)
	helper.PanicIfError(err)

	http.SetCookie(writer, &http.Cookie{
		Name:  "token",
		Value: tokenString,
	})

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userRegisterRequest := web.UserRegisterRequest{}
	helper.ReadFromRequestBody(request, &userRegisterRequest)

	userResponse := controller.UserService.Register(request.Context(), userRegisterRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Logout(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	http.SetCookie(writer, &http.Cookie{
		Name:  "token",
		Expires: time.Now(),
	})

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
