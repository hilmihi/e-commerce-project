package helper

import (
	"sirclo/api/entities"
)

type AuthFormat struct {
	Token string `json:"token"`
}

func FormatAuth(user entities.User, token string) AuthFormat {
	formatter := AuthFormat{
		Token: token,
	}
	return formatter
}

type UserFormatter struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func FormatUser(user entities.User) UserFormatter {
	formatter := UserFormatter{
		ID:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}
	return formatter
}

type ProductFormatter struct {
	ID          int     `json:"id"`
	UserID      int     `json:"user_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func FormatProduct(product entities.Product) ProductFormatter {
	formatter := ProductFormatter{
		ID:          product.Id,
		UserID:      product.UserID.Id,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}
	return formatter
}
