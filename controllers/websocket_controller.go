package controllers

import (
	"github.com/gin-gonic/gin"
	"learn/models"
	"learn/services"
	"net/http"
)

type WebSocketController struct {
	service *services.WebsocketService
}

func NewWebSocketController(service *services.WebsocketService) *WebSocketController {
	return &WebSocketController{service: service}
}

func (c *WebSocketController) HandleWebSocket(ctx *gin.Context) {
	conn, err := c.service.Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		http.NotFound(ctx.Writer, ctx.Request)
		return
	}

	c.service.AddClient(conn)
	go c.service.HandleMessage()

	for {
		var msg models.WebsocketMessage
		err := conn.ReadJSON(&msg)
		if err != nil {
			c.service.RemoveClient(conn)
			break
		}
		c.service.BroadcastService(msg)
	}
}

func (c *WebSocketController) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
