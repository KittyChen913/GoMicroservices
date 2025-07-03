package routes

import (
	"logger-service/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.Use(middlewares.ErrorHandle)

	logger := server.Group("/")
	logger.POST("/WriteLog", WriteLog)
	logger.POST("/WriteAccessLog", WriteAccessLog)
}
