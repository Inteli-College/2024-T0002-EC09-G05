package main

import (
	"g5/server/db"
	"net/http"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return r
}

func main() {
	writeAPI := db.SetupInflux()

	writeAPI.EnableBatching()

	r := SetupRouter()

	r.Run(":8080")
}
