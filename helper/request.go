package helper

import "sirclo/api/entities"

type RequestUserLogin struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type RequestUserCreate struct {
	Name         string `json:"name" form:"name"`
	Email        string `json:"email" form:"email"`
	Password     string `json:"password" form:"password"`
	Birth_date   string `json:"birth_date" form:"birth_date"`
	Phone_number int    `json:"phone_number" form:"phone_number"`
	Photo        string `json:"photo" form:"photo"`
	Gender       string `json:"gender" form:"gender"`
	Address      string `json:"address" form:"address"`
}

type RequestUserUpdate struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type RequestProductCreate struct {
	UserID      entities.User `json:"userID" form:"userID"`
	Name        string        `json:"name" form:"name"`
	Description string        `json:"description" form:"description"`
	Price       float64       `json:"price" form:"price"`
}

type RequestProductUpdate struct {
	UserID      entities.User `json:"userID" form:"userID"`
	Name        string        `json:"name" form:"name"`
	Description string        `json:"description" form:"description"`
	Price       float64       `json:"price" form:"price"`
}
