package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandle(context *gin.Context) {
	context.Next()

	if len(context.Errors) > 0 {
		err := context.Errors.Last().Err
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"ErrorMessage": err.Error()})
	}
}
