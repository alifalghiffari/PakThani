package service

import (
	"context"
	"database/sql"
	"project-workshop/go-api-ecom/exception"
	"project-workshop/go-api-ecom/helper"

	// "project-workshop/go-api-ecom/model/domain"
	"project-workshop/go-api-ecom/model/web"
	"project-workshop/go-api-ecom/repository"

	"github.com/go-playground/validator/v10"
)

type OrderItemsServiceImpl struct {
	OrderItemsRepository repository.OrderItemsRepository
	UserRepository       repository.UserRepository
	ProductRepository    repository.ProductRepository
	CartRepository       repository.CartRepository
	OrderRepository      repository.OrderRepository
	DB                   *sql.DB
	Validate             *validator.Validate
}

func NewOrderItemsService(orderItemsRepository repository.OrderItemsRepository, userRepository repository.UserRepository, productRepository repository.ProductRepository, cartRepository repository.CartRepository, orderRepository repository.OrderRepository, DB *sql.DB, validate *validator.Validate) OrderItemsService {
	return &OrderItemsServiceImpl{
		OrderItemsRepository: orderItemsRepository,
		UserRepository:       userRepository,
		ProductRepository:    productRepository,
		CartRepository:       cartRepository,
		OrderRepository:      orderRepository,
		DB:                   DB,
		Validate:             validate,
	}
}

func (service *OrderItemsServiceImpl) FindById(ctx context.Context, Id int) web.OrderItemsResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	orderItems, err := service.OrderItemsRepository.FindById(ctx, tx, Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToOrderItemsResponse(orderItems)
}

func (service *OrderItemsServiceImpl) FindByOrderId(ctx context.Context, orderId int) []web.OrderItemsResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	orderItems, err := service.OrderItemsRepository.FindByOrderId(ctx, tx, orderId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToOrderItemsResponses(orderItems)
}

func (service *OrderItemsServiceImpl) FindByUserId(ctx context.Context, userId int) []web.OrderItemsResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	orderItems, err := service.OrderItemsRepository.FindByUserId(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToOrderItemsResponses(orderItems)
}
