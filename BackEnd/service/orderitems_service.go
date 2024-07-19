package service

import (
	"context"
	"project-workshop/go-api-ecom/model/web"
)

type OrderItemsService interface {
	FindById(ctx context.Context, Id int) web.OrderItemsResponse
	FindByOrderId(ctx context.Context, orderId int) []web.OrderItemsResponse
	FindByUserId(ctx context.Context, userId int) []web.OrderItemsResponse
}
