package web

type ProductCreateRequest struct {
	Name        string   `json:"name" validate:"required,max=200,min=1"`
	Description string   `json:"description" validate:"required"`
	Image       string   `json:"image" validate:"required"`
	Price       int      `json:"price" validate:"required,min=1"`
	CategoryId  int      `json:"category_id" validate:"required"`
	Quantity    int      `json:"quantity" validate:"required"`
	Slug        string   `json:"slug" validate:"required"`
	Images      []string `json:"images"`
}

type ProductUpdateRequest struct {
	Id          int      `json:"id" validate:"required"`
	Name        string   `json:"name" validate:"required,max=200,min=1"`
	Description string   `json:"description" validate:"required"`
	Image       string   `json:"image"`
	Price       int      `json:"price" validate:"required,min=1"`
	CategoryId  int      `json:"category_id" validate:"required"`
	Quantity    int      `json:"quantity" validate:"required"`
	Slug        string   `json:"slug" validate:"required"`
	Images      []string `json:"images"`
}

type ProductDeleteRequest struct {
	Id int `json:"id" validate:"required"`
}
