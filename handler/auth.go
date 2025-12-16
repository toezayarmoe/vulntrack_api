package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/toezayarmoe/vulntrack_api/config"
	"github.com/toezayarmoe/vulntrack_api/models"
	"github.com/toezayarmoe/vulntrack_api/utils"
	"golang.org/x/crypto/bcrypt"
)

func Login(ctx *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var user models.Users
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid Input"})
		return
	}

	if err := config.DB.Where("username= ?", input.Username).First(&user).Error; err != nil {
		ctx.JSON(401, gin.H{"error": "Invalid Credentials"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Input ", string(hash))
	fmt.Println("SQL ", user.Password_Hash)
	if bcrypt.CompareHashAndPassword([]byte(user.Password_Hash), []byte(input.Password)) != nil {

		ctx.JSON(401, gin.H{"error": "Invalid Credentials"})
		return
	}

	token, _ := utils.GenerateToken(user.ID, user.IsAdmin)

	ctx.JSON(200, gin.H{"token": token})

}
