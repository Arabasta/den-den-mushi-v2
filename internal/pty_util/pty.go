package pty_util

import (
	"os"
	"os/exec"

	"github.com/creack/pty"
	"go.uber.org/zap"
)

func Spawn(cmd *exec.Cmd, log *zap.Logger) (*os.File, *exec.Cmd, error) {
	ptmx, err := pty.Start(cmd)
	if err != nil {
		return nil, nil, err
	}

	go func() {
		// retrieve exit status from OS and clean up process table entry
		// for preventing zombie
		err := cmd.Wait()
		if err != nil {
			log.Warn("SSH exited with error", zap.Error(err))
		} else {
			log.Info("SSH exited cleanly")
		}
	}()

	return ptmx, cmd, nil
}

func Resize(ptmx *os.File, cols, rows uint16) error {
	return pty.Setsize(ptmx, &pty.Winsize{
		Cols: cols,
		Rows: rows,
	})
}
