package controller

import (
	"net/http"
	"swiping-card/lib"
	"swiping-card/model"

	"github.com/gin-gonic/gin"
)

func GoodPost(c *gin.Context) {
	good := model.Good{}
	if err := c.ShouldBindJSON(&good); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	db := lib.DB()
	result := db.Create(&good)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error.Error())
	} else {
		c.JSON(http.StatusOK, "OK")
	}
}

func GoodGet(c *gin.Context) {
	var goods []model.Good
	db := lib.DB()
	result := db.Find(&goods)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error.Error())
	} else {
		c.JSON(http.StatusOK, goods)
	}
}

func GoodGetOne(c *gin.Context) {
	var good model.Good
	db := lib.DB()
	id := c.Param("id")
	result := db.First(&good, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error.Error())
	} else {
		c.JSON(http.StatusOK, good)
	}
}

func GoodPut(c *gin.Context) {
	good := model.Good{}
	if err := c.ShouldBindJSON(&good); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	id := c.Param("id")
	db := lib.DB()
	var goodDB model.Good
	db.First(&goodDB, id)
	goodDB.Name = good.Name
	goodDB.Img = good.Img
	goodDB.Price = good.Price
	goodDB.Description = good.Description
	db.Save(&goodDB)
	c.JSON(http.StatusOK, "OK")
}

func GoodDelete(c *gin.Context) {
	id := c.Param("id")
	db := lib.DB()
	db.Delete(&model.Good{}, id)
	c.JSON(http.StatusOK, "OK")
}
