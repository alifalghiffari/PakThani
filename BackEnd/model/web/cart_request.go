package web

type CartCreateRequest struct {
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
	Price     int
}

type CartIdRequest struct {
	Id []int `json:"cart_id"`
}

type CartUpdateRequest struct {
	Id        int `json:"id"`
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type CartDeleteRequest struct {
	Id        int `json:"id"`
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
