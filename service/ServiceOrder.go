package service

import (
	"fmt"
	"sirclo/api/entities"
	"sirclo/api/helper"
	"sirclo/api/repository"
	"time"
)

type ServiceOrder interface {
	ServiceOrderCreateByCart(int, helper.RequestOrderCart) (entities.Transaction, error)
	ServiceOrderCreateByProduct(int, helper.RequestOrderProduct) (entities.Transaction, error)
	ServiceOrdersGet(int) ([]helper.ResponseGetOrder, error)
	ServiceOrdersGetByID(int, int) ([]helper.ResponseGetOrderByID, error)
}

type serviceOrder struct {
	repository1 repository.RepositoryOrder
}

func NewOrderService(repository1 repository.RepositoryOrder) *serviceOrder {
	return &serviceOrder{repository1}
}

func (s *serviceOrder) ServiceOrdersGet(id_user int) ([]helper.ResponseGetOrder, error) {
	Orders, err := s.repository1.GetOrders(id_user)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(Orders)

	return Orders, nil
}

func (s *serviceOrder) ServiceOrdersGetByID(id_user int, id_transaction_detail int) ([]helper.ResponseGetOrderByID, error) {
	Orders, err := s.repository1.GetOrdersByID(id_user, id_transaction_detail)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return Orders, nil
}

func (s *serviceOrder) ServiceOrderCreateByCart(id_user int, input helper.RequestOrderCart) (entities.Transaction, error) {
	listCarts, err := s.repository1.GetCartsIn(input.Id_cart)
	if err != nil {
		fmt.Println(err)
		return entities.Transaction{}, err
	}

	var transaction entities.Transaction
	var transaction_detail []entities.TransactionDetail
	for _, v := range listCarts {
		transaction.Date = time.Now().Format("2006-01-02")
		transaction.Total_price += float64(v.Sub_total)

		var trans entities.TransactionDetail
		trans.Id_product = v.Id_product
		trans.Quantity = v.Quantity
		trans.Sub_total = float64(v.Sub_total)
		transaction_detail = append(transaction_detail, trans)
	}
	res, err := s.repository1.CreateOrder(id_user, input.Address, input.CreditCard, transaction, transaction_detail)
	return res, nil
}

func (s *serviceOrder) ServiceOrderCreateByProduct(id_user int, input helper.RequestOrderProduct) (entities.Transaction, error) {
	Product, err := s.repository1.GetProduct(input.Id_product)
	if err != nil {
		fmt.Println(err)
		return entities.Transaction{}, err
	}

	var transaction entities.Transaction
	var transaction_detail []entities.TransactionDetail
	{
		transaction.Date = time.Now().Format("2006-01-02")
		transaction.Total_price += Product.Price * float64(input.Quantity)

		var trans entities.TransactionDetail
		trans.Id_product = Product.Id
		trans.Quantity = input.Quantity
		trans.Sub_total = Product.Price * float64(input.Quantity)
		transaction_detail = append(transaction_detail, trans)
	}
	res, err := s.repository1.CreateOrder(id_user, input.Address, input.CreditCard, transaction, transaction_detail)
	return res, nil
}
