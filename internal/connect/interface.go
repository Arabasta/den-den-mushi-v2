package connect

import (
	"context"
	"os"
	"os/exec"
)

type ConnectionMethod interface {
	Connect(ctx context.Context) (*os.File, *exec.Cmd, error)
}
