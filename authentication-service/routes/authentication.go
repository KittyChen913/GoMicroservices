package routes

import (
	"authentication-service/models"
	"authentication-service/utils"
	"fmt"
	"net/http"

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
		context.JSON(http.StatusOK, gin.H{"result": true})
	} else {
		context.JSON(http.StatusOK, gin.H{"result": false})
	}
}
