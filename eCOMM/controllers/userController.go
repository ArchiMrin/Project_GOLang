package controllers

import (
	"fmt"
	"net/http"

	"github.com/ArchiMrin/Project_GOLang/eCOMM/entities"
	"github.com/ArchiMrin/Project_GOLang/eCOMM/interfaces"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService interfaces.IUser
}

func InitAuthController(authService interfaces.IUser) *AuthController {
	return &AuthController{AuthService: authService}
}

func (a *AuthController) Register(c *gin.Context) {
	fmt.Println("Invoked controller")
	var user entities.User
	err := c.BindJSON(&user)
	if err != nil {
		return
	}
	result, err := a.AuthService.Register(&user)
	fmt.Println(result)
	if err != nil {
		return
	} else {
		c.IndentedJSON(http.StatusCreated, result)
	}
}

func (a *AuthController) Login(c *gin.Context) {
	fmt.Println("Invoked controller")
	var user entities.Login
	err := c.BindJSON(&user)
	if err != nil {
		return
	}
	result, err := a.AuthService.Login(&user)
	fmt.Println(result)
	if err != nil {
		return
	} else {
		c.IndentedJSON(http.StatusCreated, result)
	}
}
func (a *AuthController) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Request.Header.Get("uid")
		var user *entities.User
		user, err := a.AuthService.GetUser(userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}
func (a *AuthController) GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Request.Header.Get("uid")
		var user *entities.User
		user, err := a.AuthService.GetUser(userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

// func (a *AuthController) Logout() string {
// 	return "user logged out"

// }

func Profile(c *gin.Context) {
	//services.GetProfile()
	c.JSON(200, gin.H{
		"data": "Welcome to gin--PROFILE",
	})

}
