package pingController

import (
	"github.com/brayanlopez8/GolangComicsStore/tools"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	tools.GenerateKeyMarvel()
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
