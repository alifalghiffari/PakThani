package service

import (
	"context"
	"database/sql"
	"log"
	"project-workshop/go-api-ecom/exception"
	"project-workshop/go-api-ecom/helper"
	"project-workshop/go-api-ecom/model/domain"
	"project-workshop/go-api-ecom/model/web"
	"project-workshop/go-api-ecom/repository"

	"github.com/go-playground/validator/v10"
)

type OrderServiceImpl struct {
	OrderRepository      repository.OrderRepository
	UserRepository       repository.UserRepository
	CartRepository       repository.CartRepository
	OrderItemsRepository repository.OrderItemsRepository
	DB                   *sql.DB
	Validate             *validator.Validate
}

func NewOrderService(orderRepository repository.OrderRepository, userRepository repository.UserRepository, cartRepository repository.CartRepository, orderItemsRepository repository.OrderItemsRepository, DB *sql.DB, validate *validator.Validate) OrderService {
	return &OrderServiceImpl{
		OrderRepository:      orderRepository,
		UserRepository:       userRepository,
		CartRepository:       cartRepository,
		OrderItemsRepository: orderItemsRepository,
		DB:                   DB,
		Validate:             validate,
	}
}

func mapToOrderStatus(orderStatus string) domain.OrderStatus {
	switch orderStatus {
	case "PENDING":
		return domain.Pending
	case "PROCESS":
		return domain.Processing
	case "SHIPPING":
		return domain.Shipped
	case "DELIVERED":
		return domain.Delivered
	case "CANCELLED":
		return domain.Cancelled
	default:
		return domain.Pending
	}
}

func mapToPaymentStatus(paymentStatus string) domain.PaymentStatus {
	switch paymentStatus {
	case "PENDING":
		return domain.PaymentPending
	case "SUCCESS":
		return domain.PaymentSuccess
	case "FAILED":
		return domain.PaymentFailed
	default:
		return domain.PaymentPending
	}
}

func (service *OrderServiceImpl) FindAll(ctx context.Context) []web.OrderResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	orders := service.OrderRepository.FindAll(ctx, tx)
	return helper.ToOrderResponses(orders)
}

// create order mengambil data total items, total price yang berasal dari order items yang dimana order items tersebut
// buatkan order items setelah berhasil membuat order
func (service *OrderServiceImpl) CreateOrder(ctx context.Context, request web.OrderCreateRequest, userId int) web.OrderResponse {
	err := service.Validate.Struct(request)
	if err != nil {
		log.Printf("Validation error: %v", err)
		helper.PanicIfError(err)
	}

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		log.Printf("User not found error: %v", err)
		panic(exception.NewNotFoundError(err.Error()))
	}

	cart, err := service.CartRepository.FindByUserId(ctx, tx, userId)
	if err != nil {
		log.Printf("Cart not found error: %v", err)
		panic(exception.NewNotFoundError(err.Error()))
	}

	if len(cart) == 0 {
		log.Println("Cart is empty")
		panic(exception.NewNotFoundError("Cart is empty"))
	}

	totalItems := 0
	totalPrice := 0
	for _, cartItem := range cart {
		if len(cartItem.Product) > 0 {
			totalItems += cartItem.Quantity
			totalPrice += cartItem.Quantity * cartItem.Product[0].Price
		}
	}

	orderStatus := domain.Pending
	paymentStatus := domain.PaymentPending

	order := domain.Order{
		UserId:        user.Id,
		TotalItems:    totalItems,
		TotalPrice:    totalPrice,
		OrderStatus:   orderStatus,
		PaymentStatus: paymentStatus,
	}

	// Insert order to get order ID
	order = service.OrderRepository.Insert(ctx, tx, order)
	if err != nil {
		log.Printf("Order insert error: %v", err)
		helper.PanicIfError(err)
	}

	var orderItems []domain.OrderItems
	for _, cartItem := range cart {
		orderItem := domain.OrderItems{
			OrderId:   order.Id, // Ensure the correct OrderId is used
			ProductId: cartItem.ProductId,
			Quantity:  cartItem.Quantity,
			Price:     cartItem.Price,
		}
		orderItems = append(orderItems, orderItem)
	}

	for _, item := range orderItems {
		service.OrderItemsRepository.Insert(ctx, tx, item)
		if err != nil {
			log.Printf("OrderItems insert error: %v", err)
			helper.PanicIfError(err)
		}
	}

	// Clear the cart after order is created
	for _, cartItem := range cart {
		service.CartRepository.DeleteCart(ctx, tx, cartItem)
		if err != nil {
			log.Printf("Cart delete error: %v", err)
			helper.PanicIfError(err)
		}
	}

	return helper.ToOrderResponse(order)
}

func (service *OrderServiceImpl) UpdateOrder(ctx context.Context, request web.OrderUpdateRequest) web.OrderResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	order := domain.Order{
		Id:            request.Id,
		OrderStatus:   mapToOrderStatus(request.OrderStatus),
		PaymentStatus: mapToPaymentStatus(request.PaymentStatus),
	}

	order = service.OrderRepository.Update(ctx, tx, order)

	return helper.ToOrderResponse(order)
}

func (service *OrderServiceImpl) FindOrderByUserId(ctx context.Context, userId int) []web.OrderResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	orders, err := service.OrderRepository.FindByUserId(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	var orderResponses []web.OrderResponse

	for _, order := range orders {
		orderItems, err := service.OrderItemsRepository.FindByOrderId(ctx, tx, order.Id)
		if err != nil {
			panic(exception.NewNotFoundError(err.Error()))
		}

		var orderItemsResponses []web.OrderItemsResponse
		for _, orderItem := range orderItems {
			orderItemsResponse := web.OrderItemsResponse{
				Id:        orderItem.Id,
				OrderId:   orderItem.OrderId,
				ProductId: orderItem.ProductId,
				Quantity:  orderItem.Quantity,
				Price:     orderItem.Price,
			}
			orderItemsResponses = append(orderItemsResponses, orderItemsResponse)
		}

		orderResponse := web.OrderResponse{
			Id:            order.Id,
			UserId:        order.UserId,
			TotalItems:    order.TotalItems,
			TotalPrice:    order.TotalPrice,
			OrderStatus:   string(order.OrderStatus),
			PaymentStatus: string(order.PaymentStatus),
			OrderItems:    orderItemsResponses,
		}

		orderResponses = append(orderResponses, orderResponse)
	}

	return orderResponses
}

func (service *OrderServiceImpl) FindOrderById(ctx context.Context, Id int, userId int) web.OrderResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	order, err := service.OrderRepository.FindById(ctx, tx, Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	if order.UserId != userId {
		helper.PanicIfError(err)
	}

	return helper.ToOrderResponse(order)
}
