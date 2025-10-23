package utils

import (
	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(code, gin.H{
		"success": true,
		"message": msg,
		"data":    data,
	})
}

func Error(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{
		"success": false,
		"error":   msg,
	})
}
