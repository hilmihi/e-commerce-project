package controller

import (
	"fmt"
	"net/http"
	"sirclo/api/entities"
	"sirclo/api/helper"
	"sirclo/api/service"

	"github.com/labstack/echo/v4"
)

type OrderHF struct {
	OrderService service.ServiceOrder
}

func NewOrderController(OrderService service.ServiceOrder) *OrderHF {
	return &OrderHF{OrderService}
}

//Order get all
func (u *OrderHF) GetOrdersController(c echo.Context) error {
	userID := c.Get("currentUser").(entities.User)
	Orders, err := u.OrderService.ServiceOrdersGet(userID.Id)
	if err != nil {
		fmt.Println(err)
		response := helper.ResponsesFormat("Failed to fetch Product data", http.StatusOK, err)
		return c.JSON(http.StatusOK, response)
	}

	return c.JSON(http.StatusOK, Orders)
}

//create order by Cart
func (u *OrderHF) CreateOrderCartController(c echo.Context) error {
	var createInput helper.RequestOrderCart
	if err := c.Bind(&createInput); err != nil {
		fmt.Println("CI", err)
		errResp := helper.ResponsesFormat("Failed to Create Cart", http.StatusBadRequest, nil)
		return c.JSON(http.StatusBadRequest, errResp)
	}

	userID := c.Get("currentUser").(entities.User)

	createdCart, err := u.OrderService.ServiceOrderCreateByCart(userID.Id, createInput)
	if err != nil {
		fmt.Println("create", err)
		errResp := helper.ResponsesFormat("Failed to create Cart", http.StatusBadRequest, nil)
		return c.JSON(http.StatusBadRequest, errResp)
	}

	resp := helper.ResponsesFormat("Successful Operation", http.StatusOK, createdCart)
	return c.JSON(http.StatusOK, resp)

}

//create order by Product
func (u *OrderHF) CreateOrderProductController(c echo.Context) error {
	var createInput helper.RequestOrderProduct
	if err := c.Bind(&createInput); err != nil {
		fmt.Println("CI", err)
		errResp := helper.ResponsesFormat("Failed to Create Cart", http.StatusBadRequest, nil)
		return c.JSON(http.StatusBadRequest, errResp)
	}

	userID := c.Get("currentUser").(entities.User)

	createdCart, err := u.OrderService.ServiceOrderCreateByProduct(userID.Id, createInput)
	if err != nil {
		fmt.Println("create", err)
		errResp := helper.ResponsesFormat("Failed to create Cart", http.StatusBadRequest, nil)
		return c.JSON(http.StatusBadRequest, errResp)
	}

	resp := helper.ResponsesFormat("Successful Operation", http.StatusOK, createdCart)
	return c.JSON(http.StatusOK, resp)

}
