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

	pg, err := db.SetupPostgres()

	if err != nil {
		panic("Erro no postgress")
	}

	routes.SetupDataRoutes(r, influx, pg)
	routes.SetupAuthRoutes(r, pg)
	routes.SetupPlatformRoutes(r, pg)

	r.Use(cors.New(cors.Config{
		AllowOrigins:           []string{"http://localhost:5173"},
		AllowBrowserExtensions: true,
		AllowHeaders:           []string{"application/json"},
		AllowCredentials:       true,
	}))

	go utils.PrintASCIIArt()

	r.Run(":8080")
}
