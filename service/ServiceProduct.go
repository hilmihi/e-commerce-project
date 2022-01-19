package service

import (
	"errors"
	"fmt"
	"sirclo/api/entities"
	"sirclo/api/helper"
	"sirclo/api/repository"
)

type ServiceProduct interface {
	ServiceProductsGet() ([]entities.Product, error)
	ServiceProductGet(id int) (entities.Product, error)
	ServiceProductCreate(input helper.RequestProductCreate) (entities.Product, error)
	ServiceProductUpdate(id int, input helper.RequestProductUpdate) (entities.Product, error)
	ServiceProductDelete(id, userID int) (entities.Product, error)
}

type serviceProduct struct {
	repository1 repository.RepositoryProduct
}

func NewProductService(repository1 repository.RepositoryProduct) *serviceProduct {
	return &serviceProduct{repository1}
}

func (su *serviceProduct) ServiceProductsGet() ([]entities.Product, error) {
	Products, err := su.repository1.GetProducts()
	if err != nil {
		return Products, err
	}
	return Products, nil
}

func (s *serviceProduct) ServiceProductGet(id int) (entities.Product, error) {
	Product, err := s.repository1.GetProduct(id)
	if err != nil {
		return Product, err
	}
	return Product, nil
}

func (s *serviceProduct) ServiceProductCreate(input helper.RequestProductCreate) (entities.Product, error) {
	Product := entities.Product{}
	Product.UserID.Id = input.UserID.Id
	Product.Name = input.Name
	Product.Description = input.Description
	Product.Price = input.Price

	createProduct, err := s.repository1.CreateProduct(Product)
	if err != nil {
		return createProduct, err
	}
	return createProduct, nil
}

func (s *serviceProduct) ServiceProductUpdate(id int, input helper.RequestProductUpdate) (entities.Product, error) {
	Product, err := s.repository1.GetProduct(id)
	if err != nil {
		return Product, err
	}
	Product.UserID = input.UserID
	Product.Name = input.Name
	Product.Description = input.Description
	Product.Price = input.Price

	if Product.UserID != input.UserID {
		return Product, errors.New("Unauthorized")
	}

	updateProduct, err := s.repository1.UpdateProduct(Product)
	if err != nil {
		return updateProduct, err
	}
	return updateProduct, nil
}

func (s *serviceProduct) ServiceProductDelete(id, userID int) (entities.Product, error) {
	ProductID, err := s.ServiceProductGet(id)
	if err != nil {
		return ProductID, err
	}

	if userID != ProductID.UserID.Id {
		fmt.Printf("uid: %d & puid: %d", userID, ProductID.UserID.Id)
		return ProductID, errors.New("unauthorized")
	}

	deleteProduct, err := s.repository1.DeleteProduct(ProductID)
	if err != nil {
		return deleteProduct, err
	} else {
		return deleteProduct, nil
	}
}
