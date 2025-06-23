package main

import (
	"logger-service/db"
	"logger-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	db.InitDb()
	routes.RegisterRoutes(server)
	server.Run()
}
