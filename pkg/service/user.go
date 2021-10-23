package service

import (
	"ecommerce-service/pkg/contract"
	"ecommerce-service/pkg/db"
	"fmt"
)

type ShopService struct {
	Db db.ShopRepo
}

type UserService interface {
	CreateUser(request contract.CreateUserRequest) (string,error)
	ValidateCredentials(request contract.LoginRequest) (string, error)
	FetchUsers() ([]*contract.User, error)
}

func generateRandomToken() string {
	return "some dummy token for now"
}

func (s *ShopService) CreateUser(request contract.CreateUserRequest) (string, error) {
	token := generateRandomToken()
	user := contract.User{
		Name:     request.Name,
		Email:    request.Email,
		Username: request.Username,
		Password: request.Password,
		Token:    token,
		Cart: contract.Cart{},
	}
	err := s.Db.CreateUser(&user)
	if err != nil {
		return "",err
	}
	return token,nil
}

func (s *ShopService) ValidateCredentials(request contract.LoginRequest) (string, error) {
	token := generateRandomToken()
	user, err := s.Db.FindUserWithUsername(request.Username)
	if err != nil {
		return "",fmt.Errorf("eror while fetching user with username: %s",err)
	}
	if request.Username != user.Username && request.Password != request.Password {
		return "", fmt.Errorf("incorrect password or user name: for username: %s",request.Username)
	}
	user.Token = token
	err = s.Db.UpdateUser(user)
	if err != nil {
		return "", fmt.Errorf("error while saving token for user: %s",request.Username)
	}
	return token,nil
}

func (s *ShopService) FetchUsers() ([]*contract.User, error) {
	users, err := s.Db.FindUsers()
	if err != nil {
		return []*contract.User{},fmt.Errorf("eror while fetching users: %s",err)
	}
	return users,err
}
