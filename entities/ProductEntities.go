package entities

type Product struct {
	Id          int
	UserID      User
	Name        string
	Description string
	Price       float64
}
