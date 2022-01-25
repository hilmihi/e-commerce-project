package controller_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	addmiddleware "sirclo/api/addMiddleware"
	"sirclo/api/config"
	"sirclo/api/controller"
	"sirclo/api/entities"
	"sirclo/api/helper"
	"sirclo/api/repository"
	"sirclo/api/service"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type mockProductRepository struct{}

func initDatabase() (*sql.DB, error) {
	_, err := config.InitDBTest("taktuku_project_test", "root:@tcp(localhost:3306)/?parseTime=True&&multiStatements=true")
	if err != nil {
		panic(err)
	}
	db, err := config.InitDB("root:@tcp(localhost:3306)/taktuku_project_test?parseTime=True&&multiStatements=true")
	return db, err
}

func InsertDataUser(db *sql.DB) error {
	UserRepository := repository.NewRepositoryUser(db)
	UserService := service.NewUserService(UserRepository)

	user := entities.User{
		Name:         "Alta",
		Password:     "12345",
		Email:        "alta@gmail.com",
		Birth_date:   "1997-12-12",
		Phone_number: "0897654312",
		Photo:        "asaaaaasa.jpg",
		Gender:       "Pria",
		Address:      "Jl Kenangan",
	}

	user.Password, _ = helper.HashPassword(user.Password)

	_, err := UserService.ServiceUserCreate(user)
	return err
}

func InsertDataProduct(db *sql.DB) error {
	ProductRepository := repository.NewRepositoryProduct(db)
	ProductService := service.NewProductService(ProductRepository)

	product := entities.Product{
		Name:        "product1",
		Price:       10000,
		Quantity:    10,
		Description: "Contoh product 1",
		Id_category: 1,
		Id_user:     1,
	}

	_, err := ProductService.ServiceProductCreate(product)
	return err
}

// 1. test create product
func TestCreateProductController(t *testing.T) {
	dbTest, err := initDatabase()
	if err != nil {
		panic(err)
	}
	InsertDataUser(dbTest)
	authService := addmiddleware.AuthService()
	UserRepository := repository.NewRepositoryUser(dbTest)
	UserService := service.NewUserService(UserRepository)
	globalToken, errCreateToken := authService.GenerateToken(1)

	t.Run("success create product", func(t *testing.T) {
		e := echo.New()
		token, err := globalToken, errCreateToken
		if err != nil {
			panic(err)
		}
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":        "product1",
			"price":       100000,
			"quantity":    10,
			"description": "jualan product 1",
			"id_category": 1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products")

		ProductRepository := repository.NewRepositoryProduct(dbTest)
		ProductService := service.NewProductService(ProductRepository)
		ProductController := controller.NewProductController(ProductService)

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, addmiddleware.AuthMiddleware(authService, UserService, ProductController.CreateProductController)(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "Successful Operation", response.Message)
		}

	})

	t.Run("failed binding", func(t *testing.T) {
		e := echo.New()
		token, err := globalToken, errCreateToken
		if err != nil {
			panic(err)
		}
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":        "product1",
			"price":       "100000",
			"quantity":    "10",
			"description": "jualan product 1",
			"id_category": 1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products")

		ProductRepository := repository.NewRepositoryProduct(dbTest)
		ProductService := service.NewProductService(ProductRepository)
		ProductController := controller.NewProductController(ProductService)

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, addmiddleware.AuthMiddleware(authService, UserService, ProductController.CreateProductController)(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "Failed to Binding Data", response.Message)
		}

	})

	t.Run("failed create product", func(t *testing.T) {
		e := echo.New()
		token, err := globalToken, errCreateToken
		if err != nil {
			panic(err)
		}
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":        "product1",
			"price":       200000,
			"quantity":    20,
			"description": "jualan product 1",
			"id_category": 4,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products")

		ProductRepository := repository.NewRepositoryProduct(dbTest)
		ProductService := service.NewProductService(ProductRepository)
		ProductController := controller.NewProductController(ProductService)

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, addmiddleware.AuthMiddleware(authService, UserService, ProductController.CreateProductController)(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "Failed to create Product", response.Message)
		}

	})
}

// 2. test update product
func TestUpdateProductController(t *testing.T) {
	dbTest, err := initDatabase()
	if err != nil {
		panic(err)
	}
	InsertDataUser(dbTest)
	InsertDataProduct(dbTest)
	authService := addmiddleware.AuthService()
	UserRepository := repository.NewRepositoryUser(dbTest)
	UserService := service.NewUserService(UserRepository)
	globalToken, errCreateToken := authService.GenerateToken(1)

	t.Run("success update product", func(t *testing.T) {
		e := echo.New()
		token, err := globalToken, errCreateToken
		if err != nil {
			panic(err)
		}
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":        "product1 edited",
			"price":       100000,
			"quantity":    10,
			"description": "jualan product 1 edited",
			"id_category": 1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		ProductRepository := repository.NewRepositoryProduct(dbTest)
		ProductService := service.NewProductService(ProductRepository)
		ProductController := controller.NewProductController(ProductService)

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, addmiddleware.AuthMiddleware(authService, UserService, ProductController.UpdateProductController)(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "Successful Operation", response.Message)
		}

	})

	t.Run("failed convert id", func(t *testing.T) {
		e := echo.New()
		token, err := globalToken, errCreateToken
		if err != nil {
			panic(err)
		}
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":        "product1 edited",
			"price":       100000,
			"quantity":    10,
			"description": "jualan product 1 edited",
			"id_category": 1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("satu")

		ProductRepository := repository.NewRepositoryProduct(dbTest)
		ProductService := service.NewProductService(ProductRepository)
		ProductController := controller.NewProductController(ProductService)

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, addmiddleware.AuthMiddleware(authService, UserService, ProductController.UpdateProductController)(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusInternalServerError, res.Code)
			assert.Equal(t, "failed to convert id", response.Message)
		}

	})

	t.Run("failed update product", func(t *testing.T) {
		e := echo.New()
		token, err := globalToken, errCreateToken
		if err != nil {
			panic(err)
		}
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":        "product1 edited",
			"price":       100000,
			"quantity":    10,
			"description": "jualan product 1 edited",
			"id_category": 4,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		ProductRepository := repository.NewRepositoryProduct(dbTest)
		ProductService := service.NewProductService(ProductRepository)
		ProductController := controller.NewProductController(ProductService)

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, addmiddleware.AuthMiddleware(authService, UserService, ProductController.UpdateProductController)(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "Failed to update data", response.Message)
		}

	})

	t.Run("failed bind product", func(t *testing.T) {
		e := echo.New()
		token, err := globalToken, errCreateToken
		if err != nil {
			panic(err)
		}
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":        "product1 edited",
			"price":       "100000",
			"quantity":    "10",
			"description": "jualan product 1 edited",
			"id_category": "4",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		ProductRepository := repository.NewRepositoryProduct(dbTest)
		ProductService := service.NewProductService(ProductRepository)
		ProductController := controller.NewProductController(ProductService)

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, addmiddleware.AuthMiddleware(authService, UserService, ProductController.UpdateProductController)(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "Failed to bind data", response.Message)
		}

	})
}

// 3. test delete product
func TestDeleteProductController(t *testing.T) {
	dbTest, err := initDatabase()
	if err != nil {
		panic(err)
	}
	InsertDataUser(dbTest)
	InsertDataProduct(dbTest)
	authService := addmiddleware.AuthService()
	UserRepository := repository.NewRepositoryUser(dbTest)
	UserService := service.NewUserService(UserRepository)
	globalToken, errCreateToken := authService.GenerateToken(1)

	t.Run("success delete product", func(t *testing.T) {
		e := echo.New()
		token, err := globalToken, errCreateToken
		if err != nil {
			panic(err)
		}

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		ProductRepository := repository.NewRepositoryProduct(dbTest)
		ProductService := service.NewProductService(ProductRepository)
		ProductController := controller.NewProductController(ProductService)

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, addmiddleware.AuthMiddleware(authService, UserService, ProductController.DeleteProductController)(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "Successful Operation", response.Message)
		}

	})

	t.Run("failed convert id", func(t *testing.T) {
		e := echo.New()
		token, err := globalToken, errCreateToken
		if err != nil {
			panic(err)
		}

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("satu")

		ProductRepository := repository.NewRepositoryProduct(dbTest)
		ProductService := service.NewProductService(ProductRepository)
		ProductController := controller.NewProductController(ProductService)

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, addmiddleware.AuthMiddleware(authService, UserService, ProductController.DeleteProductController)(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "failed to convert id", response.Message)
		}

	})

	t.Run("failed delete product", func(t *testing.T) {
		e := echo.New()
		token, err := globalToken, errCreateToken
		if err != nil {
			panic(err)
		}

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		ProductRepository := repository.NewRepositoryProduct(dbTest)
		ProductService := service.NewProductService(ProductRepository)
		ProductController := controller.NewProductController(ProductService)

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, addmiddleware.AuthMiddleware(authService, UserService, ProductController.DeleteProductController)(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusInternalServerError, res.Code)
			assert.Equal(t, "Failed to delete Product", response.Message)
		}

	})
}
