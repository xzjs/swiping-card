package controller

import (
	"net/http"
	"swiping-card/lib"

	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func TokenGet(c *gin.Context) {
	conf := lib.Conf()
	putPolicy := storage.PutPolicy{
		Scope: conf.QiNiu.Bucket,
	}
	mac := qbox.NewMac(conf.QiNiu.AK, conf.QiNiu.SK)
	upToken := putPolicy.UploadToken(mac)
	c.JSON(http.StatusOK, gin.H{"token": upToken})
}
