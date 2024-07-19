package main

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"

	"project-workshop/go-api-ecom/app"
	"project-workshop/go-api-ecom/controller"
	"project-workshop/go-api-ecom/helper"
	"project-workshop/go-api-ecom/repository"
	"project-workshop/go-api-ecom/service"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db := app.NewDB()
	validate := validator.New()
	smtpAuth := smtp.PlainAuth("PakThani", os.Getenv("SMTP_EMAIL"), os.Getenv("SMTP_PASSWORD"), os.Getenv("SMTP_HOST"))
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	
	googleConfig := service.GoogleConfig()

	categoryRepository := repository.NewCategoryRepository()
	productRepository := repository.NewProductRepository()
	userRepository := repository.NewUserRepository()
	accountRepository := repository.NewUserRepository()
	cartRepository := repository.NewCartRepositoryImpl()
	orderItemsRepository := repository.NewOrderItemsRepositoryImpl()
	orderRepository := repository.NewOrderRepositoryImpl()
	imageRepository := repository.NewImageRepositoryImpl()

	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	productService := service.NewProductService(productRepository, imageRepository, db, validate)
	productController := controller.NewProductController(productService)

	userService := service.NewUserService(userRepository, db, validate, googleConfig, smtpAuth, smtpHost, smtpPort)
	userController := controller.NewUserController(userService)

	accountService := service.NewAccountService(accountRepository, db)
	accountController := controller.NewAccountController(accountService)

	cartService := service.NewCartService(cartRepository, userRepository, productRepository, db, validate)
	cartController := controller.NewCartController(cartService)

	addressRepository := repository.NewAddressRepositoryImpl()
	addressService := service.NewAddressService(addressRepository, userRepository, db, validate)
	addressController := controller.NewAddressController(addressService)

	orderItemsService := service.NewOrderItemsService(orderItemsRepository, userRepository, productRepository, cartRepository, orderRepository, db, validate)
	orderItemsController := controller.NewOrderItemsController(orderItemsService)

	orderService := service.NewOrderService(orderRepository, userRepository, cartRepository, orderItemsRepository, db, validate)
	orderController := controller.NewOrderController(orderService)

	router := app.NewRouter(categoryController, productController, accountController, userController, cartController, addressController, orderController, orderItemsController)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	handler := corsHandler.Handler(router)

	fmt.Println("Server listening on port http://localhost:3000/")

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: handler,
	}

	err = server.ListenAndServe()
	helper.PanicIfError(err)
}
