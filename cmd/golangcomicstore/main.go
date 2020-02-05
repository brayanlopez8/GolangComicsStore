package main

import (
	pingController "github.com/brayanlopez8/GolangComicsStore/internal/controller/ping"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", pingController.Ping)

}
