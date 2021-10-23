package handler

import (
	"ecommerce-service/pkg/contract"
	"ecommerce-service/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CartHandler interface {
	AddItemToCart(cartSvc service.OrderService) gin.HandlerFunc
	FetchItems(cartSvc service.OrderService) gin.HandlerFunc
}

func AddItemToCart(cartSvc service.OrderService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body []*contract.AddItemToCartRequest
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user := c.Keys["user"].(*contract.User)
		if user != nil {
			serviceError := cartSvc.AddItems(body, user)
			if serviceError != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": serviceError.Error()})
				return
			}
		}
		c.String(http.StatusCreated, "created item successfully")
	}
}

func PurchaseOrder(cartSvc service.OrderService) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.Keys["user"].(*contract.User)
		if user != nil {
			serviceError := cartSvc.Purchase(user)
			if serviceError != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": serviceError.Error()})
				return
			}
		}
		c.String(http.StatusCreated, "created item successfully")
	}
}

func GetCarts(cartSvc service.OrderService) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.Keys["user"].(*contract.User)
		carts, err := cartSvc.FetchCarts(user.ID)

		if user != nil {
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
		c.JSON(http.StatusOK, carts)
	}
}
