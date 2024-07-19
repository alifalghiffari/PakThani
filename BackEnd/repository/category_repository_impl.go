package repository

import (
	"context"
	"database/sql"
	"project-workshop/go-api-ecom/helper"
	"project-workshop/go-api-ecom/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "insert into category(category, icon, slug) values (?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, category.Category, category.Icon, category.Slug)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "update category set category = ?, icon = ?, slug = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Category, category.Icon, category.Slug, category.Id)
	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "delete from category where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	SQL := `
		SELECT c.id, c.category, c.icon, p.id AS product_id, p.name AS product_name, p.image AS product_image, p.description AS product_description, p.price AS product_price, p.category_id, c.category AS product_category
		FROM category c
		JOIN product p ON c.id = p.category_id
		WHERE c.id = ?
	`
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)

	var category domain.Category
	category.Products = make([]domain.Product, 0)
	for rows.Next() {
		product := domain.Product{}
		err := rows.Scan(&category.Id, &category.Category, &category.Icon, &product.Id, &product.Name, &product.Image, &product.Description, &product.Price, &product.CategoryId, &product.Category.Category)
		helper.PanicIfError(err)
		category.Products = append(category.Products, product)
	}
	return category, nil
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := `
		SELECT c.id, c.category, c.icon, p.id AS product_id, p.name AS product_name, p.image AS product_image, p.description AS product_description, p.price AS product_price, p.category_id, c.category AS product_category
		FROM category c
		JOIN product p ON c.id = p.category_id
	`
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var categoriesMap = make(map[int]*domain.Category)
	for rows.Next() {
		category := domain.Category{}
		product := domain.Product{}
		err := rows.Scan(&category.Id, &category.Category, &category.Icon, &product.Id, &product.Name, &product.Image, &product.Description, &product.Price, &product.CategoryId, &product.Category.Category)
		helper.PanicIfError(err)

		if _, ok := categoriesMap[category.Id]; !ok {
			categoriesMap[category.Id] = &category
		}

		categoriesMap[category.Id].Products = append(categoriesMap[category.Id].Products, product)
	}

	var categories []domain.Category
	for _, category := range categoriesMap {
		categories = append(categories, *category)
	}

	return categories
}

func (repository *CategoryRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "SELECT id, category FROM category"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Category)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}

	return categories
}
