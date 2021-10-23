package service

import (
	"ecommerce-service/pkg/contract"
	"fmt"
)

type AuthService interface {
	GetUserDetails(token string) (*contract.User, error)
}

func (s *ShopService) GetUserDetails(token string) (*contract.User, error) {
	user, err := s.Db.FindUserWithToken(token)
	if err != nil {
		return nil, fmt.Errorf("eror while fetching user with token %s", err)
	}
	return user,nil
}