package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GETUser(c *gin.Context) {
	message := "Welcome to goalng with gorm & postgres"
	c.JSON(http.StatusOK, gin.H{"status": "success", "message":message})
}