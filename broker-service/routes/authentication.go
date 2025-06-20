package routes

import (
	"broker-service/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func Authentication(context *gin.Context) {

	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.Error(err)
		return
	}
	jsonData, err := json.Marshal(user)
	if err != nil {
		context.Error(err)
		return
	}

	res, err := http.Post(os.Getenv("AuthApiUrl"), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		context.Error(err)
		return
	}

	var jsonFromService AuthResponse

	err = json.NewDecoder(res.Body).Decode(&jsonFromService)
	if err != nil {
		context.Error(err)
		return
	}
	defer res.Body.Close()

	context.JSON(http.StatusOK, jsonFromService)
}

type AuthResponse struct {
	Result       bool
	ErrorMessage string
}
