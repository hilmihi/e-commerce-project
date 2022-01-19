package controller

import (
	"fmt"
	"net/http"
	"sirclo/api/entities"
	"sirclo/api/helper"
	"sirclo/api/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductHF struct {
	ProductService service.ServiceProduct
}

func NewProductController(ProductService service.ServiceProduct) *ProductHF {
	return &ProductHF{ProductService}
}

//Product get all
func (u *ProductHF) GetProductsController(c echo.Context) error {
	Products, err := u.ProductService.ServiceProductsGet()
	if err != nil {
		response := helper.ResponsesFormat("Failed to fetch Product data", http.StatusOK, nil)
		return c.JSON(http.StatusOK, response)
	}
	var data []helper.ProductFormatter
	for i := 0; i < len(Products); i++ {
		formatter := helper.FormatProduct(Products[i])
		data = append(data, formatter)
	}

	response := helper.ResponsesFormat("Success to fetch Product data", http.StatusOK, data)

	return c.JSON(http.StatusOK, response)
}

//Product get by id
func (u *ProductHF) GetProductController(c echo.Context) error {
	ProductId, errId := strconv.Atoi(c.Param("id"))
	if errId != nil {
		errResp := helper.ResponsesFormat("Failed to convert id", http.StatusInternalServerError, nil)
		return c.JSON(http.StatusInternalServerError, errResp)
	}

	Product, err := u.ProductService.ServiceProductGet(ProductId)
	if err != nil {
		errResp := helper.ResponsesFormat("Failed to get Product by id", http.StatusBadRequest, nil)
		return c.JSON(http.StatusBadRequest, errResp)
	}

	formatter := helper.FormatProduct(Product)
	resp := helper.ResponsesFormat("Success to get Product by id", http.StatusOK, formatter)
	return c.JSON(http.StatusOK, resp)
}

// Product update
func (u *ProductHF) UpdateProductController(c echo.Context) error {
	ProductId, errId := strconv.Atoi(c.Param("id"))
	if errId != nil {
		errResp := helper.ResponsesFormat("failed to convert id", http.StatusBadRequest, nil)
		return c.JSON(http.StatusInternalServerError, errResp)

	}

	var updateInput helper.RequestProductUpdate
	if err := c.Bind(&updateInput); err != nil {
		fmt.Println("bind", err)
		errResp := helper.ResponsesFormat("Failed to update data", http.StatusBadRequest, nil)
		return c.JSON(http.StatusBadRequest, errResp)

	}

	userID := c.Get("currentUser").(entities.User)
	updateInput.UserID = userID

	update, err := u.ProductService.ServiceProductUpdate(ProductId, updateInput)
	if err != nil {
		fmt.Println("update", err)
		errResp := helper.ResponsesFormat("unauthorized", http.StatusUnauthorized, nil)
		return c.JSON(http.StatusUnauthorized, errResp)

	}

	formatter := helper.FormatProduct(update)
	resp := helper.ResponsesFormat("Success to update Product data", http.StatusOK, formatter)
	return c.JSON(http.StatusOK, resp)

}

//create Product

func (u *ProductHF) CreateProductController(c echo.Context) error {
	var createInput helper.RequestProductCreate
	if err := c.Bind(&createInput); err != nil {
		fmt.Println("CI", err)
		errResp := helper.ResponsesFormat("Failed to Create Product", http.StatusBadRequest, nil)
		return c.JSON(http.StatusBadRequest, errResp)
	}

	userID := c.Get("currentUser").(entities.User)
	createInput.UserID = userID
	createdProduct, err := u.ProductService.ServiceProductCreate(createInput)
	if err != nil {
		fmt.Println("create", err)
		errResp := helper.ResponsesFormat("Failed to create Product", http.StatusBadRequest, nil)
		return c.JSON(http.StatusBadRequest, errResp)
	}

	formatter := helper.FormatProduct(createdProduct)
	resp := helper.ResponsesFormat("Success create Product", http.StatusOK, formatter)
	return c.JSON(http.StatusOK, resp)

}

//delete Product
func (u *ProductHF) DeleteProductController(c echo.Context) error {
	ProductId, errId := strconv.Atoi(c.Param("id"))
	if errId != nil {
		// fmt.Println("id", errId)
		errResp := helper.ResponsesFormat("Failed to delete Product", http.StatusInternalServerError, nil)
		return c.JSON(http.StatusBadRequest, errResp)
	}

	userID := c.Get("currentUser").(entities.User)
	deleteProduct, err := u.ProductService.ServiceProductDelete(ProductId, userID.Id)
	// fmt.Printf("PI %d = UI %d", ProductId, userID.Id)
	if err != nil {
		fmt.Println("del", err)
		errResp := helper.ResponsesFormat("unauthorized", http.StatusUnauthorized, nil)
		return c.JSON(http.StatusUnauthorized, errResp)
	}

	formatter := helper.FormatProduct(deleteProduct)
	resp := helper.ResponsesFormat("Success delete Product", http.StatusOK, formatter)
	return c.JSON(http.StatusOK, resp)
}
