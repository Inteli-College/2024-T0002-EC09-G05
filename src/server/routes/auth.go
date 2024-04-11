package routes

import (
	"g5/server/auth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupAuthRoutes(r *gin.Engine, pg *gorm.DB) {

	group := r.Group("/auth")

	group.POST("/login", func(c *gin.Context) { auth.Login(c, pg) })

	group.POST("/register", func(c *gin.Context) { auth.Register(c, pg) })

	group.POST("/changeUserRole", func(c *gin.Context) { auth.ChangeRole(c, pg) })

	group.POST("/changeUserDirectorate", func(c *gin.Context) { auth.ChangeDirectorate(c, pg) })

}
