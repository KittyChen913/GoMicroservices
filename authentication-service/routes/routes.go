package routes

import (
	"authentication-service/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.Use(middlewares.ErrorHandle)

	authentication := server.Group("/")
	authentication.POST("/Authentication", Authentication)
}
