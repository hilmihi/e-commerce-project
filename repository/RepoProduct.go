package repository

import (
	"database/sql"
	"fmt"
	"sirclo/api/entities"
	"sirclo/api/helper"
)

type RepositoryProduct interface {
	GetProducts() ([]entities.Product, error)
	CreateProduct(Product entities.Product) (entities.Product, error)
	GetProduct(id int) (helper.ResponseProduct, error)
	UpdateProduct(Id_product int, Product entities.Product) (entities.Product, error)
	DeleteProduct(int) error
}

type Repository_Product struct {
	db *sql.DB
}

func NewRepositoryProduct(db *sql.DB) *Repository_Product {
	return &Repository_Product{db}
}

//get Products
func (r *Repository_Product) GetProducts() ([]entities.Product, error) {
	var Products []entities.Product
	results, err := r.db.Query("select id, id_user, id_category, name, description, price, quantity, photo from products where deleted_date IS NULL")
	if err != nil {
		return nil, err
	}

	defer results.Close()

	for results.Next() {
		var Product entities.Product

		err = results.Scan(&Product.Id, &Product.Id_user, &Product.Id_category, &Product.Name, &Product.Description, &Product.Price, &Product.Quantity, &Product.Photo)
		if err != nil {
			return nil, err
		}

		Products = append(Products, Product)
	}
	return Products, nil
}

//get Product
func (r *Repository_Product) GetProduct(id int) (helper.ResponseProduct, error) {
	var Product helper.ResponseProduct

	row := r.db.QueryRow(`SELECT p.id, p.id_user, p.id_category, c.description as category, p.name, p.description, p.price, p.quantity, p.photo,
								u.id as id_user, u.name, u.email
							FROM products p
							JOIN users u ON p.id_user = u.id
							JOIN category_product c ON c.id = p.id_category
							WHERE p.id = ? AND p.deleted_date IS NULL `, id)

	err := row.Scan(&Product.Id, &Product.Id_user, &Product.Id_category, &Product.Category, &Product.Name, &Product.Description,
		&Product.Price, &Product.Quantity, &Product.Photo, &Product.User.Id, &Product.User.Name, &Product.User.Email)

	if err != nil {
		return Product, err
	}

	return Product, nil
}

//create Product
func (r *Repository_Product) CreateProduct(Product entities.Product) (entities.Product, error) {
	query := `INSERT INTO products (id_user, id_category, name, description, price, quantity, photo, created_date, updated_date) VALUES (?, ?, ?, ?, ?, ?, ?, now(), now())`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return Product, err
	}

	defer statement.Close()

	_, err = statement.Exec(Product.Id_user, Product.Id_category, Product.Name, Product.Description, Product.Price, Product.Quantity, Product.Photo)
	if err != nil {
		return Product, err
	}

	return Product, nil
}

//update Product
func (r *Repository_Product) UpdateProduct(Id_product int, Product entities.Product) (entities.Product, error) {
	fmt.Println(Id_product)
	fmt.Println(Product)
	query := `UPDATE products SET name = ?, description = ?, price = ?, quantity = ?, photo = ?, updated_date = now() WHERE id = ? AND id_user = ?`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return Product, err
	}

	defer statement.Close()

	_, err = statement.Exec(Product.Name, Product.Description, Product.Price, Product.Quantity, Product.Photo, Id_product, Product.Id_user)
	if err != nil {
		return Product, err
	}

	return Product, nil
}

//delete Product
func (r *Repository_Product) DeleteProduct(Id_product int) error {
	query := `UPDATE products SET deleted_date = now() WHERE id = ?`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(Id_product)
	if err != nil {
		return err
	}

	return nil
}
