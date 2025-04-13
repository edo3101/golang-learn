package routers

import (
	"learn/controllers"

	"github.com/gin-gonic/gin"
)

type Router struct {
	App *gin.Engine
}

func NewRouter(app *gin.Engine, baseController *controllers.WebSocketController) *Router {
	RegisterRouter(app, baseController)
	return &Router{
		App: app,
	}
}

func RegisterRouter(app *gin.Engine, baseController *controllers.WebSocketController) {
	routerGroup := app.RouterGroup
	RegisterSchedulerRouter(routerGroup, baseController)
	RegisterWebSocketRouter(routerGroup, baseController)
}
