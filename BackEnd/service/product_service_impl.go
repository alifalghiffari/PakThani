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
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	ImageRepository   repository.ImageRepository
	DB                *sql.DB
	Validate          *validator.Validate
	ImagePath         string
}

func NewProductService(productRepository repository.ProductRepository, imageRepository repository.ImageRepository, DB *sql.DB, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		ImageRepository:   imageRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *ProductServiceImpl) Create(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	base64Data, err := base64.StdEncoding.DecodeString(request.Image)
	helper.PanicIfError(err)
	filename := fmt.Sprintf("%s.png", uuid.New().String())
	filepath := path.Join("D:/PakThani-TugasAkhir/FrontEnd/src/assets/img-main", filename)

	err = os.WriteFile(filepath, base64Data, 0644)
	helper.PanicIfError(err)

	product := domain.Product{
		Name:        request.Name,
		Description: request.Description,
		Image:       filename,
		Price:       request.Price,
		CategoryId:  request.CategoryId,
		Quantity:    request.Quantity,
		Slug:        request.Slug,
	}

	product = service.ProductRepository.Save(ctx, tx, product)

	var images []domain.Image
	for _, base64String := range request.Images {
		base64Data, err := base64.StdEncoding.DecodeString(base64String)
		helper.PanicIfError(err)

		//timestamp := time.Now().Unix()
		//filename := fmt.Sprintf("%d.png", timestamp)

		filename := fmt.Sprintf("%s.png", uuid.New().String())
		filepath := path.Join("D:/PakThani-TugasAkhir/FrontEnd/src/assets/img", filename)
		err = os.WriteFile(filepath, base64Data, 0644)
		helper.PanicIfError(err)

		//hasher := sha256.New()
		//hasher.Write(base64Data)
		//hash := hex.EncodeToString(hasher.Sum(nil))
		//filename := fmt.Sprintf("%s.png", hash)

		image := domain.Image{
			ProductId:  product.Id,
			Image:      filename,
			Created_at: time.Now(),
			Updated_at: time.Now(),
		}

		images = append(images, image)
	}

	for _, img := range images {
		service.ImageRepository.AddImage(ctx, tx, img)
	}

	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	existingImages, err := service.ImageRepository.FindByProductId(ctx, tx, product.Id)
	helper.PanicIfError(err)

	if request.Image != "" {
		base64Data, err := base64.StdEncoding.DecodeString(request.Image)
		helper.PanicIfError(err)
		filename := fmt.Sprintf("%s.png", uuid.New().String())

		oldFilepath := path.Join("D:/PakThani-TugasAkhir/FrontEnd/src/assets/img-main", product.Image)
		if _, err := os.Stat(oldFilepath); err == nil {
			err = os.Remove(oldFilepath)
			helper.PanicIfError(err)
		}

		filepath := path.Join("D:/PakThani-TugasAkhir/FrontEnd/src/assets/img-main", filename)
		err = os.WriteFile(filepath, base64Data, 0644)
		helper.PanicIfError(err)

		product.Image = filename
	}

	product.Name = request.Name
	product.Description = request.Description
	product.Price = request.Price
	product.CategoryId = request.CategoryId
	product.Quantity = request.Quantity
	product.Slug = request.Slug

	product = service.ProductRepository.Update(ctx, tx, product)

	for idx, base64String := range request.Images {
		base64Data, err := base64.StdEncoding.DecodeString(base64String)
		helper.PanicIfError(err)

		filename := fmt.Sprintf("%s.png", uuid.New().String())

		image := domain.Image{
			ProductId:  product.Id,
			Image:      filename,
			Updated_at: time.Now(),
		}

		if idx < len(existingImages) {
			image.Id = existingImages[idx].Id
			image.Created_at = existingImages[idx].Created_at

			oldFilepath := path.Join("D:/PakThani-TugasAkhir/FrontEnd/src/assets/img", existingImages[idx].Image)
			err = os.Remove(oldFilepath)
			helper.PanicIfError(err)

			filepath := path.Join("D:/PakThani-TugasAkhir/FrontEnd/src/assets/img", filename)
			err = os.WriteFile(filepath, base64Data, 0644)
			helper.PanicIfError(err)

			service.ImageRepository.UpdateImage(ctx, tx, image)
		} else {
			image.Created_at = time.Now()

			filepath := path.Join("D:/PakThani-TugasAkhir/FrontEnd/src/assets/img", filename)
			err = os.WriteFile(filepath, base64Data, 0644)
			helper.PanicIfError(err)

			service.ImageRepository.AddImage(ctx, tx, image)
		}
	}

	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) Delete(ctx context.Context, productId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.ProductRepository.Delete(ctx, tx, product)
}

func (service *ProductServiceImpl) FindById(ctx context.Context, productId int) web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) FindAll(ctx context.Context) []web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	products := service.ProductRepository.FindAll(ctx, tx)

	return helper.ToProductResponses(products)
}

func (service *ProductServiceImpl) DeleteImg(ctx context.Context, id int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	image := service.ImageRepository.FindById(ctx, tx, id)

	removeFilepath := path.Join("D:/PakThani-TugasAkhir/FrontEnd/src/assets/img", image.Image)
	if _, err := os.Stat(removeFilepath); err == nil {
		err = os.Remove(removeFilepath)
		helper.PanicIfError(err)
	}

	service.ImageRepository.DeleteImage(ctx, tx, id)
}

//func saveImage(file *multipart.FileHeader, basePath string, productId int) (string, error) {
// 	src, err := file.Open()
// 	if err != nil {
// 		return "", err
// 	}
// 	defer src.Close()

// 	// Membuat direktori jika belum ada
// 	folderPath := filepath.Join(basePath, fmt.Sprintf("product-%d", productId))
// 	err = os.MkdirAll(folderPath, os.ModePerm)
// 	if err != nil {
// 		return "", err
// 	}

// 	// Mendapatkan ekstensi file
// 	fileExtension := strings.TrimPrefix(filepath.Ext(file.Filename), ".")

// 	// Membuat nama file baru
// 	filename := fmt.Sprintf("%d.%s", time.Now().UnixNano(), fileExtension)
// 	filePath := filepath.Join(folderPath, filename)

// 	dst, err := os.Create(filePath)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer dst.Close()

// 	_, err = io.Copy(dst, src)
// 	if err != nil {
// 		return "", err
// 	}

// 	return filename, nil
// }
