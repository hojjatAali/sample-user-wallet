package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"user_wallet/db"
	"user_wallet/models"
)

func CreateWallet(c *gin.Context) {

	var wallet models.Wallet

	var user models.User

	userId, err := strconv.Atoi(c.Param("user-id"))

	if err := c.BindJSON(&wallet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = db.DB.Where("id = ?", userId).First(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	err = db.DB.Where("user_id = ?", userId).First(&wallet).Error
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User wallet already exists"})
		return
	}

	if err := db.DB.Create(&wallet).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Wallet created successfully"})

}

func GetWalletByUserId(c *gin.Context) {

	userId, err := strconv.Atoi(c.Param("user-id"))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var wallet models.Wallet
	if err = db.DB.Find(&wallet, "user_id = ?", userId).Preload("User").Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, wallet)

}
