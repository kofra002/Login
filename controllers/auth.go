package controllers

import (
	"net/http"

	"github.com/kofra002/Login/config"
	"github.com/kofra002/Login/models"
	"github.com/kofra002/Login/utils"
	"github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existing models.User
	if err := config.DB.Where("username = ?", user.Username).First(&existing).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exist"})
		return
	}

	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword
	config.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var input models.User
	var user models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid crednetials"})
		return
	}

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, _ := utils.GenerateJWT(user.Username)
	refreshToken, _ := utils.GenerateRefreshToken(user.Username)

	user.RefreshToken = refreshToken
	config.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{"token": token, "refresh_token": refreshToken})
}

func RefreshToken(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := jwt.Parse(req.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		return utils.RefreshSecret, nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	var user models.User
	config.DB.Where("username = ?", username).First(&user)
	if user.RefreshToken != req.RefreshToken {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token mismatch"})
		return
	}

	newToken, _ := utils.GenerateJWT(username)
	c.JSON(http.StatusOK, gin.H{"token": newToken})
}

func Logout(c *gin.Context) {
	username, _ := c.Get("username")
	var user models.User
	config.DB.Where("username = ?", username).First(&user)
	user.RefreshToken = ""
	config.DB.Save(&user)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out succesfully"})
}
