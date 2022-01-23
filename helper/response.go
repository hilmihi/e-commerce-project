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

type ResponseGetOrder struct {
	Id_user int `json:"user_id"`
	entities.TransactionDetail
	Date    string           `json:"date" form:"date"`
	Product entities.Product `json:"product" form:"product"`
	Status  string           `json:"status" form:"status"`
}

func FormatUser(user entities.User) UserFormatter {
	formatter := UserFormatter{
		ID:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}
	return formatter
}
