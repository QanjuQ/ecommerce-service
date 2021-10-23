package service

import (
	"ecommerce-service/pkg/contract"
	"fmt"
)

type ItemService interface {
	CreateItem(request *contract.CreateItemRequest) error
	FetchItems() ([]*contract.Item, error)
}

func (s *ShopService) CreateItem(request *contract.CreateItemRequest) error {
	item := contract.Item{
		Name: request.Name,
	}
	err := s.Db.CreateItem(&item)
	if err != nil {
		return err
	}
	return nil
}

func (s *ShopService) FetchItems() ([]*contract.Item, error) {
	users, err := s.Db.FindItems()
	if err != nil {
		return []*contract.Item{}, fmt.Errorf("eror while fetching users: %s", err)
	}
	return users, err
}
