package router

import (
	"database/sql"
	addmiddleware "sirclo/api/addMiddleware"
	"sirclo/api/controller"
	"sirclo/api/repository"
	"sirclo/api/service"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(db *sql.DB) *echo.Echo {
	//new echo
	e := echo.New()
	//setting cors
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	// e.Pre(middleware.RemoveTrailingSlash(), middleware.Logger())
	authService := addmiddleware.AuthService()

	//User
	UserRepository := repository.NewRepositoryUser(db)
	UserService := service.NewUserService(UserRepository)
	UserController := controller.NewUserController(authService, UserService)

	e.POST("/login", UserController.AuthUserController)
	e.GET("/users", addmiddleware.AuthMiddleware(authService, UserService, UserController.GetUsersController))
	e.GET("/users/:id", addmiddleware.AuthMiddleware(authService, UserService, UserController.GetUserController))
	e.POST("/users", UserController.CreateUserController)
	e.PUT("/users/:id", addmiddleware.AuthMiddleware(authService, UserService, UserController.UpdateUserController))
	e.DELETE("/users/:id", addmiddleware.AuthMiddleware(authService, UserService, UserController.DeleteUserController))

	//Product
	ProductRepository := repository.NewRepositoryProduct(db)
	ProductService := service.NewProductService(ProductRepository)
	ProductController := controller.NewProductController(ProductService)

	e.GET("products", ProductController.GetProductsController)
	e.GET("products/:id", ProductController.GetProductController)
	e.POST("products", addmiddleware.AuthMiddleware(authService, UserService, ProductController.CreateProductController))
	e.PUT("products/:id", addmiddleware.AuthMiddleware(authService, UserService, ProductController.UpdateProductController))
	e.DELETE("products/:id", addmiddleware.AuthMiddleware(authService, UserService, ProductController.DeleteProductController))

	//Cart
	CartRepository := repository.NewRepositoryCart(db)
	CartService := service.NewCartService(CartRepository)
	CartController := controller.NewCartController(CartService)

	e.GET("carts", addmiddleware.AuthMiddleware(authService, UserService, CartController.GetCartsController))
	e.GET("carts/:id", addmiddleware.AuthMiddleware(authService, UserService, CartController.GetCartController))
	e.POST("carts", addmiddleware.AuthMiddleware(authService, UserService, CartController.CreateCartController))
	e.PUT("carts/:id", addmiddleware.AuthMiddleware(authService, UserService, CartController.UpdateCartController))
	e.DELETE("carts/:id", addmiddleware.AuthMiddleware(authService, UserService, CartController.DeleteCartController))

	//Order
	OrderRepository := repository.NewRepositoryOrder(db)
	OrderService := service.NewOrderService(OrderRepository)
	OrderController := controller.NewOrderController(OrderService)

	e.POST("order/cart", addmiddleware.AuthMiddleware(authService, UserService, OrderController.CreateOrderCartController))
	e.POST("order/product", addmiddleware.AuthMiddleware(authService, UserService, OrderController.CreateOrderProductController))
	e.GET("order", addmiddleware.AuthMiddleware(authService, UserService, OrderController.GetOrdersController))
	e.GET("order/:id", addmiddleware.AuthMiddleware(authService, UserService, OrderController.GetOrderController))

	return e
}
