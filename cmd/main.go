package main

import (
	"context"
	root "den-den-mushi-v2"
	"den-den-mushi-v2/internal/config"
	"den-den-mushi-v2/internal/server"
	"den-den-mushi-v2/pkg/logger"
	"flag"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	_ = godotenv.Load(".env")
	cfg := config.Load(configPath())

	log := logger.Init(cfg.Logger, cfg.App)
	if log == nil {
		panic("failed to initialize logger")
	}
	defer func() {
		_ = log.Sync()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-stop
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		time.AfterFunc(10*time.Second, func() {
			os.Exit(0)
		})

		<-ctx.Done()
		os.Exit(0)
	}()

	s := server.New(root.Files, cfg, log)

	if err := server.Start(s, cfg, log); err != nil {
		log.Fatal("failed to start server: %v", zap.Error(err))
	}
}

// configPath usage: go run main.go -config /path/to/config.json
func configPath() string {
	path := flag.String("config", "", "path to config file (optional)")
	flag.Parse()

	var finalPath string
	if *path != "" {
		finalPath = *path
	} else {
		// default to config.json next to executable
		exe, _ := os.Executable()
		exeDir := filepath.Dir(exe)
		finalPath = filepath.Join(exeDir, "config.json")
	}

	return finalPath
}
