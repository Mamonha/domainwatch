package main

import (
	"gin-quickstart/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	api := router.Group("/api")

	api.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	domainRoutes := api.Group("/domain")

	domainRoutes.GET("/validate/:domain", func(c *gin.Context) {
		domain := c.Param("domain")
		controllers.Validate(domain, c.Writer)
	})

	domainRoutes.GET("/whois/:domain", func(c *gin.Context) {
		domain := c.Param("domain")
		controllers.Whois(domain, c.Writer)
	})
	domainRoutes.POST("/tls", func(c *gin.Context) {
		controllers.CheckTLS(c.Writer, c.Request)
	})

	router.Run()
}
