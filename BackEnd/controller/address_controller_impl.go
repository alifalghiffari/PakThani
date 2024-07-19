package controller

import (
	"net/http"
	"project-workshop/go-api-ecom/helper"
	"project-workshop/go-api-ecom/model/web"
	"project-workshop/go-api-ecom/service"

	"github.com/julienschmidt/httprouter"
)

type AddressControllerImpl struct {
	AddressService service.AddressService
}

func NewAddressController(addressService service.AddressService) AddressController {
	return &AddressControllerImpl{
		AddressService: addressService,
	}
}

func (controller *AddressControllerImpl) AddAddress(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	addressCreateRequest := web.AddressCreateRequest{}
	userId := request.Context().Value("userId").(int)
	helper.ReadFromRequestBody(request, &addressCreateRequest)

	addressResponse := controller.AddressService.AddAddress(request.Context(), addressCreateRequest, userId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   addressResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AddressControllerImpl) UpdateAddress(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	addressUpdateRequest := web.AddressUpdateRequest{}
	userId := request.Context().Value("userId").(int)
	helper.ReadFromRequestBody(request, &addressUpdateRequest)

	addressResponse := controller.AddressService.UpdateAddress(request.Context(), addressUpdateRequest, userId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   addressResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AddressControllerImpl) FindByUserId(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := request.Context().Value("userId").(int)

	addressResponse := controller.AddressService.FindByUserId(request.Context(), userId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   addressResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
