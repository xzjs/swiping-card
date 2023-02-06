package controller

import (
	"net/http"
	"swiping-card/lib"
	"swiping-card/model"

	"github.com/gin-gonic/gin"
)

func CyclePost(c *gin.Context) {
	cycle := model.Cycle{}
	if err := c.ShouldBindJSON(&cycle); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	db := lib.DB()
	result := db.Create(&cycle)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error.Error())
	} else {
		c.JSON(http.StatusOK, "OK")
	}
}

func CycleGet(c *gin.Context) {
	var cycles []model.Cycle
	db := lib.DB()
	result := db.Find(&cycles)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error.Error())
	} else {
		c.JSON(http.StatusOK, cycles)
	}
}

func CycleGetOne(c *gin.Context) {
	var cycle model.Cycle
	db := lib.DB()
	id := c.Param("id")
	result := db.First(&cycle, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error.Error())
	} else {
		c.JSON(http.StatusOK, cycle)
	}
}

func CyclePut(c *gin.Context) {
	cycle := model.Cycle{}
	if err := c.ShouldBindJSON(&cycle); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	id := c.Param("id")
	db := lib.DB()
	var cycleDB model.Cycle
	db.First(&cycleDB, id)
	cycleDB.Name = cycle.Name
	db.Save(&cycleDB)
	c.JSON(http.StatusOK, "OK")
}

func CycleDelete(c *gin.Context) {
	id := c.Param("id")
	db := lib.DB()
	db.Delete(&model.Cycle{}, id)
	c.JSON(http.StatusOK, "OK")
}
