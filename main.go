package main

import (
	"github.com/gin-gonic/gin"
	"user_wallet/db"
	"user_wallet/routes"
)

func main() {

	db.Connect()

	router := gin.Default()

	routes.InitRoutes(router)

	router.Run(":8000")
}
