package repository

import (
	"database/sql"
	"log"
	"sirclo/api/entities"
)

type RepositoryProduct interface {
	GetProducts() ([]entities.Product, error)
	CreateProduct(Product entities.Product) (entities.Product, error)
	GetProduct(id int) (entities.Product, error)
	UpdateProduct(Product entities.Product) (entities.Product, error)
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
	results, err := r.db.Query("select * from Products")
	if err != nil {
		log.Fatalf("Error")
	}

	defer results.Close()

	for results.Next() {
		var Product entities.Product

		err = results.Scan(&Product.Id, &Product.UserID.Id, &Product.Name, &Product.Description, &Product.Price)
		if err != nil {
			log.Fatalf("Error")
		}

		Products = append(Products, Product)
	}
	return Products, nil
}

//get Product
func (r *Repository_Product) GetProduct(id int) (entities.Product, error) {
	var Product entities.Product

	row := r.db.QueryRow(`SELECT * FROM Products WHERE id = ?`, id)

	err := row.Scan(&Product.Id, &Product.UserID.Id, &Product.Name, &Product.Description, &Product.Price)
	if err != nil {
		return Product, err
	}

	return Product, nil
}

//create Product
func (r *Repository_Product) CreateProduct(Product entities.Product) (entities.Product, error) {
	query := `INSERT INTO products (userID, productName, ProductDescription, Price) VALUES (?, ?, ?, ?)`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return Product, err
	}

	defer statement.Close()

	_, err = statement.Exec(Product.UserID.Id, Product.Name, Product.Description, Product.Price)
	if err != nil {
		return Product, err
	}

	return Product, nil
}

//update Product
func (r *Repository_Product) UpdateProduct(Product entities.Product) (entities.Product, error) {
	query := `UPDATE Products SET ProductName = ?, ProductDescription = ?, Price = ? WHERE id = ?`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return Product, err
	}

	defer statement.Close()

	_, err = statement.Exec(Product.Name, Product.Description, Product.Price, Product.Id)
	if err != nil {
		return Product, err
	}

	return Product, nil
}

//delete Product
func (r *Repository_Product) DeleteProduct(Product entities.Product) (entities.Product, error) {
	query := `DELETE FROM Products WHERE id = ?`

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
