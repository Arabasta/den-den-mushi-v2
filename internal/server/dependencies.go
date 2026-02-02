package server

import (
	"den-den-mushi-v2/internal/config"
	connect2 "den-den-mushi-v2/internal/connect"

	"den-den-mushi-v2/internal/pty_util"

	"den-den-mushi-v2/internal/websocket"

	"go.uber.org/zap"
)

type Deps struct {
	WebsocketService *websocket.Service
}

func initDependencies(cfg *config.Config, log *zap.Logger) *Deps {

	connMethodStrategy := connect2.NewRegistry(
		connect2.NewDeps(
			pty_util.NewBuilder(log, cfg.Ssh),
			cfg,
			log))

	websocketService := websocket.NewService(connMethodStrategy, log, cfg)

	return &Deps{
		WebsocketService: websocketService,
	}
}
