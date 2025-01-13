package user

import "github.com/gin-gonic/gin"

func InitRoutes(c *gin.Engine) {
	userRoutes := c.Group("/users")
	{
		userRoutes.POST("", CreateUser)
		userRoutes.GET("", GetUsers)
		userRoutes.GET("/:id/wallet", GetUserByWallet)
		userRoutes.PUT("/:id", UpdateUser)
		userRoutes.DELETE("/:id", DeleteUser)
	}

}
