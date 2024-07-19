package repository

import (
	"context"
	"database/sql"
	"project-workshop/go-api-ecom/model/domain"
)

type AddressRepository interface {
	AddAddress(ctx context.Context, tx *sql.Tx, address domain.Address) domain.Address
	UpdateAddress(ctx context.Context, tx *sql.Tx, address domain.Address) domain.Address
	FindByUserId(ctx context.Context, tx *sql.Tx, userId int) ([]domain.Address, error)
}
