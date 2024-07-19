package app

import (
	// "net/http"
	"project-workshop/go-api-ecom/controller"
	"project-workshop/go-api-ecom/exception"
	"project-workshop/go-api-ecom/middleware"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController,
	productController controller.ProductController,
	accountController controller.AccountController,
	userController controller.UserController,
	cartController controller.CartController,
	addressController controller.AddressController,
	orderController controller.OrderController,
	orderItemsController controller.OrderItemsController) *httprouter.Router {
	router := httprouter.New()

	// Middleware
	authMiddleware := middleware.Middleware{}

	// Order
	router.GET("/api/orders/users", authMiddleware.ApplyMiddleware(orderController.FindOrderByUserId))
	router.GET("/api/orders/edit/:orderId", authMiddleware.ApplyMiddleware(orderController.FindOrderById))
	router.POST("/api/orders", authMiddleware.ApplyMiddleware(orderController.CreateOrder))
	router.PUT("/api/orders/:orderId", authMiddleware.ApplyAdminMiddleware(orderController.UpdateOrder))
	router.GET("/api/order", authMiddleware.ApplyAdminMiddleware(orderController.FindAll))

	// Order Items
	router.GET("/api/orderitems/edit/:orderItemId", authMiddleware.ApplyMiddleware(orderItemsController.FindOrderItemsById))
	router.GET("/api/orderitems/carts/:orderId", authMiddleware.ApplyMiddleware(orderItemsController.FindOrderItemsByOrderId))
	router.GET("/api/orderitems/admin/:orderId", authMiddleware.ApplyAdminMiddleware(orderItemsController.FindOrderItemsByOrderId))
	router.GET("/api/orderitems/users", authMiddleware.ApplyMiddleware(orderItemsController.FindOrderItemsByUserId))
	// router.POST("/api/orderitems", authMiddleware.ApplyMiddleware(orderItemsController.CreateOrderItems))

	// Keranjang
	router.POST("/api/carts", authMiddleware.ApplyMiddleware(cartController.AddToCart))
	router.PUT("/api/carts", authMiddleware.ApplyMiddleware(cartController.UpdateCart))
	router.DELETE("/api/remove/carts", authMiddleware.ApplyMiddleware(cartController.RemoveCart))
	router.DELETE("/api/carts", authMiddleware.ApplyMiddleware(cartController.DeleteCart))
	router.GET("/api/carts/edit", authMiddleware.ApplyMiddleware(cartController.FindById))
	router.GET("/api/carts/users", authMiddleware.ApplyMiddleware(cartController.FindByUserId))
	router.GET("/api/carts", authMiddleware.ApplyMiddleware(cartController.FindAll))

	// Kategori
	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/category", authMiddleware.ApplyAdminMiddleware(categoryController.GetAll))
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", authMiddleware.ApplyAdminMiddleware(categoryController.Create))
	router.PUT("/api/categories/:categoryId", authMiddleware.ApplyAdminMiddleware(categoryController.Update))
	router.DELETE("/api/categories/:categoryId", authMiddleware.ApplyAdminMiddleware(categoryController.Delete))

	// Produk
	router.GET("/api/products", productController.FindAll)
	router.GET("/api/products/:productId", productController.FindById)
	router.POST("/api/products", authMiddleware.ApplyAdminMiddleware(productController.Create))
	router.PUT("/api/products/:productId", authMiddleware.ApplyAdminMiddleware(productController.Update))
	router.DELETE("/api/products/:productId", authMiddleware.ApplyAdminMiddleware(productController.Delete))
	router.DELETE("/api/img/:imgId", authMiddleware.ApplyAdminMiddleware(productController.DeleteImg))

	// Akun
	router.GET("/api/accounts", authMiddleware.ApplyMiddleware(accountController.UserDetailByID))
	router.POST("/api/address", authMiddleware.ApplyMiddleware(addressController.AddAddress))
	router.PUT("/api/address/update", authMiddleware.ApplyMiddleware(addressController.UpdateAddress))
	router.GET("/api/address/users", authMiddleware.ApplyMiddleware(addressController.FindByUserId))

	// Pengguna
	router.POST("/api/users/register", userController.Register)
	router.POST("/api/users/login", userController.Login)
	router.POST("/api/users/verify-email", userController.VerifyEmail)
	router.GET("/api/users/login-google", userController.LoginGoogle)

	// Password
	router.POST("/api/users/send-reset", userController.SendResetPassword)
	router.POST("/api/users/verify-reset-password", userController.VerifyResetPassword)
	router.POST("/api/users/reset-password", authMiddleware.ApplyMiddleware(userController.ResetPassword))

	router.PanicHandler = exception.ErrorHandler

	return router
}
