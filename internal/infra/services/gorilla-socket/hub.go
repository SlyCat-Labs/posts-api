package websockets

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	// You can customize options here, such as:
	// ReadBufferSize:  1024,
	// WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Add your origin checking logic here
		return true // Example: Allow all origins for now
	},
}

type Hub struct {
	clients    []*Client
	register   chan *Client
	unregister chan *Client
	mutex      *sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		clients:    make([]*Client, 0),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		mutex:      &sync.Mutex{},
	}
}

func (hub *Hub) HandleSocket() gin.HandlerFunc {
	return func(c *gin.Context) {
		socket, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"Could not open connection": err.Error()})
			return
		}

		client := NewClient(hub, socket)
		hub.register <- client

		go client.Write()
	}
}

func (hub *Hub) Run() {
	for {
		select {
		case client := <-hub.register:
			hub.onConnect(client)
		case client := <-hub.unregister:
			hub.onDisconnect(client)
		}
	}
}

func (hub *Hub) onConnect(client *Client) {
	hub.mutex.Lock()
	defer hub.mutex.Unlock()

	client.id = client.socket.RemoteAddr().String()
	hub.clients = append(hub.clients, client)

	log.Println("Client connected")
}

func (hub *Hub) onDisconnect(client *Client) {
	hub.mutex.Lock()
	client.socket.Close()
	defer hub.mutex.Unlock()

	for i, other := range hub.clients {
		if other.id == client.id {
			hub.clients = append(hub.clients[:i], hub.clients[i+1:]...)
			break
		}
	}

	log.Println("Client disconnected")
}

func (hub *Hub) Broadcast(message any, ignore *Client) {
	data, _ := json.Marshal(message)

	for _, client := range hub.clients {
		if client != ignore {
			client.outbound <- data
		}
	}
}
