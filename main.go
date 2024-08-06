package main

import (
	"os"

	routes "github.com/SAMGREENHEAD/user_server/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())
	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "welcome"})
	})

	router.GET("/user-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "welcome"})
	})

	router.Run(":" + port)

}
