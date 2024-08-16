package domain

import "github.com/gorilla/websocket"

// Connection represents a WebSocket connection.
type Connection struct {
    Conn *websocket.Conn
}

// Message represents a WebSocket message.
type Message struct {
    Content []byte
}
