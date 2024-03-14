package routes

import (
	"g5/server/db"
	"github.com/gin-gonic/gin"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"gorm.io/gorm"
	"net/http"
)

type GetDataByRelativeTimeDTO struct {
	StartTime  int      `json:"start_time"`
	EndTime    *int     `json:"end_time,omitempty"`
	Aggregator int      `json:"aggregator"`
	Field      []string `json:"fields"`
	Sensor     []string `json:"sensors"`
}

func SetupRoutes(r *gin.Engine,influx  pg *gorm.DB) *gin.Engine {

	group := r.Group("/")

	group.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Online!")
	})

	group.POST("/getDataByRelativeTime", func(c *gin.Context) {

		var req GetDataByRelativeTimeDTO

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res, err := db.GetDataByRelativeTime(
			influx,
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

	group.GET("/getAllSensors", func(c *gin.Context) {

		res, err := db.GetAllSensors(influx)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, (res))

	})

	group.GET("/getAllFields", func(c *gin.Context) {

		res, err := db.GetAllFields(influx)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, (res))

	})

	return r
}

