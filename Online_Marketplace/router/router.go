package router

import (
	"net/http"
	"onlinemarketplace/controller"
	"os"

	log "github.com/sirupsen/logrus"

	"onlinemarketplace/router/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

	file, err := os.Create("online_marketplace_error.log")
	if err != nil {
		panic(err)
	}
	log.SetOutput(file)
	r := gin.Default()
	r.Use(gin.LoggerWithWriter(file))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "welcome..."})
	})

	r.POST("/signup", func(c *gin.Context) {
		controller.Signup(c)
	})

	r.POST("/login", func(c *gin.Context) {
		controller.Login(c)
	})

	r.GET("/products", func(c *gin.Context) {
		controller.GetProduct(c)
	})

	r.GET("/products/:id", func(c *gin.Context) {
		controller.GetProductById(c)
	})

	r.Use(middleware.Authtoken())

	r.DELETE("/products/:id", func(c *gin.Context) {
		controller.DeleteProduct(c)
	})

	r.POST("/products", func(c *gin.Context) {
		controller.CreateProduct(c)
	})

	r.PUT("/products/:id", func(c *gin.Context) {
		controller.UpdateProduct(c)
	})

	return r
}
