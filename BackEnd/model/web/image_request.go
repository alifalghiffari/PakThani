package web

type ImageCreateRequest struct {
	Images []string `json:"images"`
}

type ImageUpdateRequest struct {
	Image []string `json:"image"`
}
