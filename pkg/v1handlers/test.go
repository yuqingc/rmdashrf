package v1handlers

import (
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello rmdashrf!",
	})
}
