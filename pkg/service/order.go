package service

import (
	"ecommerce-service/pkg/contract"
	"fmt"
)


type OrderService interface {
	AddItems(items []*contract.AddItemToCartRequest, user *contract.User) error
	Purchase(user *contract.User) error
	FetchOrders(id uint) ([]*contract.Order, error)
	FetchCarts(id uint) ([]*contract.Cart, error)
}

func (s *ShopService) AddItems(items []*contract.AddItemToCartRequest, user *contract.User) error {
	for _, item := range items {
		cartItem := contract.CartItem{
			Cart:     user.Cart,
			CartId:   user.CartId,
			ItemId:   item.Id,
			Quantity: item.Quantity,
		}

		err := s.Db.AddItem(&cartItem)

		if err != nil {
			return fmt.Errorf("error while adding item to cart: %s", user.Username)
		}
	}

	return nil
}

func (s *ShopService) Purchase(user *contract.User) error {
	user.Cart.Purchase()
	err := s.Db.Update(&user.Cart)
	if err != nil {
		return fmt.Errorf("error while adding updating cart for purchase: %s", user.Username)
	}

	order := contract.Order{
		CartId: user.CartId,
		Cart:   user.Cart,
		UserId: user.ID,
		User:   *user,
	}

	err = s.Db.CreateOrder(&order)

	if err != nil {
		return fmt.Errorf("error adding cart as order for user: %s", user.Username)
	}

	return nil
}

func (s *ShopService) FetchOrders(id uint) ([]*contract.Order, error) {
	orders, err := s.Db.FindOrders(id)
	if err != nil {
		return nil, fmt.Errorf("error while fetching orders for user id: %d", id)
	}
	return orders, nil
}

func (s *ShopService) FetchCarts(id uint) ([]*contract.Cart, error) {
	orders, err := s.Db.FindCarts(id)
	if err != nil {
		return nil, fmt.Errorf("error while fetching carts for user id: %d", id)
	}
	return orders, nil
}
