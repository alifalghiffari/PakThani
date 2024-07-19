package web

type CartResponse struct {
	Id        int               `json:"id"`
	UserId    int               `json:"user_id"`
	ProductId int               `json:"product_id"`
	Product   []ProductResponse `json:"product"`
	Quantity  int               `json:"quantity"`
	Price     int               `json:"price"`
}
