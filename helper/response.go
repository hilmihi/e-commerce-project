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
	Id        int             `json:"id" form:"id"`
	Id_user   int             `json:"user_id"`
	Quantity  int             `json:"quantity" form:"quantity"`
	Sub_total float64         `json:"sub_total" form:"sub_total"`
	Date      string          `json:"date" form:"date"`
	Product   ResponseProduct `json:"product" form:"product"`
	Status    string          `json:"status" form:"status"`
}

type ResponseProduct struct {
	Id          int     `json:"id" form:"id"`
	Id_user     int     `json:"id_seller" form:"id_user"`
	Id_category int     `json:"id_category" form:"id_category"`
	Name        string  `json:"name" form:"name"`
	Description string  `json:"description" form:"description"`
	Price       float64 `json:"price" form:"price"`
	Photo       string  `json:"photo" form:"photo"`
}

func FormatUser(user entities.User) UserFormatter {
	formatter := UserFormatter{
		ID:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}
	return formatter
}
