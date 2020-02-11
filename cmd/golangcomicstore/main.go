package main

import (
	pingcontroller "github.com/brayanlopez8/GolangComicsStore/internal/controller/ping"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", pingcontroller.Ping)
	// r.GET("/comics", comicscontroller.)
	r.Run()
}
