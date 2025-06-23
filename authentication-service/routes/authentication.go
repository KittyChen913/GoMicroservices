package routes

import (
	"authentication-service/models"
	"authentication-service/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Authentication(context *gin.Context) {

	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.Error(err)
		return
	}

	// 確認此用戶是否存在
	dbUser, err := user.QueryUserByEmail(user.Email)
	if err != nil {
		context.Error(fmt.Errorf("[%v] user not found", user.Email))
		return
	}
	fmt.Println(dbUser)

	// 驗證密碼
	hashCompareResult := utils.CompareHashAndPassword(dbUser.Password, user.Password)
	if hashCompareResult == nil {
		err := WriteLog("authentication", fmt.Sprintf("%v logged in", user.Email))
		if err != nil {
			context.Error(err)
			return
		}
		context.JSON(http.StatusOK, gin.H{"result": true})
	} else {
		context.JSON(http.StatusOK, gin.H{"result": false})
	}
}

func WriteLog(name, data string) error {

	var logDetail struct {
		Name string
		Data string
	}
	logDetail.Name = name
	logDetail.Data = data

	jsonData, _ := json.Marshal(logDetail)
	_, err := http.Post(os.Getenv("LoggerApiUrl"), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	return nil
}
