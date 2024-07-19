package repository

import (
	"context"
	"database/sql"
	"project-workshop/go-api-ecom/model/domain"
)

type UserRepository interface {
	Register(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Login(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.User
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error)
	FindByEmailGoogle(ctx context.Context, tx *sql.Tx, email string) (domain.User, error)
	FindByRole(ctx context.Context, tx *sql.Tx, role bool) (domain.User, error)
	VerifyEmail(ctx context.Context, tx *sql.Tx, verificationtoken string) (domain.User, error)
	UpdateVerificationStatus(ctx context.Context, tx *sql.Tx, email string, verificationtoken string) (domain.User, error)
	InsertResetPassword(ctx context.Context, tx *sql.Tx, token domain.ResetPassword) (domain.ResetPassword, error)
	FindByToken(ctx context.Context, tx *sql.Tx, token string) (domain.ResetPassword, error)
	DeletedByUserId(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error)
	UpdatePassword(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
}
