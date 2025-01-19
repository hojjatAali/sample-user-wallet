package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"user_wallet/service"
	"user_wallet/struct"
)

func CreateUser(c *gin.Context) {
	var user structs.UserCreateRQ

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userService := service.UserService{}

	newUser, err := userService.CreateUser(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.New("can not create user")})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

func GetUser(c *gin.Context) {

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userService := service.UserService{}
	userWalletResponse, err := userService.GetUser(userID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if userWalletResponse.Wallet.ID == 0 {
		c.JSON(http.StatusOK, userWalletResponse.User)
		return
	}

	c.JSON(http.StatusOK, userWalletResponse)

}

func UpdateUser(c *gin.Context) {
	var user structs.UserUpdateRQ
	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err = c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userService := service.UserService{}

	updatedUser, err := userService.UpdateUser(userId, user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

func DeleteUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))

	userService := service.UserService{}

	err = userService.DeleteUser(userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
func GetUsers(c *gin.Context) {
	userService := service.UserService{}
	users, err := userService.GetUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, users)

}
