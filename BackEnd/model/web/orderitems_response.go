package web

type OrderItemsResponse struct {
	Id        int             `json:"id"`
	OrderId   int             `json:"order_id"`
	ProductId int             `json:"product_id"`
	Product   ProductResponse `json:"product"`
	Quantity  int             `json:"quantity"`
	Price     int             `json:"price"`
}
