package client

import (
	"den-den-mushi-v2/internal/protocol"
	"github.com/gorilla/websocket"
)

func (c *Connection) WriteClient() {
	defer func() {
		if r := recover(); r != nil {
		}
	}()

	for {
		pkt, ok := <-c.WsWriteCh
		if !ok {
			return
		}
		if err := c.Sock.WriteMessage(websocket.BinaryMessage, protocol.PacketToByte(pkt)); err != nil {
			return
		}
	}
}
