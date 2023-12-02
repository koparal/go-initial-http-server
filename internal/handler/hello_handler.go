package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}
