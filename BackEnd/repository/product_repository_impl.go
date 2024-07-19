package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"project-workshop/go-api-ecom/helper"
	"project-workshop/go-api-ecom/model/domain"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "insert into product(name, description, image, price, category_id, quantity, slug) values (?, ?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, product.Name, product.Description, product.Image, product.Price, product.CategoryId, product.Quantity, product.Slug)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	product.Id = int(id)
	return product
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "update product set name = ?, description = ?, image = ?, price = ?, category_id = ?, quantity = ?, slug = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Name, product.Description, product.Image, product.Price, product.CategoryId, product.Quantity, product.Slug, product.Id)
	helper.PanicIfError(err)

	return product
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, product domain.Product) {
	SQL := "delete from product where id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Id)
	helper.PanicIfError(err)
}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId int) (domain.Product, error) {
	SQL := `
		SELECT p.id, p.name, p.image, p.description, p.price, p.category_id, p.quantity, p.slug, c.category, COALESCE(i.id, 0), i.image
		FROM product p
		JOIN category c ON p.category_id = c.id
		LEFT JOIN images i ON p.id = i.product_id
		WHERE p.id = ?
	`
	rows, err := tx.QueryContext(ctx, SQL, productId)
	if err != nil {
		return domain.Product{}, err
	}
	defer rows.Close()

	product := domain.Product{}
	images := []domain.Image{}

	productSet := false

	for rows.Next() {
		var img domain.Image
		var category string
		var imgID sql.NullInt64
		var image sql.NullString

		err := rows.Scan(
			&product.Id, &product.Name, &product.Image, &product.Description,
			&product.Price, &product.CategoryId, &product.Quantity, &product.Slug,
			&category, &imgID, &image,
		)
		if err != nil {
			return domain.Product{}, err
		}

		if !productSet {
			product.Category = domain.Category{Category: category}
			productSet = true
		}

		if imgID.Valid {
			img.Id = int(imgID.Int64)
			img.ProductId = product.Id
			img.Image = string(image.String)
			images = append(images, img)
		}
	}

	if !productSet {
		return domain.Product{}, errors.New("product is not found")
	}

	product.Images = images
	return product, nil
}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Product {
	SQL := `
		SELECT p.id, p.name, p.image, p.description, p.price, p.category_id, p.quantity, p.slug, c.category, i.id, i.product_id, i.image
		FROM product p
		JOIN category c ON p.category_id = c.id
		LEFT JOIN images i ON p.id = i.product_id
	`
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		fmt.Println("err", err)
	}
	defer rows.Close()

	var products []domain.Product
	productMap := make(map[int]*domain.Product)
	for rows.Next() {
		product := domain.Product{}
		image := domain.Image{}
		var imgId, prdId *int
		var img *string
		err := rows.Scan(&product.Id, &product.Name, &product.Image, &product.Description, &product.Price, &product.CategoryId, &product.Quantity, &product.Slug, &product.Category.Category, &imgId, &prdId, &img)
		helper.PanicIfError(err)

		if imgId != nil {
			image.Id = *imgId
		}

		if prdId != nil {
			image.ProductId = *prdId
		}

		if img != nil {
			image.Image = *img
		}

		if _, ok := productMap[product.Id]; !ok {
			productMap[product.Id] = &product
		}
		productMap[product.Id].Images = append(productMap[product.Id].Images, image)
	}

	for _, product := range productMap {
		products = append(products, *product)
	}

	if len(products) == 0 {
		fmt.Println("err2", err)
	}

	return products
}
