package middle

import (
	"net/http"
	"swiping-card/lib"

	"github.com/gin-gonic/gin"
)

// IsLogin 登录态验证中间件
func IsLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		_cookie, err := c.Cookie(lib.Conf().Cookie.Name)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusUnauthorized, "token is not valid")
			return
		}
		cookie, err := lib.CookieDecrypt(_cookie)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusUnauthorized, "token is not valid")
			return
		}
		c.Set("userID", cookie.ID)
		c.Set("type", cookie.Type)

		c.Next()
	}
}
