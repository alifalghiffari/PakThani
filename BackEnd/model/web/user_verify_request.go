package web

type UserVerifyRequest struct {
	Email             string `json:"email" validate:"required,max=200,min=1"`
	VerificationToken string `json:"verification_token" validate:"required,max=200,min=1"`
}
