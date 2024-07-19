package service

import (
	"context"
	"project-workshop/go-api-ecom/model/web"
)

type AddressService interface {
	AddAddress(ctx context.Context, request web.AddressCreateRequest, userId int) web.AddressResponse
	UpdateAddress(ctx context.Context, request web.AddressUpdateRequest, userId int) web.AddressResponse
	FindByUserId(ctx context.Context, userId int) []web.AddressResponse
}
