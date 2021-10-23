package main

import (
	"ecommerce-service/pkg/config"
	"ecommerce-service/pkg/db"
	"ecommerce-service/pkg/handler"
	"ecommerce-service/pkg/service"
	"github.com/gin-gonic/gin"
)

func Pong (c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}


func main() {
	r := gin.Default()
	r.GET("/ping",Pong)
	dbConfig := config.DBConfig{
		Name:     "apnastore",
		Host:     "localhost",
		Username: "shop-admin",
		Password: "password",
		Port:     5432,
	}
	shopDb := db.ShopDb{}
	shopDb.Init(dbConfig.GetConnectionURI())
	shopService := service.ShopService{
		Db: &shopDb,
	}

	itemRouter := r.Group("/item")
	itemRouter.POST("/create",handler.CreateItem(&shopService))
	itemRouter.GET("/list",handler.FetchItems(&shopService))

	cartRouter := r.Group("/cart")
	cartRouter.Use(handler.Authorize(&shopService))
	cartRouter.PUT("/add/items",handler.AddItemToCart(&shopService))
	cartRouter.GET("/list",handler.FetchItems(&shopService))
	cartRouter.PUT("/:cart-id/complete",handler.PurchaseOrder(&shopService))

	orderRouter := r.Group("/order")
	orderRouter.Use(handler.Authorize(&shopService))
	orderRouter.GET("/list",handler.FetchOrders(&shopService))

	r.POST("/user/create",handler.CreateUser(&shopService))
	r.POST("/user/login",handler.UserLogin(&shopService))
	r.GET("/user/list",handler.GetUsers(&shopService))

	r.Run() // 8080
}
