package controllers

import (
	"fmt"
	"net/http"

	"github.com/ArchiMrin/Project_GOLang/eCOMM/entities"
	"github.com/ArchiMrin/Project_GOLang/eCOMM/interfaces"
	"github.com/gin-gonic/gin"
	//"github.com/ArchiMrin/tecnotree-go-advanced/ekart_1/services"
)

type ProductController struct {
	ProductService interfaces.IProduct
	//ProductController services.ProductService //this is our hard coupling dependency which we won't do
}

func InitProductController(productSvc interfaces.IProduct) *ProductController {
	return &ProductController{ProductService: productSvc}
}

func (p ProductController) InsertProduct(c *gin.Context) {
	fmt.Println("Invoked Controller")
	//extract the product from the Request Object
	var product entities.Product
	err := c.BindJSON(&product) // we call the rest api to extract the product details
	if err != nil {
		return
	}
	//call the service.Insert
	result, err := p.ProductService.Insert(&product)
	if err != nil {
		return
	} else {
		c.IndentedJSON(http.StatusCreated, result)
	}
	//call p.Product service.InsertOne()
}
func (p ProductController) GetProducts(c *gin.Context) {
	result, err := p.ProductService.GetProducts()
	if err != nil {
		return
	} else {
		c.IndentedJSON(http.StatusCreated, result)
	}
}
func (p ProductController) GetProductById(c *gin.Context) {

	productID := c.Param("_id")

	product, err := p.ProductService.GetProductById(productID)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve desired product"})

		return

	}

	if product == nil {

		c.JSON(http.StatusNotFound, gin.H{"message": "Product not found in DB"})

		return

	}

	c.IndentedJSON(http.StatusOK, product)

}
