package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {

	broker := server.Group("/")
	broker.POST("/Authentication", Authentication)
}
