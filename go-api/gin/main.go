package main

import (
	"api/controller"

	"github.com/gin-gonic/gin"
)

var (
	// db             *gorm.DB                  = config.SetupDatabaseConnection()
	authController controller.AuthController = controller.NewAuthController()
)

func main() {
	// defer config.CloseDatabaseConnection(db)
	router := gin.Default()

	authRoutes := router.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	router.Run()
}
