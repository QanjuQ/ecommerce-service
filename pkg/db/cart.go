package db

import (
	"ecommerce-service/pkg/contract"
	"fmt"
)

type CartRepo interface {
	FindCarts(id uint) ([]*contract.Cart, error)
	AddItem(cartItem *contract.CartItem) error
	Update(cart *contract.Cart) error
}

func (s *ShopDb) FindCarts(id uint) ([]*contract.Cart, error) {
	var carts []*contract.Cart
	result := s.db.Find(&carts)
	if result.Error != nil {
		return nil, fmt.Errorf("error occurred while fetching carts for user: %s", result.Error)
	}
	return carts, nil
}

func (s *ShopDb) AddItem(cartItem *contract.CartItem) error {
	result := s.db.Create(&cartItem)
	if result.Error != nil {
		return fmt.Errorf("add item to cart: %d error: %s", cartItem.CartId, result.Error)
	}
	return result.Error
}

func (s *ShopDb) Update(cart *contract.Cart) error {
	result := s.db.Model(&cart.ID).Updates(&cart)
	if result.Error != nil {
		return fmt.Errorf("error purchasing the cart items: %d error: %s", cart.ID, result.Error)
	}
	return result.Error
}
