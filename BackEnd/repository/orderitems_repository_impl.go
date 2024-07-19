package repository

import (
	"context"
	"database/sql"
	"project-workshop/go-api-ecom/helper"
	"project-workshop/go-api-ecom/model/domain"
)

type OrderItemsRepositoryImpl struct {
}

func NewOrderItemsRepositoryImpl() OrderItemsRepository {
	return &OrderItemsRepositoryImpl{}
}

func (repository *OrderItemsRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, Id int) (domain.OrderItems, error) {
	query := "SELECT id, order_id, product_id, quantity, price FROM order_items WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, Id)
	helper.PanicIfError(err)

	orderItem := domain.OrderItems{}
	if rows.Next() {
		err = rows.Scan(&orderItem.Id, &orderItem.OrderId, &orderItem.ProductId, &orderItem.Quantity, &orderItem.Price)
		helper.PanicIfError(err)
	}

	return orderItem, nil
}

func (repository *OrderItemsRepositoryImpl) FindByOrderId(ctx context.Context, tx *sql.Tx, orderId int) ([]domain.OrderItems, error) {
	//query := "SELECT id, order_id, product_id, quantity, price FROM order_items WHERE order_id = ?"
	query := `
			SELECT o.id, o.order_id, o.product_id, o.quantity, o.price, p.id, p.name, p.image, p.price
			FROM order_items o
			JOIN product p ON o.product_id = p.id
			WHERE o.order_id = ?
	`
	rows, err := tx.QueryContext(ctx, query, orderId)
	helper.PanicIfError(err)

	var orderItems []domain.OrderItems
	for rows.Next() {
		orderItem := domain.OrderItems{}
		err = rows.Scan(&orderItem.Id, &orderItem.OrderId, &orderItem.ProductId, &orderItem.Quantity, &orderItem.Price, &orderItem.Product.Id, &orderItem.Product.Name, &orderItem.Product.Image, &orderItem.Product.Price)
		helper.PanicIfError(err)
		orderItems = append(orderItems, orderItem)
	}

	return orderItems, nil
}

func (repository *OrderItemsRepositoryImpl) FindByUserId(ctx context.Context, tx *sql.Tx, userId int) ([]domain.OrderItems, error) {
	query := "SELECT id, order_id, product_id, quantity, price FROM order_items WHERE user_id = ?"
	rows, err := tx.QueryContext(ctx, query, userId)
	helper.PanicIfError(err)

	var orderItems []domain.OrderItems
	for rows.Next() {
		orderItem := domain.OrderItems{}
		err = rows.Scan(&orderItem.Id, &orderItem.OrderId, &orderItem.ProductId, &orderItem.Quantity, &orderItem.Price)
		helper.PanicIfError(err)
		orderItems = append(orderItems, orderItem)
	}

	return orderItems, nil
}

func (repository *OrderItemsRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, orderItems domain.OrderItems) domain.OrderItems {
	query := "INSERT INTO order_items (order_id, product_id, quantity, price) VALUES (?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, query, orderItems.OrderId, orderItems.ProductId, orderItems.Quantity, orderItems.Price)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	orderItems.Id = int(id)
	return orderItems
}
