package client

import (
	"den-den-mushi-v2/internal/protocol"

	"github.com/gorilla/websocket"
)

type Connection struct {
	Sock *websocket.Conn

	WsWriteCh chan protocol.Packet
}

func New(sock *websocket.Conn) *Connection {
	return &Connection{
		Sock:      sock,
		WsWriteCh: make(chan protocol.Packet, 100),
	}
}
