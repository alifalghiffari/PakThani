package web

type ImageResponse struct {
	Id        int    `json:"id"`
	ProductId int    `json:"product_id"`
	Image     string `json:"image"`
}
