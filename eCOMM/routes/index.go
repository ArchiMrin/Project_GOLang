package routes

import (
	"github.com/ArchiMrin/Project_GOLang/eCOMM/controllers"
	"github.com/ArchiMrin/Project_GOLang/eCOMM/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, a controllers.AuthController) {
	// user routes
	user := r.Group("/api/user")
	user.POST("/register", a.Register)
	//user.POST("/register", controllers.Register)
	user.POST("/login", a.Login)
	//user.GET("/logout", a.Logout)
	//user.GET("/profile/:id", controllers.Profile)
}

//Product Routes

func ProductRoutes(r *gin.Engine, p controllers.ProductController) {
	product := r.Group("/api/product")
	product.POST("/insert", p.InsertProduct)
	product.GET("/getProducts", p.GetProducts)
	product.GET("/getProducts/:id", p.GetProductById)

}

func AuthRoutes(incomingRoutes *gin.Engine, a controllers.AuthController) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/api/users/usersdata", a.GetUser())
	incomingRoutes.Use(middleware.Authorize())
	incomingRoutes.GET("/api/users/getallusers", a.GetAllUsers())

}
