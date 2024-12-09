package controller

import (
	"github.com/gin-gonic/gin"
)

func Greetings(c *gin.Context) {
	c.IndentedJSON(200, gin.H{
		"message": "Hello World",
	})
}
