package web

type OrderResponse struct {
	Id            int                  `json:"id"`
	UserId        int                  `json:"user_id"`
	TotalItems    int                  `json:"total_items"`
	TotalPrice    int                  `json:"total_price"`
	OrderStatus   string               `json:"order_status"`
	PaymentStatus string               `json:"payment_status"`
	OrderItems    []OrderItemsResponse `json:"order_items"`
	User          UserResponse         `json:"user"`
	Address       AddressResponse      `json:"address"`
}
