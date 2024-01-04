package models

type SocketMessage struct {
	Type    string `json:"type"`
	Payload any    `json:"payload"`
}
