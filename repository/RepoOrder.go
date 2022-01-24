package repository

import (
	"database/sql"
	"fmt"
	"sirclo/api/entities"
	"sirclo/api/helper"
	"strings"
)

type RepositoryOrder interface {
	GetOrders(int) ([]helper.ResponseGetOrder, error)
	GetOrdersByID(int, int) ([]helper.ResponseGetOrderByID, error)
	CreateOrder(int, entities.Address, entities.CreditCard, entities.Transaction, []entities.TransactionDetail) (entities.Transaction, error)
	GetCartsIn([]int) ([]entities.Cart, error)
	GetProduct(id int) (entities.Product, error)
}

type Repository_Order struct {
	db *sql.DB
}

func NewRepositoryOrder(db *sql.DB) *Repository_Order {
	return &Repository_Order{db}
}

//get Orders
func (r *Repository_Order) GetOrders(id_user int) ([]helper.ResponseGetOrder, error) {
	var Orders []helper.ResponseGetOrder

	results, err := r.db.Query(`
		SELECT td.id, t.date, td.quantity, td.sub_total, t.id_user, p.id as id_product,
		p.name as product_name, p.price, p.description, p.photo, ts.description as status,
		p.id_user as id_seller, p.id_category
		FROM transaction t
		JOIN transaction_detail td ON td.id_transaction = t.id
		JOIN products p ON td.id_product = p.id
		JOIN transaction_status ts ON ts.id = td.id_status
		WHERE t.id_user = ? AND t.deleted_date IS NULL
	`, id_user)
	if err != nil {
		return nil, err
	}

	defer results.Close()

	for results.Next() {
		var Order helper.ResponseGetOrder

		err = results.Scan(&Order.Id, &Order.Date, &Order.Quantity, &Order.Sub_total, &Order.Id_user, &Order.Product.Id,
			&Order.Product.Name, &Order.Product.Price, &Order.Product.Description, &Order.Product.Photo,
			&Order.Status, &Order.Product.Id_user, &Order.Product.Id_category)

		if err != nil {
			return nil, err
		}

		Orders = append(Orders, Order)
	}
	return Orders, nil
}

//get Orders
func (r *Repository_Order) GetOrdersByID(id_user int, id_transaction_detail int) ([]helper.ResponseGetOrderByID, error) {
	var Orders []helper.ResponseGetOrderByID

	results, err := r.db.Query(`
		SELECT td.id, t.date, td.quantity, td.sub_total, t.id_user, p.id as id_product,
		p.name as product_name, p.price, p.description, p.photo, ts.description as status,
		p.id_user as id_seller, p.id_category, a.state, a.street, a.zip
		FROM transaction t
		JOIN transaction_detail td ON td.id_transaction = t.id
		JOIN products p ON td.id_product = p.id
		JOIN transaction_status ts ON ts.id = td.id_status
		JOIN address a ON t.id = a.id_transaction
		WHERE td.id = ? AND t.id_user = ? AND t.deleted_date IS NULL
	`, id_transaction_detail, id_user)
	if err != nil {
		return nil, err
	}

	defer results.Close()

	for results.Next() {
		var Order helper.ResponseGetOrderByID

		err = results.Scan(&Order.Id, &Order.Date, &Order.Quantity, &Order.Sub_total, &Order.Id_user, &Order.Product.Id,
			&Order.Product.Name, &Order.Product.Price, &Order.Product.Description, &Order.Product.Photo,
			&Order.Status, &Order.Product.Id_user, &Order.Product.Id_category, &Order.Address.State,
			&Order.Address.Street, &Order.Address.Zip)
		if err != nil {
			return nil, err
		}

		Orders = append(Orders, Order)
	}
	return Orders, nil
}

//create Order
func (r *Repository_Order) CreateOrder(id_user int, address entities.Address, card entities.CreditCard, transaction entities.Transaction, transaction_detail []entities.TransactionDetail) (entities.Transaction, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return entities.Transaction{}, nil
	}
	//insert credit card
	var id_credit_card int64
	{
		stmt, err := tx.Prepare(`INSERT INTO credit_card (name, type, number, cvv, month, year, created_date, updated_date)
                     VALUES(?,?,?,?,?,?,now(),now());`)
		if err != nil {
			fmt.Println("prepare credit")
			fmt.Println(err)
			tx.Rollback()
			return entities.Transaction{}, err
		}
		defer stmt.Close()

		res, err := stmt.Exec(card.Name, card.Type, card.Number, card.CVV, card.Month, card.Year)
		if err != nil {
			fmt.Println("exec credit")
			fmt.Println(err)
			tx.Rollback() // return an error too, we may want to wrap them
			return entities.Transaction{}, err
		}

		id_credit_card, err = res.LastInsertId()
	}

	//insert transaction
	var id_transaction int64
	{
		stmt, err := tx.Prepare(`INSERT INTO transaction (id_user, id_credit_card, date, total_price, created_date, updated_date)
                     VALUES(?,?,?,?,now(),now());`)
		if err != nil {
			fmt.Println("prepare transaction")
			fmt.Println(err)
			tx.Rollback()
			return entities.Transaction{}, err
		}
		defer stmt.Close()

		res, err := stmt.Exec(id_user, id_credit_card, transaction.Date, transaction.Total_price)
		if err != nil {
			fmt.Println("exec transaction")
			fmt.Println(err)
			tx.Rollback() // return an error too, we may want to wrap them
			return entities.Transaction{}, err
		}

		id_transaction, err = res.LastInsertId()
	}

	//insert address
	{
		stmt, err := tx.Prepare(`INSERT INTO address (id_transaction, state, street, zip, created_date, updated_date)
                     VALUES(?,?,?,?,now(),now());`)
		if err != nil {
			fmt.Println("prepare address")
			fmt.Println(err)
			tx.Rollback()
			return entities.Transaction{}, err
		}
		defer stmt.Close()

		_, err = stmt.Exec(id_transaction, address.State, address.Street, address.Zip)
		if err != nil {
			fmt.Println("exec address")
			fmt.Println(err)
			tx.Rollback() // return an error too, we may want to wrap them
			return entities.Transaction{}, err
		}
	}

	//insert transaction detail
	{
		stmt, err := tx.Prepare(`INSERT INTO transaction_detail (id_transaction, id_product, id_status, quantity, sub_total, created_date, updated_date)
                     VALUES(?,?,?,?,?,now(),now())` + strings.Repeat(",(?,?,?,?,?,now(),now())", len(transaction_detail)-1) + `;`)
		if err != nil {
			fmt.Println("prepare transaction detail")
			fmt.Println(err)
			tx.Rollback()
			return entities.Transaction{}, err
		}
		defer stmt.Close()

		args := []interface{}{}
		for _, v := range transaction_detail {
			args = append(args, id_transaction)
			args = append(args, v.Id_product)
			args = append(args, 1)
			args = append(args, v.Quantity)
			args = append(args, v.Sub_total)
		}
		_, err = stmt.Exec(args...)
		if err != nil {
			fmt.Println("exec transaction detail")
			fmt.Println(err)
			tx.Rollback() // return an error too, we may want to wrap them
			return entities.Transaction{}, err
		}
	}

	tx.Commit()
	return entities.Transaction{}, nil
}

//get Carts IN
func (r *Repository_Order) GetCartsIn(list_id []int) ([]entities.Cart, error) {
	var Carts []entities.Cart
	args := []interface{}{}
	for _, v := range list_id {
		args = append(args, v)
	}

	results, err := r.db.Query("select id, id_user, id_product, quantity, sub_total from cart_items WHERE deleted_date IS NULL AND id IN (?"+strings.Repeat(",?", len(list_id)-1)+")", args...)
	if err != nil {
		return nil, err
	}

	defer results.Close()

	for results.Next() {
		var Cart entities.Cart

		err = results.Scan(&Cart.Id, &Cart.Id_user, &Cart.Id_product, &Cart.Quantity, &Cart.Sub_total)
		if err != nil {
			return nil, err
		}

		Carts = append(Carts, Cart)
	}
	return Carts, nil
}

//get Product
func (r *Repository_Order) GetProduct(id int) (entities.Product, error) {
	var Product entities.Product

	row := r.db.QueryRow(`SELECT id, id_user, id_category, name, description, price, quantity, photo FROM products WHERE id = ? AND deleted_date IS NULL `, id)

	err := row.Scan(&Product.Id, &Product.Id_user, &Product.Id_category, &Product.Name, &Product.Description, &Product.Price, &Product.Quantity, &Product.Photo)
	if err != nil {
		return Product, err
	}

	return Product, nil
}
