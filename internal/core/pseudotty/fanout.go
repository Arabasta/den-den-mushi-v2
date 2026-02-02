package pseudotty

import (
	"den-den-mushi-v2/internal/core/core_helpers"
	"den-den-mushi-v2/internal/handler"
	"den-den-mushi-v2/internal/protocol"
)

func (s *Session) handleConnPacket(pkt protocol.Packet) {
	h, _ := handler.Get[pkt.Header]
	_, err := h.Handle(pkt, s.pty)
	if err != nil {
		return
	}
}

func (s *Session) fanout(pkt protocol.Packet) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	core_helpers.SendToConn(s.conn, pkt)
}
