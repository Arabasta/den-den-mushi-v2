package websocket

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RegisterWebsocketRoutes(r *gin.RouterGroup, log *zap.Logger, svc *Service) {
	ws := r.Group("/v1")
	ws.GET("/ws", websocketHandler(svc, log))
}

func websocketHandler(svc *Service, log *zap.Logger) gin.HandlerFunc {
	upgrader := svc.websocketUpgrader()

	return func(c *gin.Context) {
		log.Info("websocket upgrader", zap.String("url", c.Request.URL.String()))
		ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Error("Failed to upgrade to websocket", zap.Error(err))
			c.AbortWithStatus(500)
			return
		}

		svc.run(c.Request.Context(), ws)
	}
}
