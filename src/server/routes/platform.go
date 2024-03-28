package routes

import (
	"g5/server/platform"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupPlatformRoutes(r *gin.Engine, pg *gorm.DB) {

	group := r.Group("/platform")

	group.POST("/getComponents", func(c *gin.Context) { platform.GetAll(c, pg) })

	//group.POST("/updateLayout", func(c *gin.Context) { platform.ChangeRole(c, pg) })

}
