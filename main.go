package main

import (
	"github.com/gin-gonic/gin"
	"user_wallet/db"
	"user_wallet/handlers/user"
	"user_wallet/handlers/wallet"
)

func main() {

	db.Connect()

	router := gin.Default()

	user.InitRoutes(router)
	wallet.InitRoutes(router)

	router.Run(":8000")
}
