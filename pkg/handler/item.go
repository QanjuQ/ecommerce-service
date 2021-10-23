package handler

import (
	"ecommerce-service/pkg/contract"
	"ecommerce-service/pkg/service"
	"net/http"
	"github.com/gin-gonic/gin"
)

type ItemHandler interface {
	CreateItem(itemSvc service.ItemService) gin.HandlerFunc
	FetchItems(itemSvc service.ItemService) gin.HandlerFunc
}

func CreateItem(itemSvc service.ItemService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body contract.CreateItemRequest
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		serviceError := itemSvc.CreateItem(&body)
		if serviceError != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": serviceError.Error()})
			return
		}
		c.String(http.StatusCreated, "created item successfully")
	}
}

func FetchItems(itemSvc service.ItemService) gin.HandlerFunc {
	return func(c *gin.Context) {
		items, serviceError := itemSvc.FetchItems()
		if serviceError != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": serviceError.Error()})
			return
		}
		c.JSON(http.StatusOK, items)
	}
}
