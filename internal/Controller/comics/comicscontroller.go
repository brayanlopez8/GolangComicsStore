package comicsController

import (
	"net/http"

	usesCase "github.com/brayanlopez8/GolangComicsStore/internal/usesCase"
	"github.com/gin-gonic/gin"
)

//Comic contains injected interface from usesCase layer.
type Comic struct {
	UseCase usesCase.ComicUseCase
}

// SetGameEndpoints sets endpoints for Game entity.
func SetGameEndpoints(version *gin.RouterGroup, c usesCase.ComicUseCase) {
	comics := &Comic{
		UseCase: c,
	}

	endpoints := version.Group("/comics")
	{
		endpoints.GET("", comics.findAll)
		// endpoints.GET("/:id", game.getByID)
		// endpoints.POST("", game.post)
		// endpoints.PATCH("", game.patch)
		// endpoints.DELETE("/:id", game.deleteByID)
	}
}
func (co *Comic) findAll(c *gin.Context) {
	comics, _ := co.UseCase.FindAll()

	c.JSON(
		http.StatusOK,
		comics,
	)
}

//ListComics get comics
// func ListComics(c *gin.Context) {
// 	var key = tools.GenerateKeyMarvel()
// 	fmt.Println(key)
// 	// fmt.Println("Hola Mundo")
// 	c.JSON(200, gin.H{
// 		"Message": key,
// 	})
// }
