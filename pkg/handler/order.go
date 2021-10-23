package handler

import (
	"ecommerce-service/pkg/contract"
	"ecommerce-service/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderHandler interface {
	FetchOrders(svc service.OrderService) gin.HandlerFunc
}

func FetchOrders(svc service.OrderService) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.Keys["user"].(*contract.User)
		orders,err := svc.FetchOrders(user.ID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, orders)
	}
}
