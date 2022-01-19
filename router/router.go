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
	e.Pre(middleware.RemoveTrailingSlash(), middleware.Logger())
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

	return e
}
