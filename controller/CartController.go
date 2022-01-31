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

type CartHF struct {
	CartService service.ServiceCart
}

func NewCartController(CartService service.ServiceCart) *CartHF {
	return &CartHF{CartService}
}

//Cart get all
func (u *CartHF) GetCartsController(c echo.Context) error {
	id_user := c.Get("currentUser").(entities.User)
	Carts, err := u.CartService.ServiceCartsGet(id_user.Id)
	if err != nil {
		fmt.Println(err)
		response := helper.ResponsesFormat("Failed to fetch Cart data", http.StatusOK, err)
		return c.JSON(http.StatusOK, response)
	}

	return c.JSON(http.StatusOK, Carts)
}

//Cart get by id
func (u *CartHF) GetCartController(c echo.Context) error {
	CartId, errId := strconv.Atoi(c.Param("id"))
	if errId != nil {
		errResp := helper.ResponsesFormat("Failed to convert id", http.StatusInternalServerError, nil)
		return c.JSON(http.StatusInternalServerError, errResp)
	}

	Cart, err := u.CartService.ServiceCartGet(CartId)
	if err != nil {
		fmt.Println(err)
		errResp := helper.ResponsesFormat("Failed to get Cart by id", http.StatusBadRequest, nil)
		return c.JSON(http.StatusBadRequest, errResp)
	}

	return c.JSON(http.StatusOK, Cart)
}

// Cart update
func (u *CartHF) UpdateCartController(c echo.Context) error {
	CartId, errId := strconv.Atoi(c.Param("id"))
	if errId != nil {
		errResp := helper.ResponsesFormat("failed to convert id", http.StatusBadRequest, nil)
		return c.JSON(http.StatusInternalServerError, errResp)

	}

	var updateInput entities.Cart
	if err := c.Bind(&updateInput); err != nil {
		fmt.Println("bind", err)
		errResp := helper.ResponsesFormat("Failed to update data", http.StatusBadRequest, nil)
		return c.JSON(http.StatusBadRequest, errResp)

	}

	userID := c.Get("currentUser").(entities.User)
	updateInput.Id_user = userID.Id

	update, err := u.CartService.ServiceCartUpdate(CartId, updateInput)
	if err != nil {
		fmt.Println("update", err)
		errResp := helper.ResponsesFormat("unauthorized", http.StatusUnauthorized, nil)
		return c.JSON(http.StatusUnauthorized, errResp)

	}

	resp := helper.ResponsesFormat("Successful Operation", http.StatusOK, update)
	return c.JSON(http.StatusOK, resp)

}

//create Cart

func (u *CartHF) CreateCartController(c echo.Context) error {
	var createInput entities.Cart
	if err := c.Bind(&createInput); err != nil {
		fmt.Println("CI", err)
		errResp := helper.ResponsesFormat("Failed to Create Cart", http.StatusBadRequest, nil)
		return c.JSON(http.StatusBadRequest, errResp)
	}

	userID := c.Get("currentUser").(entities.User)
	createInput.Id_user = userID.Id
	createdCart, err := u.CartService.ServiceCartCreate(createInput)
	if err != nil {
		fmt.Println("create", err)
		errResp := helper.ResponsesFormat("Failed to create Cart", http.StatusBadRequest, nil)
		return c.JSON(http.StatusBadRequest, errResp)
	}

	resp := helper.ResponsesFormat("Successful Operation", http.StatusOK, createdCart)
	return c.JSON(http.StatusOK, resp)

}

//delete Cart
func (u *CartHF) DeleteCartController(c echo.Context) error {
	CartId, errId := strconv.Atoi(c.Param("id"))
	if errId != nil {
		// fmt.Println("id", errId)
		errResp := helper.ResponsesFormat("Failed to delete Cart", http.StatusInternalServerError, nil)
		return c.JSON(http.StatusBadRequest, errResp)
	}

	userID := c.Get("currentUser").(entities.User)

	deleteCart, err := u.CartService.ServiceCartDelete(CartId, userID.Id)
	// fmt.Printf("PI %d = UI %d", CartId, userID.Id)
	if err != nil {
		fmt.Println("del", err)
		errResp := helper.ResponsesFormat("unauthorized", http.StatusUnauthorized, nil)
		return c.JSON(http.StatusUnauthorized, errResp)
	}

	resp := helper.ResponsesFormat("Successful Operation", http.StatusOK, deleteCart)
	return c.JSON(http.StatusOK, resp)
}
