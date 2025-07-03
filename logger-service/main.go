package main

import (
	"logger-service/db"
	"logger-service/logger"
	"logger-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	logger.InitLogger()
	defer logger.Log.Sync()

	server := gin.Default()
	db.InitDb()
	routes.RegisterRoutes(server)
	server.Run()
}
