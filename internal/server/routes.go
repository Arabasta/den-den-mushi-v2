package server

import (
	"den-den-mushi-v2/internal/websocket"
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func registerWebsocketRoutes(r *gin.Engine, deps *Deps, log *zap.Logger) {
	unprotected := r.Group("")
	websocket.RegisterWebsocketRoutes(unprotected, log, deps.WebsocketService)
}

func addStaticRoutes(r *gin.Engine, staticFiles embed.FS, log *zap.Logger) {
	subFS, err := fs.Sub(staticFiles, "static")
	if err != nil {
		log.Fatal("Failed to load embedded static files", zap.Error(err))
	}

	r.StaticFS("/static", http.FS(subFS))

	r.GET("/", func(c *gin.Context) {
		data, err := fs.ReadFile(subFS, "index.html")
		if err != nil {
			c.String(http.StatusInternalServerError, "failed to load index.html")
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", data)
	})
}
