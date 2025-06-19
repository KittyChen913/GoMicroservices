package main

import (
	"authentication-service/db"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	db.InitDb()
	server.Run()
}
