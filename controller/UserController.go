package controller

import (
	"fmt"
	"net/http"
	addmiddleware "sirclo/api/addMiddleware"
	"sirclo/api/entities"
	"sirclo/api/helper"
	"sirclo/api/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHF struct {
	authService addmiddleware.JWTService
	userService service.ServiceUser
}

func NewUserController(authService addmiddleware.JWTService, userService service.ServiceUser) *UserHF {
	return &UserHF{authService, userService}
}

func (h *UserHF) AuthUserController(c echo.Context) error {
	var input helper.RequestUserLogin
	if err := c.Bind(&input); err != nil {
		response := helper.ResponsesFormat("Failed to Login as User", http.StatusUnprocessableEntity, nil)
		fmt.Println("bind: ", err)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	loginUser, err := h.userService.ServiceUserLogin(input)
	if err != nil {
		response := helper.ResponsesFormat("Failed to Login as User", http.StatusBadRequest, nil)
		fmt.Println("login: ", err)
		return c.JSON(http.StatusBadRequest, response)
	}

	token, err := h.authService.GenerateToken(loginUser.Id)
	if err != nil {
		response := helper.ResponsesFormat("Failed to Login as User", http.StatusBadRequest, nil)
		fmt.Println("GT: ", err)
		return c.JSON(http.StatusBadRequest, response)
	}

	formatter := helper.FormatAuth(loginUser, token)
	response := helper.ResponsesAuth("Success to Login as User", http.StatusOK, formatter)
	return c.JSON(http.StatusOK, response)
}

//user get all
func (u *UserHF) GetUsersController(c echo.Context) error {
	users, err := u.userService.ServiceUsersGet()
	if err != nil {
		response := helper.ResponsesFormat("Failed to get user data ", http.StatusOK, nil)
		return c.JSON(http.StatusOK, response)
	}

	return c.JSON(http.StatusOK, users)
}

//user get by id
func (u *UserHF) GetUserController(c echo.Context) error {
	userId, errId := strconv.Atoi(c.Param("id"))
	if errId != nil {
		errResp := helper.ResponsesFormat("Failed to convert id", http.StatusInternalServerError, nil)
		return c.JSON(http.StatusInternalServerError, errResp)
	}

	user, err := u.userService.ServiceUserGet(userId)
	if err != nil {
		fmt.Println(err)
		errResp := helper.ResponsesFormat("Failed to get user by id", http.StatusBadRequest, nil)
		return c.JSON(http.StatusBadRequest, errResp)
	}

	return c.JSON(http.StatusOK, user)
}

//user get by id
func (u *UserHF) GetMyUserController(c echo.Context) error {
	userID := c.Get("currentUser").(entities.User)
	user, err := u.userService.ServiceUserGet(userID.Id)
	if err != nil {
		fmt.Println(err)
		errResp := helper.ResponsesFormat("Failed to get user by id", http.StatusBadRequest, nil)
		return c.JSON(http.StatusBadRequest, errResp)
	}

	return c.JSON(http.StatusOK, user)
}

// user update
func (u *UserHF) UpdateUserController(c echo.Context) error {
	userId, errId := strconv.Atoi(c.Param("id"))
	if errId != nil {
		errResp := helper.ResponsesFormat("failed to convert id", http.StatusBadRequest, nil)
		return c.JSON(http.StatusInternalServerError, errResp)

	}

	var updateInput entities.User
	if err := c.Bind(&updateInput); err != nil {
		errResp := helper.ResponsesFormat("Failed to update data", http.StatusBadRequest, nil)
		return c.JSON(http.StatusBadRequest, errResp)

	}

	update, err := u.userService.ServiceUserUpdate(userId, updateInput)
	if err != nil {
		errResp := helper.ResponsesFormat("Failed to update data", http.StatusBadRequest, nil)
		return c.JSON(http.StatusBadRequest, errResp)

	}

	formatter := helper.FormatUser(update)
	resp := helper.ResponsesFormat("Success to update user data", http.StatusOK, formatter)
	return c.JSON(http.StatusOK, resp)

}

//create user

func (u *UserHF) CreateUserController(c echo.Context) error {
	var createInput entities.User
	if err := c.Bind(&createInput); err != nil {
		errResp := helper.ResponsesFormat("Failed to Create user", http.StatusBadRequest, nil)
		return c.JSON(http.StatusBadRequest, errResp)
	}

	createdUser, err := u.userService.ServiceUserCreate(createInput)
	if err != nil {
		errResp := helper.ResponsesFormat("Failed to create user", http.StatusBadRequest, nil)
		return c.JSON(http.StatusBadRequest, errResp)
	}

	resp := helper.ResponsesFormat("Successful Operation", http.StatusOK, createdUser)
	return c.JSON(http.StatusOK, resp)

}

//delete user
func (u *UserHF) DeleteUserController(c echo.Context) error {
	userId, errId := strconv.Atoi(c.Param("id"))
	if errId != nil {
		errResp := helper.ResponsesFormat("Failed to delete user", http.StatusInternalServerError, nil)
		return c.JSON(http.StatusBadRequest, errResp)
	}

	deleteUser, err := u.userService.ServiceUserDelete(userId)
	if err != nil {
		errResp := helper.ResponsesFormat("Failed to delete user", http.StatusBadRequest, nil)
		return c.JSON(http.StatusBadRequest, errResp)
	}

	formatter := helper.FormatUser(deleteUser)
	resp := helper.ResponsesFormat("Success delete user", http.StatusOK, formatter)
	return c.JSON(http.StatusOK, resp)
}
