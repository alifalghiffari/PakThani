package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AddressController interface {
	AddAddress(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateAddress(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByUserId(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
