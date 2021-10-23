package db

import (
	"ecommerce-service/pkg/contract"
	"fmt"
)

type ItemRepo interface {
	FindItems() ([]*contract.Item, error)
	CreateItem(item *contract.Item) error
}

func (s *ShopDb) FindItems() ([]*contract.Item, error) {
	var items []*contract.Item
	result := s.db.Find(&items)
	if result.Error != nil {
		return nil, fmt.Errorf("error occurred while fetching items: %s", result.Error)
	}
	return items, nil
}

func (s *ShopDb) CreateItem(item *contract.Item) error {
	result := s.db.Create(&item)
	return result.Error
}
