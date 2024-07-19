package service

import (
	"context"
	"project-workshop/go-api-ecom/model/web"
)

type UserService interface {
	Register(ctx context.Context, request web.UserCreateRequest) web.UserResponse
	VerifyEmail(ctx context.Context, request web.UserVerifyRequest) web.UserResponse
	Login(ctx context.Context, request web.UserLoginRequest) (web.UserResponse, error)
	LoginGoogle(ctx context.Context, code string) (web.UserResponse, error)
	SendResetPassword(ctx context.Context, request web.ForgotPasswordRequest) web.UserResponse
	VerifyResetPassword(ctx context.Context, request web.ResetPasswordRequest) web.UserResponse
    ResetPassword(ctx context.Context, request web.ChangePasswordRequest, Id int) web.UserResponse
}
