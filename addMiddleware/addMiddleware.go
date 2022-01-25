package addmiddleware

import (
	"fmt"
	"net/http"
	"sirclo/api/helper"
	"sirclo/api/service"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(authService JWTService, userService service.ServiceUser, next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.ResponsesFormat("Unauthorized", http.StatusUnauthorized, nil)
			return c.JSON(http.StatusUnauthorized, response)
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			fmt.Println("token", err)
			response := helper.ResponsesFormat("Unauthorized", http.StatusUnauthorized, nil)
			return c.JSON(http.StatusUnauthorized, response)
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			fmt.Println("claims", ok)
			response := helper.ResponsesFormat("Unauthorized", http.StatusUnauthorized, nil)
			return c.JSON(http.StatusUnauthorized, response)
		}

		userID := int(claims["id"].(float64))
		user, err := userService.ServiceUserGet(userID)
		if err != nil {
			fmt.Println("userID", err)
			response := helper.ResponsesFormat("Unauthorized", http.StatusUnauthorized, nil)
			return c.JSON(http.StatusUnauthorized, response)
		}
		c.Set("currentUser", user)
		return next(c)
	}
}
