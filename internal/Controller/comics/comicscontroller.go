package comicscontroller

import (
	"fmt"

	"github.com/brayanlopez8/GolangComicsStore/tools"
	"github.com/gin-gonic/gin"
)

//ListComics get comics
func ListComics(c *gin.Context) {
	var key = tools.GenerateKeyMarvel()
	fmt.Println(key)
	// fmt.Println("Hola Mundo")
	c.JSON(200, gin.H{
		"Message": key,
	})
}
