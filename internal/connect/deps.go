package connect

import (
	"den-den-mushi-v2/internal/config"
	"den-den-mushi-v2/internal/pty_util"

	"go.uber.org/zap"
)

type Deps struct {
	commandBuilder *pty_util.Builder
	cfg            *config.Config
	log            *zap.Logger
}

func NewDeps(commandBuilder *pty_util.Builder, cfg *config.Config, log *zap.Logger) Deps {
	return Deps{
		commandBuilder: commandBuilder,
		cfg:            cfg,
		log:            log,
	}
}
