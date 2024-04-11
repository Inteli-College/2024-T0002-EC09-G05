package platform

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCookieHandler(c *gin.Context) {
	cookie, err := c.Cookie("token")
	if err != nil {
		c.String(http.StatusNotFound, "Cookie not found")
		return
	}
	c.String(http.StatusOK, "Cookie value: %s", cookie)
}
