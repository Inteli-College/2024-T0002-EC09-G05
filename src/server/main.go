package main

import (
	"g5/server/db"
	"g5/server/routes"
	"g5/server/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	influx := db.CreateinfluxConnection()

	pg := db.SetupPostgres()

	routes.SetupDataRoutes(r, influx, pg)

	routes.SetupAuthRoutes(r, pg)

	r.Use(cors.New(cors.Config{
		AllowOrigins:           []string{"http://localhost:5173", "http://localhost:3000"},
		AllowBrowserExtensions: true,
		AllowHeaders:           []string{"Content-Type", "application/json","Access-Control-Allow-Headers", "Access-Control-Allow-Origin", },
		AllowCredentials:       true,
	}))

	go utils.PrintASCIIArt()

	r.Run(":8080")
}
