package controller

import (
	"net/http"
	"swiping-card/lib"
	"swiping-card/model"
	"time"

	"github.com/gin-gonic/gin"
)

func DoPost(c *gin.Context) {
	do := model.Do{}
	if err := c.ShouldBindJSON(&do); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	db := lib.DB()
	result := db.Create(&do)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error.Error())
	} else {
		c.JSON(http.StatusOK, "OK")
	}
}

func DoGet(c *gin.Context) {
	var dos []model.Do
	db := lib.DB()
	var plans []model.Plan
	result := db.Preload("Cycle").Find(&plans)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error.Error())
	}
	for _, v := range plans {
		start, end := getTimeSection(v.Cycle)
		c.JSON(http.StatusOK, gin.H{"start": start, "end": end})
	}
	result = db.Find(&dos)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error.Error())
	} else {
		c.JSON(http.StatusOK, dos)
	}
}

func DoGetOne(c *gin.Context) {
	var do model.Do
	db := lib.DB()
	id := c.Param("id")
	result := db.First(&do, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error.Error())
	} else {
		c.JSON(http.StatusOK, do)
	}
}

func DoPut(c *gin.Context) {
	do := model.Do{}
	if err := c.ShouldBindJSON(&do); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	id := c.Param("id")
	db := lib.DB()
	var doDB model.Do
	db.First(&doDB, id)
	// doDB.Name = do.Name
	// doDB.Img = do.Img
	// doDB.Price = do.Price
	// doDB.Description = do.Description
	// db.Save(&doDB)
	c.JSON(http.StatusOK, "OK")
}

func DoDelete(c *gin.Context) {
	id := c.Param("id")
	db := lib.DB()
	db.Delete(&model.Do{}, id)
	c.JSON(http.StatusOK, "OK")
}

func getTimeSection(cycle model.Cycle) (start time.Time, end time.Time) {
	switch cycle.Name {
	case "周":
		now := time.Now()
		offset := int(time.Monday - now.Weekday())
		if offset > 0 {
			offset = -6
		}
		start = now.AddDate(0, 0, offset)
		end = now.AddDate(0, 0, 6+offset)
	}
	return start, end
}
