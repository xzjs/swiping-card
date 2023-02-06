package controller

import (
	"net/http"
	"swiping-card/lib"
	"swiping-card/model"

	"github.com/gin-gonic/gin"
)

func AdminLoginPost(c *gin.Context) {
	admin := model.Admin{}
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	db := lib.DB()
	adminDB := model.Admin{}
	db.Where("name=?", admin.Name).First(&adminDB)
	pwd := lib.MD5(admin.Pwd)
	if adminDB.Pwd == pwd {
		cookie := lib.Cookie{ID: adminDB.ID, Type: 1}
		cookieStr, err := lib.CookieEncrypt(cookie)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		} else {
			c.SetCookie(lib.Conf().Cookie.Name, cookieStr, 3600, "/", lib.Conf().Cookie.Domain, false, true)
			c.JSON(http.StatusOK, "OK")
		}
	} else {
		c.JSON(http.StatusBadRequest, "用户名或密码错误")
	}
}
