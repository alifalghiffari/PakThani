package helper

import (
	"project-workshop/go-api-ecom/model/domain"
	"project-workshop/go-api-ecom/model/web"
	// "strconv"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:       category.Id,
		Category: category.Category,
		Icon:     category.Icon,
		Products: ToProductResponses(category.Products),
		Slug:     category.Slug,
	}
}

func ToUserResponse(user domain.User) web.UserResponse {
	role := "user"
	if !user.Role {
		role = "admin"
	}

	return web.UserResponse{
		Id:        user.Id,
		Username:  user.Username,
		Password:  user.Password,
		Email:     user.Email,
		NoTelepon: user.NoTelepon,
		Role:      role,
	}
}

func ToAccountResponse(user domain.User) web.AccountResponse {
	role := "user"
	if !user.Role {
		role = "admin"
	}

	return web.AccountResponse{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		NoTelepon: user.NoTelepon,
		Role:      role,
	}
}

func ToProductResponse(product domain.Product) web.ProductResponse {
	return web.ProductResponse{
		Id:          product.Id,
		Name:        product.Name,
		Image:       product.Image,
		Description: product.Description,
		Price:       product.Price,
		CategoryId:  product.CategoryId,
		Category:    product.Category.Category,
		Quantity:    product.Quantity,
		Slug:        product.Slug,
		Images:      ToImageReponses(product.Images),
	}
}

func ToImageResponse(image domain.Image) web.ImageResponse {
	return web.ImageResponse{
		Id:        image.Id,
		ProductId: image.ProductId,
		Image:     image.Image,
	}
}

func ToOrderResponse(order domain.Order) web.OrderResponse {
	return web.OrderResponse{
		Id:            order.Id,
		UserId:        order.UserId,
		TotalItems:    order.TotalItems,
		TotalPrice:    order.TotalPrice,
		OrderStatus:   string(order.OrderStatus),
		PaymentStatus: string(order.PaymentStatus),
		OrderItems:    ToOrderItemsResponses(order.OrderItems),
		User:          ToUserResponse(order.User),
		Address:       ToAddressResponse(order.Address),
	}
}

func ToOrderItemsResponse(orderItems domain.OrderItems) web.OrderItemsResponse {
	return web.OrderItemsResponse{
		Id:        orderItems.Id,
		OrderId:   orderItems.OrderId,
		ProductId: orderItems.ProductId,
		Product:   ToProductResponse(orderItems.Product),
		Quantity:  orderItems.Quantity,
		Price:     orderItems.Price,
	}
}

func ToCartResponse(cart domain.Cart) web.CartResponse {
	return web.CartResponse{
		Id:        cart.Id,
		UserId:    cart.UserId,
		ProductId: cart.ProductId,
		Product:   ToProductResponses(cart.Product),
		Quantity:  cart.Quantity,
		Price:     cart.Price,
	}
}

func ToAddressResponse(address domain.Address) web.AddressResponse {
	return web.AddressResponse{
		Id:           address.Id,
		UserId:       address.UserId,
		NamaPenerima: address.NamaPenerima,
		Kabupaten:    address.Kabupaten,
		Kecamatan:    address.Kecamatan,
		Kelurahan:    address.Kelurahan,
		Alamat:       address.Alamat,
		NoHP:         address.NoHP,
		Note:         address.Note,
		User:         ToUserResponses(address.User),
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}

func ToUserResponses(users []domain.User) []web.UserResponse {
	var userResponses []web.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}
	return userResponses
}

func ToAccountResponses(users []domain.User) []web.AccountResponse {
	var accountResponses []web.AccountResponse
	for _, user := range users {
		accountResponses = append(accountResponses, ToAccountResponse(user))
	}
	return accountResponses
}

func ToProductResponses(products []domain.Product) []web.ProductResponse {
	var productResponses []web.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, ToProductResponse(product))
	}
	return productResponses
}

func ToOrderResponses(orders []domain.Order) []web.OrderResponse {
	var orderResponses []web.OrderResponse
	for _, order := range orders {
		orderResponses = append(orderResponses, ToOrderResponse(order))
	}
	return orderResponses
}

func ToOrderItemsResponses(orderItems []domain.OrderItems) []web.OrderItemsResponse {
	var orderItemsResponse []web.OrderItemsResponse
	for _, orderItem := range orderItems {
		orderItemsResponse = append(orderItemsResponse, ToOrderItemsResponse(orderItem))
	}
	return orderItemsResponse
}

func ToCartResponses(carts []domain.Cart) []web.CartResponse {
	var cartResponses []web.CartResponse
	for _, cart := range carts {
		cartResponses = append(cartResponses, ToCartResponse(cart))
	}
	return cartResponses
}

func ToAddressResponses(addresses []domain.Address) []web.AddressResponse {
	var addressResponses []web.AddressResponse
	for _, address := range addresses {
		addressResponses = append(addressResponses, ToAddressResponse(address))
	}

	return addressResponses
}

func ToImageReponses(images []domain.Image) []web.ImageResponse {
	var imagesResponse []web.ImageResponse
	for _, image := range images {
		imagesResponse = append(imagesResponse, ToImageResponse(image))
	}

	return imagesResponse
}
