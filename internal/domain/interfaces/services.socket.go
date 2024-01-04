package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/jdashel/posts-api/internal/domain/models"
)

type SocketService interface {
	Broadcast(message models.SocketMessage)
	RequestHandler() gin.HandlerFunc
}
