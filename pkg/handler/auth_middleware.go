package handler

import (
	"ecommerce-service/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Authorize(authSvc service.AuthService) gin.HandlerFunc
}

func Authorize(authSvc service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, cookieErr := c.Cookie("token")
		if cookieErr != nil {
			c.String(http.StatusUnauthorized, "please login to get authorized")
		}
		userDetails, err := authSvc.GetUserDetails(token)
		if err != nil {
			c.String(http.StatusUnauthorized, "please login to get authorized")
		}
		c.Set("user", userDetails)
		c.Next()
	}
}
