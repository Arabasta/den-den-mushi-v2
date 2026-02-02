package pty_util

import (
	"den-den-mushi-v2/pkg/config"
	"os"
	"os/exec"

	"go.uber.org/zap"
)

type Builder struct {
	log *zap.Logger
	cfg *config.Ssh
}

func NewBuilder(log *zap.Logger, cfg *config.Ssh) *Builder {
	return &Builder{log: log, cfg: cfg}
}

func (b *Builder) BuildBashCmd() *exec.Cmd {
	cmd := exec.Command("bash")
	cmd.Env = append(os.Environ(), "TERM=xterm-256color")
	cmd.Dir = os.Getenv("HOME")

	return cmd
}
