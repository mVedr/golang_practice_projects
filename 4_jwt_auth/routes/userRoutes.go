package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mVedr/jwt_auth_go/controllers"
	"github.com/mVedr/jwt_auth_go/middlewares"
)

func UserRoutes(router *gin.Engine) {
	router.Use(middlewares.Authenticate())
	router.GET("/users", controllers.GetUsers())
	router.GET("/users/:id", controllers.GetUser())
}
