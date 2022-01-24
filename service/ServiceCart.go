package service

import (
	"errors"
	"fmt"
	"sirclo/api/entities"
	"sirclo/api/helper"
	"sirclo/api/repository"
)

type ServiceCart interface {
	ServiceCartsGet(id_user int) ([]helper.ResponseCart, error)
	ServiceCartGet(id int) (entities.Cart, error)
	ServiceCartCreate(input entities.Cart) (entities.Cart, error)
	ServiceCartUpdate(id int, input entities.Cart) (entities.Cart, error)
	ServiceCartDelete(id, userID int) (entities.Cart, error)
}

type serviceCart struct {
	repository1 repository.RepositoryCart
}

func NewCartService(repository1 repository.RepositoryCart) *serviceCart {
	return &serviceCart{repository1}
}

func (su *serviceCart) ServiceCartsGet(id_user int) ([]helper.ResponseCart, error) {
	Carts, err := su.repository1.GetCarts(id_user)
	if err != nil {
		return Carts, err
	}
	return Carts, nil
}

func (s *serviceCart) ServiceCartGet(id int) (entities.Cart, error) {
	Cart, err := s.repository1.GetCart(id)
	if err != nil {
		return Cart, err
	}
	return Cart, nil
}

func (s *serviceCart) ServiceCartCreate(input entities.Cart) (entities.Cart, error) {
	createCart, err := s.repository1.CreateCart(input)
	if err != nil {
		return createCart, err
	}
	return createCart, nil
}

func (s *serviceCart) ServiceCartUpdate(id int, input entities.Cart) (entities.Cart, error) {
	updateCart, err := s.repository1.UpdateCart(id, input)
	if err != nil {
		return updateCart, err
	}
	return updateCart, nil
}

func (s *serviceCart) ServiceCartDelete(id, userID int) (entities.Cart, error) {
	CartID, err := s.ServiceCartGet(id)
	if err != nil {
		return CartID, err
	}

	if userID != CartID.Id_user {
		fmt.Printf("uid: %d & puid: %d", userID, CartID.Id_user)
		return CartID, errors.New("unauthorized")
	}

	deleteCart, err := s.repository1.DeleteCart(CartID)
	if err != nil {
		return deleteCart, err
	} else {
		return deleteCart, nil
	}
}
