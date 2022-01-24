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
	ServiceProductGet(id int) (helper.ResponseProduct, error)
	ServiceProductCreate(input entities.Product) (entities.Product, error)
	ServiceProductUpdate(id int, input entities.Product) (entities.Product, error)
	ServiceProductDelete(id, userID int) error
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

func (s *serviceProduct) ServiceProductGet(id int) (helper.ResponseProduct, error) {
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

func (s *serviceProduct) ServiceProductDelete(id, userID int) error {
	ProductID, err := s.ServiceProductGet(id)
	if err != nil {
		return err
	}

	if userID != ProductID.Id_user {
		fmt.Printf("uid: %d & puid: %d", userID, ProductID.Id_user)
		return errors.New("unauthorized")
	}

	err = s.repository1.DeleteProduct(id)
	if err != nil {
		return err
	} else {
		return nil
	}
}
