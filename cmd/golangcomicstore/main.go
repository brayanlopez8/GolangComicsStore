package main

import (
	ping "github.com/brayanlopez8/GolangComicsStore/internal/controller/ping"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", ping.Ping)
	r.GET("/comics", comics.listComics)
	r.Run()
}
