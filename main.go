package main

import (
	"github.com/kofra002/Login/config"
	"github.com/kofra002/Login/controllers"
	"github.com/kofra002/Login/models"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{})

	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.POST("/refresh", controllers.RefreshToken)
	r.POST("/logout", controllers.AuthMiddleware(), controllers.Logout)

	protected := r.Group("/api")
	protected.Use(controllers.AuthMiddleware())
	protected.GET("/protected", controllers.Protected)

	r.Run(":8080")
}
