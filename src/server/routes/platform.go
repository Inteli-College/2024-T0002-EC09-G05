package routes

import (
	"g5/server/platform"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupPlatformRoutes(r *gin.Engine, pg *gorm.DB) {

	group := r.Group("/platform")

	group.POST("/getComponents", func(c *gin.Context) { platform.GetAll(c, pg) })

	group.GET("/get_all_sources", func(c *gin.Context) { platform.GetAllComponents(c, pg) }) //get_all_components

	//group.POST("/updateLayout", func(c *gin.Context) { platform.ChangeRole(c, pg) })

}
