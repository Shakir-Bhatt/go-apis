package main

import (
	"simpleapi_go/controllers"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := SetupRouter()
	router.Run("localhost:8083")
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/products", getProducts)

	router.GET("/product/:id", getProduct)

	// router.POST("/products", addProduct)

	return router
}

func getProducts(c *gin.Context) {

	products := controllers.GetProducts()

	if products == nil || len(products) == 0 {

		c.AbortWithStatus(http.StatusNotFound)

	} else {

		c.IndentedJSON(http.StatusOK, products)

	}
}

func getProduct(c *gin.Context) {

	code := c.Param("code")

	product := controllers.GetProduct(code)

	if product == nil {
		c.AbortWithStatus(http.StatusNotFound)

	} else {

		c.IndentedJSON(http.StatusOK, product)

	}

}

// func addProduct(c *gin.Context) {

// 	var prod models.Product

// 	if err := c.BindJSON(&prod); err != nil {
// 		c.AbortWithStatus(http.StatusBadRequest)
// 	} else {
// 		controllers.AddProduct(prod)
// 		c.IndentedJSON(http.StatusCreated, prod)
// 	}

// }
