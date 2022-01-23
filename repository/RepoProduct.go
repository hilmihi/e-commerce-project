package repository

import (
	"database/sql"
	"fmt"
	"sirclo/api/entities"
)

type RepositoryProduct interface {
	GetProducts() ([]entities.Product, error)
	CreateProduct(Product entities.Product) (entities.Product, error)
	GetProduct(id int) (entities.Product, error)
	UpdateProduct(Id_product int, Product entities.Product) (entities.Product, error)
	DeleteProduct(Product entities.Product) (entities.Product, error)
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
func (r *Repository_Product) GetProduct(id int) (entities.Product, error) {
	var Product entities.Product

	row := r.db.QueryRow(`SELECT id, id_user, id_category, name, description, price, quantity, photo FROM products WHERE id = ? AND deleted_date IS NULL `, id)

	err := row.Scan(&Product.Id, &Product.Id_user, &Product.Id_category, &Product.Name, &Product.Description, &Product.Price, &Product.Quantity, &Product.Photo)
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
func (r *Repository_Product) DeleteProduct(Product entities.Product) (entities.Product, error) {
	query := `UPDATE products SET deleted_date = now() WHERE id = ?`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return Product, err
	}

	defer statement.Close()

	_, err = statement.Exec(Product.Id)
	if err != nil {
		return Product, err
	}

	return Product, nil
}
