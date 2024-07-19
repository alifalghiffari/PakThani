package service

import (
	"context"
	"database/sql"
	"encoding/base64"
	"fmt"
	"os"
	"path"
	"project-workshop/go-api-ecom/exception"
	"project-workshop/go-api-ecom/helper"
	"project-workshop/go-api-ecom/model/domain"
	"project-workshop/go-api-ecom/model/web"
	"project-workshop/go-api-ecom/repository"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	base64Data, err := base64.StdEncoding.DecodeString(request.Icon)
	helper.PanicIfError(err)
	filename := fmt.Sprintf("%s.png", uuid.New().String())
	filepath := path.Join("D:/PakThani-TugasAkhir/FrontEnd/src/assets/Icon", filename)

	err = os.WriteFile(filepath, base64Data, 0644)
	helper.PanicIfError(err)

	category := domain.Category{
		Category: request.Category,
		Icon:     filename,
		Slug:     request.Slug,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	if request.Icon != "" {
		base64Data, err := base64.StdEncoding.DecodeString(request.Icon)
		helper.PanicIfError(err)
		filename := fmt.Sprintf("%s.png", uuid.New().String())

		oldFilepath := path.Join("D:/PakThani-TugasAkhir/FrontEnd/src/assets/Icon", category.Icon)
		if _, err := os.Stat(oldFilepath); err == nil {
			err = os.Remove(oldFilepath)
			helper.PanicIfError(err)
		}

		filepath := path.Join("D:/PakThani-TugasAkhir/FrontEnd/src/assets/Icon", filename)
		err = os.WriteFile(filepath, base64Data, 0644)
		helper.PanicIfError(err)

		category.Icon = filename
	}

	category.Category = request.Category
	category.Slug = request.Slug

	category = service.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)
}

func (service *CategoryServiceImpl) GetAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.GetAll(ctx, tx)

	return helper.ToCategoryResponses(categories)
}
