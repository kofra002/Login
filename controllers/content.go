package controllers

import (
	"net/http"

	"github.com/kofra002/config"
	"github.com/gin-gonic/gin"
)

func updateContent(c *gin.Context) {
	var secret struct (
		content string `json:"content"`
	)
	if err := c.ShouldBindJSON(&secret); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
