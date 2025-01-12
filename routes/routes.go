package routes

import (
	"github.com/gin-gonic/gin"
	"user_wallet/handlers"
)

func InitRoutes(r *gin.Engine) {
	r.POST("/users", handlers.CreateUser)
	r.GET("/users", handlers.GetUsers)
	r.PUT("/users/:id", handlers.UpdateUser)
	r.DELETE("/users/:id", handlers.DeleteUser)
	r.GET("/users/:id/wallet", handlers.GetUserByWallet)

	r.POST("/wallet/:user-id", handlers.CreateWallet)
	//r.PUT("/wallets/:id")
	r.GET("/Wallets/:user-id", handlers.GetWalletByUserId)

	r.GET("/health", handlers.Health)

}
