package wallet

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"user_wallet/db"
	"user_wallet/struct"
)

func CreateWallet(c *gin.Context) {
	var wallet structs.Wallet
	var user structs.User

	var createWalletRQ structs.WalletCreateRQ

	if err := c.BindJSON(&createWalletRQ); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if createWalletRQ.UserId != nil {
		wallet.UserId = *createWalletRQ.UserId
	}
	if createWalletRQ.Balance != nil {
		wallet.Balance = *createWalletRQ.Balance
	}

	err := db.DB.Where("id = ?", wallet.UserId).First(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	err = db.DB.Where("user_id = ?", wallet.UserId).First(&wallet).Error
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
	var user structs.User
	var wallet structs.Wallet

	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	err = db.DB.Where("id = ?", userId).First(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	err = db.DB.Where("user_id = ?", userId).First(&wallet).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not have wallet"})
		return
	}

	if err = db.DB.Find(&wallet, "user_id = ?", userId).Preload("User").Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userWalletResponse := structs.UserWalletResponse{
		User:   user,
		Wallet: wallet,
	}

	c.JSON(http.StatusOK, userWalletResponse)

}
