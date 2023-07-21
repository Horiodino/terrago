package jwt

import (
	"os"

	"github.com/gin-gonic/gin"
)

func env() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	// it is creating
	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-v1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success ": "Acess granted for api-1"})
	})

	router.GET("/api-v2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success ": "Acess granted for api-2"})
	})

	router.Run("", port)

}
