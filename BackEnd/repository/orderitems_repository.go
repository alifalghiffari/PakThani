package repository

import (
	"context"
	"database/sql"
	"project-workshop/go-api-ecom/model/domain"
)

type OrderItemsRepository interface {
	FindById(ctx context.Context, tx *sql.Tx, Id int) (domain.OrderItems, error)
	FindByOrderId(ctx context.Context, tx *sql.Tx, orderId int) ([]domain.OrderItems, error)
	FindByUserId(ctx context.Context, tx *sql.Tx, userId int) ([]domain.OrderItems, error)
	Insert(ctx context.Context, tx *sql.Tx, orderItems domain.OrderItems) domain.OrderItems
}