package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mVedr/jwt_auth_go/routes"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	routes.AuthRoutes(router)

	router.GET("/api-1", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"sucess": "Access granted for api-1"})
	})

	router.GET("/api-2", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"sucess": "Access granted for api-2"})
	})

	router.Run("localhost:" + port)

}
