package controller

import (
	"net/http"
	"swiping-card/lib"
	"swiping-card/model"

	"github.com/gin-gonic/gin"
)

func WayPost(c *gin.Context) {
	way := model.Way{}
	if err := c.ShouldBindJSON(&way); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	db := lib.DB()
	result := db.Create(&way)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error.Error())
	} else {
		c.JSON(http.StatusOK, "OK")
	}
}

func WayGet(c *gin.Context) {
	var ways []model.Way
	db := lib.DB()
	result := db.Find(&ways)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error.Error())
	} else {
		c.JSON(http.StatusOK, ways)
	}
}

func WayGetOne(c *gin.Context) {
	var way model.Way
	db := lib.DB()
	id := c.Param("id")
	result := db.First(&way, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error.Error())
	} else {
		c.JSON(http.StatusOK, way)
	}
}

func WayPut(c *gin.Context) {
	way := model.Way{}
	if err := c.ShouldBindJSON(&way); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	id := c.Param("id")
	db := lib.DB()
	var wayDB model.Way
	db.First(&wayDB, id)
	wayDB.Name = way.Name
	db.Save(&wayDB)
	c.JSON(http.StatusOK, "OK")
}

func WayDelete(c *gin.Context) {
	id := c.Param("id")
	db := lib.DB()
	db.Delete(&model.Way{}, id)
	c.JSON(http.StatusOK, "OK")
}
