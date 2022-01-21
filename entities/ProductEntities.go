package entities

type Product struct {
	Id          int     `json:"id" form:"id"`
	Id_user     int     `json:"id_user" form:"id_user"`
	Id_category int     `json:"id_category" form:"id_category"`
	Name        string  `json:"name" form:"name"`
	Description string  `json:"description" form:"description"`
	Price       float64 `json:"price" form:"price"`
	Quantity    int     `json:"quantity" form:"quantity"`
	Photo       string  `json:"photo" form:"photo"`
}
