package controller

import (
	"first_task/helper"
	"first_task/middleware"
	"first_task/model/web"
	"first_task/service"
	// "fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func NewRoleController(roleService service.RoleService) RoleController {
	return &RoleControllerImpl{
		RoleService: roleService,
	}
}

type RoleControllerImpl struct {
	RoleService service.RoleService
}

func (controller *RoleControllerImpl) CreateRole(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	middleware.AuthAdmin(request)

	// roleCreateRequest := web.RoleCreateRequest{}
	// helper.ReadFromRequestBody(request, &roleCreateRequest)
	// fmt.Println(request)

	// controller.RoleService.CreateRole(request.Context(), roleCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RoleControllerImpl) UpdateRole(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	middleware.AuthAdmin(request)

	roleUpdateRequest := web.RoleUpdateRequest{}
	helper.ReadFromRequestBody(request, &roleUpdateRequest)

	roleId := params.ByName("roleId")
	id, err := strconv.Atoi(roleId)
	helper.PanicIfError(err)

	roleUpdateRequest.RoleId = id

	roleResponse := controller.RoleService.UpdateRole(request.Context(), roleUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   roleResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RoleControllerImpl) DeleteRole(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	middleware.AuthAdmin(request)

	categoryId := params.ByName("roleId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.RoleService.DeleteRole(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RoleControllerImpl) FindRoleById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	middleware.AuthAdmin(request)

	roleId := params.ByName("roleId")
	id, err := strconv.Atoi(roleId)
	helper.PanicIfError(err)

	roleResponse := controller.RoleService.FindByIdRole(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   roleResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RoleControllerImpl) FindAllRole(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	middleware.AuthAdmin(request)

	roleResponse := controller.RoleService.FindAllRole(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   roleResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
