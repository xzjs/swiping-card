package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"swiping-card/lib"
	"swiping-card/model"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type WechatLoginForm struct {
	Code string `json:"code"`
}

type WechatLoginResponse struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

func UserLogin(c *gin.Context) {
	o := WechatLoginForm{}
	if err := c.ShouldBindJSON(&o); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	client := resty.New()
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", lib.Conf().Wechat.Appid, lib.Conf().Wechat.Secret, o.Code)
	resp, err := client.R().EnableTrace().Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	wlr := WechatLoginResponse{}
	if err = json.Unmarshal(resp.Body(), &wlr); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if wlr.ErrCode != 0 {
		c.JSON(http.StatusInternalServerError, wlr.ErrMsg)
		return
	}
	var user model.User
	db := lib.DB()
	db.Where(model.User{OpenID: wlr.OpenID}).FirstOrCreate(&user)
	cookie := lib.Cookie{ID: user.ID, Type: 0}
	cookieStr, err := lib.CookieEncrypt(cookie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.SetCookie(lib.Conf().Cookie.Name, cookieStr, 3600, "/", lib.Conf().Cookie.Domain, false, true)
	c.JSON(http.StatusOK, "OK")
}

func TestLogin(c *gin.Context) {
	var user model.User
	db := lib.DB()
	db.Where(model.User{OpenID: "1"}).FirstOrCreate(&user)
	cookie := lib.Cookie{ID: user.ID, Type: 0}
	cookieStr, err := lib.CookieEncrypt(cookie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.SetCookie(lib.Conf().Cookie.Name, cookieStr, 3600, "/", lib.Conf().Cookie.Domain, false, true)
	c.JSON(http.StatusOK, "OK")
}
