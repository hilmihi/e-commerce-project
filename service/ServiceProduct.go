package service

import (
	"errors"
	"fmt"
	"sirclo/api/entities"
	"sirclo/api/repository"
)

type ServiceProduct interface {
	ServiceProductsGet() ([]entities.Product, error)
	ServiceProductGet(id int) (entities.Product, error)
	ServiceProductCreate(input entities.Product) (entities.Product, error)
	ServiceProductUpdate(id int, input entities.Product) (entities.Product, error)
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

func (s *serviceProduct) ServiceProductCreate(input entities.Product) (entities.Product, error) {
	createProduct, err := s.repository1.CreateProduct(input)
	if err != nil {
		return createProduct, err
	}
	return createProduct, nil
}

func (s *serviceProduct) ServiceProductUpdate(id int, input entities.Product) (entities.Product, error) {
	updateProduct, err := s.repository1.UpdateProduct(id, input)
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

	if userID != ProductID.Id_user {
		fmt.Printf("uid: %d & puid: %d", userID, ProductID.Id_user)
		return ProductID, errors.New("unauthorized")
	}

	deleteProduct, err := s.repository1.DeleteProduct(ProductID)
	if err != nil {
		return deleteProduct, err
	} else {
		return deleteProduct, nil
	}
}
