package core_helpers

import (
	"den-den-mushi-v2/internal/core/client"
	"den-den-mushi-v2/internal/protocol"
)

// SendToConn sends packet to a specific connection, used for targeted messages
func SendToConn(c *client.Connection, pkt protocol.Packet) {
	if c == nil {
		return
	}

	Send(c.WsWriteCh, pkt)
}

// Send just sends a packet to a channel
func Send(ch chan protocol.Packet, pkt protocol.Packet) {
	defer func() {
		if r := recover(); r != nil {
		}
	}()

	select {
	case ch <- pkt:
	default:
		// queue full
	}
}
