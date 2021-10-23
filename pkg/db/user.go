package db

import (
	"ecommerce-service/pkg/contract"
	"fmt"
)

type UserRepo interface {
	CreateUser(user *contract.User) error
	FindUserWithUsername(username string) (*contract.User, error)
	FindUserWithToken(token string) (*contract.User, error)
	UpdateUser(user *contract.User) error
	FindUsers() ([]*contract.User, error)
}

func (s *ShopDb) CreateUser(user *contract.User) error {
	cart := contract.Cart{}
	result := s.db.Create(&cart)
	if result.Error != nil {
		return fmt.Errorf("error occurred while creating card %s", result.Error)
	}
	user.Cart = cart
	result = s.db.Create(&user)
	return result.Error
}

func (s *ShopDb) FindUserWithUsername(username string) (*contract.User, error) {
	var user contract.User
	result := s.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, fmt.Errorf("error occurred while creating card %s", result.Error)
	}
	return &user, nil
}

func (s *ShopDb) FindUserWithToken(token string) (*contract.User, error) {
	var user contract.User
	result := s.db.Where("token = ?", token).First(&user)
	if result.Error != nil {
		return nil, fmt.Errorf("error occurred while creating card %s", result.Error)
	}
	return &user, nil
}

func (s *ShopDb) UpdateUser(user *contract.User) error {
	result := s.db.Model(&contract.User{Username: user.Username}).Updates(user)
	if result.Error != nil {
		return fmt.Errorf("error occurred while creating card %s", result.Error)
	}
	return nil
}

func (s *ShopDb) FindUsers() ([]*contract.User, error) {
	var users []*contract.User
	result := s.db.Find(&users)
	if result.Error != nil {
		return nil, fmt.Errorf("error occurred while creating card %s", result.Error)
	}
	return users, nil
}
