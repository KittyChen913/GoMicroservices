package routes

import (
	"logger-service/logger"
	"logger-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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

func WriteAccessLog(context *gin.Context) {
	var log models.ELKLogDetail

	err := context.ShouldBindJSON(&log)
	if err != nil {
		context.Error(err)
		return
	}
	request := map[string]any{
		"method": log.Request.Method,
		"url":    log.Request.Url,
		"ip":     log.Request.Ip,
		"body":   log.Request.Body,
	}
	response := map[string]any{
		"status": log.Response.Status,
		"body":   log.Response.Body,
	}
	logger.Log.Info(log.Message,
		zap.String("service.name", log.ServiceName),
		zap.Any("request", request),
		zap.Any("response", response))

	context.JSON(http.StatusOK, gin.H{"result": true})
}
