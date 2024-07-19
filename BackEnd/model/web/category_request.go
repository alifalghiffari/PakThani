package web

type CategoryCreateRequest struct {
	Category string `json:"category" validate:"required,min=1,max=100"`
	Icon     string `json:"image" validate:"required"`
	Slug     string `json:"slug" validate:"required"`
}

type CategoryUpdateRequest struct {
	Id       int    `json:"id" validate:"required"`
	Category string `json:"category" validate:"required,min=1,max=100"`
	Icon     string `json:"image"`
	Slug     string `json:"slug" validate:"required"`
}

type CategoryDeleteRequest struct {
	Id int `validate:"required"`
}
