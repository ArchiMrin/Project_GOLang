package middleware

import (
	"fmt"

	"github.com/ArchiMrin/Project_GOLang/eCOMM/utils"
	"github.com/gin-gonic/gin"

	"net/http"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization Header Provided")})
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("first_name", claims.First_name)
		c.Set("last_name", claims.Last_name)
		c.Set("uid", claims.Uid)
		c.Set("user_type", claims.User_type)
		c.Request.Header.Set("uid", claims.Uid)
		c.Request.Header.Set("user_type", claims.User_type)
		c.Next()
	}
}

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_type := c.Request.Header.Get("user_type")
		if user_type == "admin" {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("User is not authorized")})
			c.Abort()
			return
		}

	}
}
