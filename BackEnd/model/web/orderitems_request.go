package web

type OrderItemsCreateRequest struct {
	CartId    []int `json:"cart_id"`
}