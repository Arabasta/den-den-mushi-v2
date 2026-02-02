package websocket

import (
	"context"
	"den-den-mushi-v2/internal/config"
	"den-den-mushi-v2/internal/connect"
	"den-den-mushi-v2/internal/core/client"
	"den-den-mushi-v2/internal/core/pseudotty"
	"den-den-mushi-v2/pkg/types"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type Service struct {
	connStrategy *connect.Strategy
	log          *zap.Logger
	cfg          *config.Config
}

func NewService(c *connect.Strategy, log *zap.Logger,
	cfg *config.Config) *Service {
	return &Service{
		connStrategy: c,
		log:          log,
		cfg:          cfg,
	}
}

func (s *Service) run(ctx context.Context, ws *websocket.Conn) {
	conn := client.New(ws)

	connMethod, err := s.connStrategy.Get(types.LocalShell)
	if err != nil {
		s.log.Error("Failed to get connection method", zap.Error(err))
		return
	}

	pty, cmd, err := connMethod.Connect(ctx)
	if err != nil {
		s.log.Error("Failed to connect", zap.Error(err))
		return
	}

	err = pseudotty.New(pty, cmd, conn)
	if err != nil {
		s.log.Error("Failed to create pty", zap.Error(err))
		return
	}

	return
}
