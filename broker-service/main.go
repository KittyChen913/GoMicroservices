package main

import (
	"broker-service/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.Use(cors.Default())
	routes.RegisterRoutes(server)
	server.Run()
}
