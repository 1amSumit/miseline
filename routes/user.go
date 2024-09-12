package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sumit.com/mise-link/model"
	"sumit.com/mise-link/utils"
)

func createUser(context *gin.Context) {
	var user model.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse user data", "error": err.Error()})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "failed to save user", "error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}

func login(context *gin.Context) {

	var loginUserData model.LoginUser

	err := context.BindJSON(&loginUserData)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse login data", "error": err.Error()})
		return
	}

	err = loginUserData.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "invalid credentials", "error": err.Error()})
		return
	}

	token, err := utils.GenerateJWTToken(loginUserData.Email, loginUserData.Id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Authentication failed", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Loged in successfully", "token": token})

}
