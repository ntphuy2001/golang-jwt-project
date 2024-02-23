package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	env "github.com/ntphuy2001/golang-jwt-project/dotenvsingleton"
	"github.com/ntphuy2001/golang-jwt-project/routes"
)

func main() {
	port := env.GetDotenvInstance().GetPort()
	if port == "" {
		port = "8000"
	}
	fmt.Println(port)
	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted to api-1"})
	})

	router.GET("api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted to api-2"})
	})

	err := router.Run(":" + port)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
