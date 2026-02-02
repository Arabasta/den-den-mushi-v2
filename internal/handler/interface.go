package handler

import (
	"den-den-mushi-v2/internal/protocol"
	"io"
)

// Handler should not have any websocket writes
type Handler interface {
	Handle(pkt protocol.Packet, pty io.Writer) (string, error)
}
