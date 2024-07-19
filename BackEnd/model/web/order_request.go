package web

type OrderCreateRequest struct {
	CartId int `json:"cart_id"`
}

type OrderUpdateRequest struct {
	Id            int    `json:"id"`
	OrderStatus   string `json:"order_status"`
	PaymentStatus string `json:"payment_status"`
}
