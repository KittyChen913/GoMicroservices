package routes

import (
	"logger-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func WriteLog(context *gin.Context) {
	var log models.LogDetail

	err := context.ShouldBindJSON(&log)
	if err != nil {
		context.Error(err)
		return
	}

	err = log.Insert()
	if err != nil {
		context.Error(err)
		return
	}
	context.JSON(http.StatusOK, gin.H{"result": true})
}
