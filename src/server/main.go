package main

import (
	"fmt"
	"net/http"
	"g5/server/db"
	"g5/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type GetDataByRelativeTimeBody struct {
	StartTime  int      `json:"start_time"`
	EndTime    *int     `json:"end_time,omitempty"`
	Aggregator int      `json:"aggregator"`
	Field      []string `json:"field"`
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

		fmt.Println(req.EndTime)

		res, err := db.GetDataByRelativeTime(
			queryAPI,
			req.StartTime,
			req.EndTime,
			req.Aggregator,
			req.Field,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, (res))
	})

	r.GET("/GetDataByTimestamp", func(c *gin.Context) {
	})

	r.GET("/getAllSensors", func(c *gin.Context) {

		res, err := db.GetAllSensors(queryAPI)

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

	go utils.PrintASCIIArt()
	
	r.Run(":8080")
}
