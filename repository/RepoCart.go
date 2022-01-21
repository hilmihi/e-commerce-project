package repository

import (
	"database/sql"
	"fmt"
	"sirclo/api/entities"
)

type RepositoryCart interface {
	GetCarts() ([]entities.Cart, error)
	CreateCart(Cart entities.Cart) (entities.Cart, error)
	GetCart(id int) (entities.Cart, error)
	UpdateCart(Id_Cart int, Cart entities.Cart) (entities.Cart, error)
	DeleteCart(Cart entities.Cart) (entities.Cart, error)
}

type Repository_Cart struct {
	db *sql.DB
}

func NewRepositoryCart(db *sql.DB) *Repository_Cart {
	return &Repository_Cart{db}
}

//get Carts
func (r *Repository_Cart) GetCarts() ([]entities.Cart, error) {
	var Carts []entities.Cart
	results, err := r.db.Query("select id, id_user, id_product, quantity, sub_total from cart_items WHERE deleted_date IS NULL")
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

//get Cart
func (r *Repository_Cart) GetCart(id int) (entities.Cart, error) {
	var Cart entities.Cart

	row := r.db.QueryRow("SELECT id, id_user, id_product, quantity, sub_total FROM cart_items WHERE id = ? AND  deleted_date IS NULL", id)

	err := row.Scan(&Cart.Id, &Cart.Id_user, &Cart.Id_product, &Cart.Quantity, &Cart.Sub_total)
	if err != nil {
		return Cart, err
	}

	return Cart, nil
}

//create Cart
func (r *Repository_Cart) CreateCart(Cart entities.Cart) (entities.Cart, error) {
	query := `INSERT INTO cart_items (id_user, id_product, quantity, sub_total, created_date, updated_date) VALUES (?, ?, ?, ?, now(), now())`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return Cart, err
	}

	defer statement.Close()

	_, err = statement.Exec(Cart.Id_user, Cart.Id_product, Cart.Quantity, Cart.Sub_total)
	if err != nil {
		return Cart, err
	}

	return Cart, nil
}

//update Cart
func (r *Repository_Cart) UpdateCart(Id_Cart int, Cart entities.Cart) (entities.Cart, error) {
	fmt.Println(Id_Cart)
	fmt.Println(Cart)
	query := `UPDATE cart_items SET quantity = ?, updated_date = now() WHERE id = ? AND id_product = ?`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return Cart, err
	}

	defer statement.Close()

	_, err = statement.Exec(Cart.Quantity, Id_Cart, Cart.Id_product)
	if err != nil {
		return Cart, err
	}

	return Cart, nil
}

//delete Cart
func (r *Repository_Cart) DeleteCart(Cart entities.Cart) (entities.Cart, error) {
	query := `UPDATE cart_items SET deleted_date = now() WHERE id = ?`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return Cart, err
	}

	defer statement.Close()

	_, err = statement.Exec(Cart.Id)
	if err != nil {
		return Cart, err
	}

	return Cart, nil
}
