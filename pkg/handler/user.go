package handler

import (
	"ecommerce-service/pkg/contract"
	"ecommerce-service/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler interface {
	CreateUser(userSvc service.UserService) gin.HandlerFunc
	UserLogin(userSvc service.UserService) gin.HandlerFunc
	GetUsers(userSvc service.UserService) gin.HandlerFunc
}

func CreateUser(userSvc service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body contract.CreateUserRequest
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		token, serviceError := userSvc.CreateUser(body)
		if serviceError != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": serviceError.Error()})
			return
		}
		c.SetCookie("token", token, 3600, "/", "localhost", false, true)
		c.String(http.StatusCreated, "created account successfully")
	}
}

func UserLogin(userSvc service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body contract.LoginRequest
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		token, serviceError := userSvc.ValidateCredentials(body)
		if serviceError != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": serviceError.Error()})
			return
		}
		c.SetCookie("token", token, 3600, "/", "localhost", false, true)
		c.String(http.StatusOK, "logged in successfully")
	}
}

func GetUsers(userSvc service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, serviceError := userSvc.FetchUsers()
		if serviceError != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": serviceError.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}
