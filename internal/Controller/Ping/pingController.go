package pingcontroller

import (
	"fmt"

	"github.com/brayanlopez8/GolangComicsStore/tools"
	"github.com/gin-gonic/gin"
)

//Ping test service
func Ping(c *gin.Context) {
	var key = tools.GenerateKeyMarvel()
	fmt.Println(key)
	// fmt.Println("Hola Mundo")
	c.JSON(200, gin.H{
		"Message": "pong",
	})
}
