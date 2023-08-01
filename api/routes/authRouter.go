package routes

import (
	controller "JWT/api/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	// signup
	incomingRoutes.POST("user/signup", controller.Signup())

	// login
	incomingRoutes.POST("user/login", controller.Login())
}