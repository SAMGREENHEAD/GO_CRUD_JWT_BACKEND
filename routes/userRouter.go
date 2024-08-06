package routes

import (
	controller "github.com/SAMGREENHEAD/user_server/controllers"
	"github.com/SAMGREENHEAD/user_server/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/user/:user_id", controller.GetUser())
}
