package connect

import (
	"context"
	"den-den-mushi-v2/internal/config"
	"den-den-mushi-v2/internal/pty_util"
	"errors"
	"os"
	"os/exec"

	"go.uber.org/zap"
)

type LocalShellConnection struct {
	cfg            *config.Config
	log            *zap.Logger
	commandBuilder *pty_util.Builder
}

func (c *LocalShellConnection) Connect(_ context.Context) (*os.File, *exec.Cmd, error) {
	if !c.cfg.Ssh.IsLocalSshKeyEnabled {
		return nil, nil, errors.New("LocalShell is not supported in this build")
	}

	cmd := c.commandBuilder.BuildBashCmd()
	return pty_util.Spawn(cmd, c.log)
}
