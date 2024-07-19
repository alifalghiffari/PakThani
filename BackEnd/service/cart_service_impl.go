package service

import (
	"context"
	"database/sql"
	"errors"
	"project-workshop/go-api-ecom/exception"
	"project-workshop/go-api-ecom/helper"
	"project-workshop/go-api-ecom/model/domain"
	"project-workshop/go-api-ecom/model/web"
	"project-workshop/go-api-ecom/repository"

	"github.com/go-playground/validator/v10"
)

type CartServiceImpl struct {
	CartRepository    repository.CartRepository
	UserRepository    repository.UserRepository
	ProductRepository repository.ProductRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewCartService(cartRepository repository.CartRepository, userRepository repository.UserRepository, productRepository repository.ProductRepository, DB *sql.DB, validate *validator.Validate) CartService {
	return &CartServiceImpl{
		CartRepository:    cartRepository,
		UserRepository:    userRepository,
		ProductRepository: productRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *CartServiceImpl) AddToCart(ctx context.Context, request web.CartCreateRequest, userId int) web.CartResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(err)
	}

	// Fetch product information based on ProductId from the request
	product, err := service.ProductRepository.FindById(ctx, tx, request.ProductId)
	if err != nil {
		panic(err)
	}

	if product.Quantity < request.Quantity {
		panic(errors.New("quantity is not enough"))
	}

	products := []domain.Product{product}

	cart := domain.Cart{
		UserId:    user.Id,
		ProductId: request.ProductId,
		Quantity:  request.Quantity,
		Product:   products,
		Price:     product.Price * request.Quantity,
	}
	cart = service.CartRepository.AddToCart(ctx, tx, cart)

	return helper.ToCartResponse(cart)
}

func (service *CartServiceImpl) UpdateCart(ctx context.Context, request web.CartUpdateRequest, userId int) web.CartResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(err)
	}

	product, err := service.ProductRepository.FindById(ctx, tx, request.ProductId)
	if err != nil {
		panic(err)
	}

	products := []domain.Product{product}

	cart := domain.Cart{
		Id:        request.Id,
		UserId:    user.Id,
		ProductId: request.ProductId,
		Quantity:  request.Quantity,
		Product:   products,
		Price:     product.Price * request.Quantity,
	}
	cart = service.CartRepository.UpdateCart(ctx, tx, cart)

	return helper.ToCartResponse(cart)
}

func (service *CartServiceImpl) RemoveCart(ctx context.Context, request web.CartDeleteRequest, userId int) web.CartResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(err)
	}

	cart := domain.Cart{
		Id:        request.Id,
		UserId:    user.Id,
		ProductId: request.ProductId,
		Quantity:  request.Quantity,
	}
	cart = service.CartRepository.RemoveCart(ctx, tx, cart)

	return helper.ToCartResponse(cart)
}

func (service *CartServiceImpl) DeleteCart(ctx context.Context, request web.CartDeleteRequest, userId int) web.CartResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(err)
	}

	cart := domain.Cart{
		Id:     request.Id,
		UserId: user.Id,
	}
	cart = service.CartRepository.DeleteCart(ctx, tx, cart)

	return helper.ToCartResponse(cart)
}

// Service method to find a single cart by ID
func (service *CartServiceImpl) FindById(ctx context.Context, cartId []int) web.CartResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Fetch a single cart by its ID
	cart, err := service.CartRepository.FindById(ctx, tx, cartId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	var cartResponses []web.CartResponse
	for _, carts := range cart {
		cartResponse := web.CartResponse{
			Id:        carts.Id,
			ProductId: carts.ProductId,
			Quantity:  carts.Quantity,
			Price:     carts.Price,
		}
		cartResponses = append(cartResponses, cartResponse)
	}
	return cartResponses[0]
}

func (service *CartServiceImpl) FindByUserId(ctx context.Context, userId int) []web.CartResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	carts, err := service.CartRepository.FindByUserId(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCartResponses(carts)
}

func (service *CartServiceImpl) FindAll(ctx context.Context) []web.CartResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	carts := service.CartRepository.FindAll(ctx, tx)
	return helper.ToCartResponses(carts)
}
