package pseudotty

import (
	"den-den-mushi-v2/internal/protocol"

	"github.com/labstack/gommon/log"
	"go.uber.org/zap"
)

func (s *Session) readPtyLoop() {
	defer func() {
		if r := recover(); r != nil {
			log.Error("panic", zap.Any("panic", r), zap.Stack("stack"))
		}
	}()

	buf := make([]byte, 4096)
	for {
		n, err := s.pty.Read(buf)
		if err != nil {
			return
		}

		data := append([]byte{}, buf[:n]...)
		pkt := protocol.Packet{
			Header: protocol.Output,
			Data:   data,
		}

		s.fanout(pkt)
	}
}
