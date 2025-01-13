package wallet

import "github.com/gin-gonic/gin"

func InitRoutes(c *gin.Engine) {
	walletRoutes := c.Group("wallet/")
	{
		walletRoutes.POST("", CreateWallet)
		walletRoutes.GET("/:id/user", GetWalletByUserId)
	}

}
