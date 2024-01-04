package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jdashel/posts-api/internal/domain/interfaces"
)

type WebsocketHandler struct {
	socketService interfaces.SocketService
}

func NewWebsocketHandler(socketService interfaces.SocketService) *WebsocketHandler {
	return &WebsocketHandler{socketService}
}

func (wsh *WebsocketHandler) RequestHandler() gin.HandlerFunc {
	return wsh.socketService.RequestHandler()
}
