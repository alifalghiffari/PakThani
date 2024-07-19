package controller

import (
	"net/http"
	"project-workshop/go-api-ecom/helper"
	"project-workshop/go-api-ecom/model/web"
	"project-workshop/go-api-ecom/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type OrderItemsControllerImpl struct {
	OrderItemsService service.OrderItemsService
}

func NewOrderItemsController(orderItemsService service.OrderItemsService) OrderItemsController {
	return &OrderItemsControllerImpl{
		OrderItemsService: orderItemsService,
	}
}

func (controller *OrderItemsControllerImpl) FindOrderItemsById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	Id, err := strconv.Atoi(params.ByName("orderItemId"))
	helper.PanicIfError(err)

	orderItemsResponse := controller.OrderItemsService.FindById(request.Context(), Id)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   orderItemsResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrderItemsControllerImpl) FindOrderItemsByOrderId(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	orderId, err := strconv.Atoi(params.ByName("orderId"))
	helper.PanicIfError(err)

	orderItemsResponse := controller.OrderItemsService.FindByOrderId(request.Context(), orderId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   orderItemsResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrderItemsControllerImpl) FindOrderItemsByUserId(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := request.Context().Value("userId").(int)

	orderItemsResponse := controller.OrderItemsService.FindByUserId(request.Context(), userId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   orderItemsResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
