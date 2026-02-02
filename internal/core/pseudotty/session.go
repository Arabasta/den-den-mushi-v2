package pseudotty

import (
	"context"

	"den-den-mushi-v2/internal/core/client"

	"os"
	"os/exec"
	"sync"
)

type Session struct {
	pty  *os.File
	cmd  *exec.Cmd
	conn *client.Connection

	ctx    context.Context
	cancel context.CancelFunc

	mu   sync.RWMutex
	once sync.Once
}

func New(pty *os.File, cmd *exec.Cmd, c *client.Connection) error {
	s := &Session{
		pty: pty,
		cmd: cmd,
	}

	s.ctx, s.cancel = context.WithCancel(context.Background())
	return s.setup(c)
}

func (s *Session) setup(c *client.Connection) error {
	s.conn = c
	go s.conn.WriteClient()
	go s.conn.PrimaryReadLoop(s.handleConnPacket)
	go s.readPtyLoop()
	return nil
}
