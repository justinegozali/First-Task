package main

import (
	"first_task/app"
	"first_task/controller"
	"first_task/exception"
	"first_task/helper"
	"first_task/repository"
	"first_task/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.OpenConnection()
	validate := validator.New()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)


	// roleRepository := repository.NewRoleRepository()
	// roleService := service.NewRoleService(roleRepository, db, validate)
	// roleController := controller.NewRoleController(roleService)

	router := httprouter.New()

	router.GET("/api/login", userController.Login)
	router.POST("/api/register", userController.Register)
	router.GET("/api/logout", userController.Logout)


	// router.GET("/api/roles", roleController.FindAllRole)
	// router.POST("/api/roles", roleController.CreateRole)
	// router.GET("/api/roles/:roleId", roleController.FindRoleById)
	// router.PUT("/api/roles/:roleId", roleController.UpdateRole)
	// router.DELETE("/api/roles/:roleId", roleController.DeleteRole)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	defer server.Close()
	helper.PanicIfError(err)
}