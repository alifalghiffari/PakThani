package web

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email,max=200,min=1"`
	Password string `json:"password" validate:"required,max=200,min=1"`
}
