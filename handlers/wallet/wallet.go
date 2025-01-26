package wallet

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"user_wallet/service"
	"user_wallet/struct"
)

var wS = service.WalletService{}

func CreateWallet(c *gin.Context) {

	var createWalletRQ structs.WalletCreateRQ

	if err := c.BindJSON(&createWalletRQ); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	wallet, err := wS.CreateWallet(createWalletRQ)

	if err != nil {
		var status int
		switch err.Error() {
		case "user not found":
			status = http.StatusNotFound
		case "user wallet already exists":
			status = http.StatusConflict
		default:
			status = http.StatusInternalServerError
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Wallet created successfully"})
	c.JSON(http.StatusCreated, wallet)
}

func GetWalletByUserId(c *gin.Context) {

	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	user, wallet, err := wS.GetUserWallet(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	userWalletResponse := structs.UserWalletResponse{
		User:   user,
		Wallet: wallet,
	}

	c.JSON(http.StatusOK, userWalletResponse)

}
