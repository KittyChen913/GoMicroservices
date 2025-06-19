package main

import (
	"authentication-service/db"
	"authentication-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	db.InitDb()
	routes.RegisterRoutes(server)
	server.Run()
}
