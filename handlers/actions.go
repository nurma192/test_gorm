package handlers

import (
	"Go_rest_now_android/database"
	"Go_rest_now_android/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Registration(c *gin.Context) {
	var user *model.User

	decode := json.NewDecoder(c.Request.Body).Decode(&user)
	if decode != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"response": decode.Error(),
		})
		return
	}

	isExist := database.IsExistUserByLogin(user.Login)
	if isExist {
		c.JSON(http.StatusOK, gin.H{
			"response": "User with this login is exist",
		})
		return
	}

	newUser := &model.User{Id: user.Id, Login: user.Login, Name: user.Name, Pass: user.Pass}
	isSuccessAddError := database.Add(newUser)

	if isSuccessAddError == nil {
		c.JSON(http.StatusOK, gin.H{
			"response": gin.H{
				"code":     http.StatusOK,
				"response": "You success registration",
			},
		})
	}
}
