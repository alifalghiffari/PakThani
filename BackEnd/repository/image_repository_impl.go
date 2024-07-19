package repository

import (
	"context"
	"database/sql"
	"fmt"
	"project-workshop/go-api-ecom/helper"
	"project-workshop/go-api-ecom/model/domain"
)

type ImageRepositoryImpl struct {
}

func NewImageRepositoryImpl() ImageRepository {
	return &ImageRepositoryImpl{}
}

func (repository *ImageRepositoryImpl) AddImage(ctx context.Context, tx *sql.Tx, image domain.Image) domain.Image {
	SQL := "INSERT INTO images(product_id, image, created_at, updated_at) VALUES (?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, image.ProductId, image.Image, image.Created_at, image.Updated_at)
	if err != nil {
		fmt.Println("err", err)
	}
	//helper.PanicIfError(err)

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("err1", err)
	}
	//helper.PanicIfError(err)

	image.Id = int(id)

	return image
}

func (repository *ImageRepositoryImpl) UpdateImage(ctx context.Context, tx *sql.Tx, image domain.Image) domain.Image {
	SQL := "UPDATE images SET image = ?, updated_at = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, image.Image, image.Updated_at, image.Id)
	helper.PanicIfError(err)

	return image
}

func (repository *ImageRepositoryImpl) DeleteImage(ctx context.Context, tx *sql.Tx, id int) domain.Image {
	SQL := "DELETE FROM images WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, id)
	helper.PanicIfError(err)

	return domain.Image{}
}

func (repository *ImageRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) domain.Image {
	SQL := "SELECT id, product_id, image FROM images WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)
	defer rows.Close()

	var img domain.Image
	if rows.Next() {
		err := rows.Scan(&img.Id, &img.ProductId, &img.Image)
		helper.PanicIfError(err)
	}

	return img
}

func (repository *ImageRepositoryImpl) FindByProductId(ctx context.Context, tx *sql.Tx, productId int) ([]domain.Image, error) {
	SQL := "SELECT id, product_id, image FROM images WHERE product_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, productId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var images []domain.Image
	for rows.Next() {
		var img domain.Image
		err := rows.Scan(&img.Id, &img.ProductId, &img.Image)
		if err != nil {
			return nil, err
		}
		images = append(images, img)
	}

	return images, nil
}
