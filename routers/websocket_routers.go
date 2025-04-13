package routers

import (
	"learn/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterWebSocketRouter(group gin.RouterGroup, controller *controllers.WebSocketController) {
	group.GET("/ws", controller.HandleWebSocket)
}
