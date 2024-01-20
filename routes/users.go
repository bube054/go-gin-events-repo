package routes

import (
	"fmt"
	_ "fmt"
	"net/http"

	// "strconv"

	"example.com/learning/models"
	"example.com/learning/utils"
	"github.com/gin-gonic/gin"
)

func SignUp(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		fmt.Println("1")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data!"})
		return
	}

	err = user.Save()

	if err != nil {
		fmt.Println("2")

		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user!"})
		return
	}

	fmt.Println("3")
	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func Login(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data!"})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}
