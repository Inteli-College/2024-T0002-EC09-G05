package main

import (
	"g5/server/db"
	"g5/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"net/http"
)

type GetDataByRelativeTimeBody struct {
	StartTime  int      `json:"start_time"`
	EndTime    *int     `json:"end_time,omitempty"`
	Aggregator int      `json:"aggregator"`
	Field      []string `json:"fields"`
	Sensor     []string `json:"sensors"`
}

func SetupRouter(queryAPI api.QueryAPI) *gin.Engine {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Online!")
	})

	r.POST("/getDataByRelativeTime", func(c *gin.Context) {

		var req GetDataByRelativeTimeBody

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res, err := db.GetDataByRelativeTime(
			queryAPI,
			req.StartTime,
			req.EndTime,
			req.Aggregator,
			req.Field,
			req.Sensor,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, (res))
	})

	r.GET("/getAllSensors", func(c *gin.Context) {

		res, err := db.GetAllSensors(queryAPI)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, (res))

	})

	r.GET("/getAllFields", func(c *gin.Context) {

		res, err := db.GetAllFields(queryAPI)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, (res))

	})

	return r
}

func main() {

	queryAPI := db.CreateinfluxConnection()

	r := SetupRouter(queryAPI)

	r.Use(cors.Default())

	go utils.PrintASCIIArt()

	r.Run(":8080")
}
