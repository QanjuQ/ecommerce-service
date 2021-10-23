package db

import (
	"ecommerce-service/pkg/contract"
	"fmt"
)

type OrderRepo interface {
	FindOrders(id uint) ([]*contract.Order, error)
	CreateOrder(order *contract.Order) error
}

func (s *ShopDb) FindOrders(id uint) ([]*contract.Order, error) {
	var orders []*contract.Order
	result := s.db.Find(&orders)
	if result.Error != nil {
		return nil, fmt.Errorf("error occurred while fetching orders for username: %d %s", id, result.Error)
	}
	return orders, nil
}

func (s *ShopDb) CreateOrder(order *contract.Order) error {
	result := s.db.Create(&order)
	return result.Error
}
