package entities

type User struct {
	Id           int    `json:"id" form:"id"`
	Name         string `json:"name" form:"name"`
	Email        string `json:"email" form:"email"`
	Password     string `json:"password" form:"password"`
	Birth_date   string `json:"birth_date" form:"birth_date"`
	Phone_number string `json:"phone_number" form:"phone_number"`
	Photo        string `json:"photo" form:"photo"`
	Gender       string `json:"gender" form:"gender"`
	Address      string `json:"address" form:"address"`
}
