package db

import (
	"ecommerce-service/pkg/contract"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ShopDb struct {
	db *gorm.DB
}

type ShopRepo interface {
	UserRepo
	CartRepo
	OrderRepo
	ItemRepo
}

func (s *ShopDb) Init(dbUri string) error {
	conn, err := gorm.Open(postgres.New(postgres.Config{DSN: dbUri, PreferSimpleProtocol: true}), &gorm.Config{})
	if err != nil {
		fmt.Print(err)
		return err
	}

	s.db = conn
	s.db.Debug().AutoMigrate(&contract.User{}, &contract.Cart{}, &contract.Item{}, &contract.Order{}, &contract.CartItem{})
	
	return nil
}

func (s *ShopDb) Close()  {
	s.db.Statement.ReflectValue.Close()
}