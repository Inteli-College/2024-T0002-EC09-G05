package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"gorm.io/gorm"
)


func SetupRoutes(r *gin.Engine, influx api.QueryAPI, pg *gorm.DB) {

	group := r.Group("/auth")

	group.POST("/login", func(c *gin.Context) {})

	group.GET("/register", func(c *gin.Context) {})

}
