package routers

import (
	"learn/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterSchedulerRouter(group gin.RouterGroup, controller *controllers.WebSocketController) {
	// Add routes
	group.GET("/ping", controller.Ping)
}
