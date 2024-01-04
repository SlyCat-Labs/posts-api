package services

import (
	"github.com/gin-gonic/gin"
	"github.com/jdashel/posts-api/internal/domain/models"
	websockets "github.com/jdashel/posts-api/internal/infra/services/gorilla-socket"
)

type GorillaSocketService struct {
	Hub *websockets.Hub
}

func NewGorillaSocketService() *GorillaSocketService {
	service := &GorillaSocketService{
		Hub: websockets.NewHub(),
	}

	go service.Hub.Run()

	return service
}

func (socket *GorillaSocketService) Broadcast(message models.SocketMessage) {
	socket.Hub.Broadcast(message, nil)
}

func (socket *GorillaSocketService) RequestHandler() gin.HandlerFunc {
	return socket.Hub.HandleSocket()
}
