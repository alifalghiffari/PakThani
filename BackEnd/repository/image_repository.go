package repository

import (
	"context"
	"database/sql"
	"project-workshop/go-api-ecom/model/domain"
)

type ImageRepository interface {
	AddImage(ctx context.Context, tx *sql.Tx, image domain.Image) domain.Image
	UpdateImage(ctx context.Context, tx *sql.Tx, image domain.Image) domain.Image
	DeleteImage(ctx context.Context, tx *sql.Tx, id int) domain.Image
	FindById(ctx context.Context, tx *sql.Tx, id int) domain.Image
	FindByProductId(ctx context.Context, tx *sql.Tx, productId int) ([]domain.Image, error)
}
