package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hellow(c *gin.Context) {
	token := c.GetString("Token")
	c.JSON(http.StatusOK, gin.H{
		"token":    token,
		"response": "Hello world",
	})
}
func Hellow2(c *gin.Context) {
	role := c.GetString("role")
	c.JSON(http.StatusOK, gin.H{
		"token":    role,
		"response": "Hello world",
	})
}
