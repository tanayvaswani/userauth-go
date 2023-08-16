package route

import (
	"github.com/gin-gonic/gin"
	controller "github.com/tanayvaswani/userauth-go/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controller.Signup())
	incomingRoutes.POST("users/login", controller.Login())
}