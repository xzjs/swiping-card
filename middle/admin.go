package middle

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// IsAdmin 管理员中间件
func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := c.GetInt("type")
		if t == 1 {
			c.Next()
		} else {
			c.Abort()
			c.JSON(http.StatusForbidden, "permission forbid")
		}
	}
}
