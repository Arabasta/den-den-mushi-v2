package client

import (
	"den-den-mushi-v2/internal/protocol"
)

func (c *Connection) PrimaryReadLoop(onPacket func(protocol.Packet)) {
	defer func() {
		if r := recover(); r != nil {
		}
	}()

	for {
		msgType, msg, err := c.Sock.ReadMessage()
		if err != nil {
			return
		}

		pkt := protocol.Parse(msgType, msg)
		if pkt.Header == protocol.ParseError {
			continue
		}

		onPacket(pkt)
	}
}
