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
	Id        int              `json:"id" form:"id"`
	Id_user   int              `json:"user_id"`
	Quantity  int              `json:"quantity" form:"quantity"`
	Sub_total float64          `json:"sub_total" form:"sub_total"`
	Date      string           `json:"date" form:"date"`
	Product   entities.Product `json:"product" form:"product"`
	Status    string           `json:"status" form:"status"`
}

type ResponseGetOrderByID struct {
	Id        int              `json:"id" form:"id"`
	Id_user   int              `json:"user_id"`
	Quantity  int              `json:"quantity" form:"quantity"`
	Sub_total float64          `json:"sub_total" form:"sub_total"`
	Date      string           `json:"date" form:"date"`
	Product   ResponseProduct  `json:"product" form:"product"`
	Status    string           `json:"status" form:"status"`
	Address   entities.Address `json:"address" form:"address"`
}

type ResponseProduct struct {
	entities.Product
	Category string       `json:"category" form:"category"`
	User     ResponseUser `json:"user" form:"user"`
}

type ResponseUser struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ResponseCart struct {
	entities.Cart
	Product ResponseProduct2 `json:"product" form:"product"`
}

type ResponseProduct2 struct {
	Id          int     `json:"id" form:"id"`
	Name        string  `json:"name" form:"name"`
	Description string  `json:"description" form:"description"`
	Price       float64 `json:"price" form:"price"`
	Quantity    int     `json:"quantity" form:"quantity"`
}

func FormatUser(user entities.User) UserFormatter {
	formatter := UserFormatter{
		ID:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}
	return formatter
}
