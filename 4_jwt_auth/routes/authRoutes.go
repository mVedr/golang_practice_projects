package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mVedr/jwt_auth_go/controllers"
)

func AuthRoutes(router *gin.Engine) {
	router.POST("/users/signup", controllers.Signup())
	router.POST("/users/login", controllers.Login())
}
