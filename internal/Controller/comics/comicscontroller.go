package comicscontroller

import (
	"fmt"
	"github.com/brayanlopez8/GolangComicsStore/tools"
	"github.com/gin-gonic/gin"
)

//listComics get comics
func listComics(c *gin.Contex {
	var key = tools.GenerateKeyMarvel()
	fmt.Println(key)
	// fmt.Println("Hola Mundo")
	c.JSON(200, gin.H{
		"Message": "pong",
	})
}
