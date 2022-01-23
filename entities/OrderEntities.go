package entities

type Transaction struct {
	Id             int     `json:"id" form:"id"`
	Id_user        int     `json:"id_user" form:"id_user"`
	Id_credit_card int     `json:"id_credit_card" form:"id_credit_card"`
	Date           string  `json:"date" form:"date"`
	Total_price    float64 `json:"total_price" form:"total_price"`
}

type TransactionDetail struct {
	Id             int     `json:"id" form:"id"`
	Id_transaction int     `json:"id_transaction" form:"id_transaction"`
	Id_product     int     `json:"id_product" form:"id_product"`
	Id_status      int     `json:"id_status" form:"id_status"`
	Quantity       int     `json:"quantity" form:"quantity"`
	Sub_total      float64 `json:"sub_total" form:"sub_total"`
}

type Address struct {
	Id             int    `json:"id" form:"id"`
	Id_transaction int    `json:"id_transaction" form:"id_transaction"`
	State          string `json:"state" form:"state"`
	Street         string `json:"street" form:"street"`
	Zip            int    `json:"zip" form:"zip"`
}

type CreditCard struct {
	Id     int    `json:"id" form:"id"`
	Name   string `json:"name" form:"name"`
	Type   string `json:"type" form:"type"`
	Number int    `json:"number" form:"number"`
	CVV    int    `json:"cvv" form:"cvv"`
	Month  int    `json:"month" form:"month"`
	Year   int    `json:"year" form:"year"`
}
