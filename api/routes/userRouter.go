package route

import (
	"github.com/gin-gonic/gin"
	controller "github.com/tanayvaswani/userauth-go/controllers"
	"github.com/tanayvaswani/userauth-go/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}