package controller

import (
	"net/http"
	"swiping-card/lib"
	"swiping-card/model"

	"github.com/gin-gonic/gin"
)

func PlanPost(c *gin.Context) {
	plan := model.Plan{}
	if err := c.ShouldBindJSON(&plan); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	userID := c.GetUint("userID")
	plan.UserID = userID
	db := lib.DB()
	result := db.Create(&plan)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error.Error())
	} else {
		c.JSON(http.StatusOK, "OK")
	}
}

func PlanGet(c *gin.Context) {
	var plans []model.Plan
	db := lib.DB()
	result := db.Preload("Ways").Find(&plans)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error.Error())
	} else {
		c.JSON(http.StatusOK, plans)
	}
}

func PlanGetOne(c *gin.Context) {
	var plan model.Plan
	db := lib.DB()
	id := c.Param("id")
	result := db.First(&plan, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error.Error())
	} else {
		c.JSON(http.StatusOK, plan)
	}
}

func PlanPut(c *gin.Context) {
	plan := model.Plan{}
	if err := c.ShouldBindJSON(&plan); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	id := c.Param("id")
	db := lib.DB()
	var planDB model.Plan
	db.First(&planDB, id)
	planDB.CardID = plan.CardID
	planDB.Sum = plan.Sum
	planDB.CycleID = plan.CycleID
	planDB.Total = plan.Total
	db.Save(&planDB)
	c.JSON(http.StatusOK, "OK")
}

func PlanDelete(c *gin.Context) {
	id := c.Param("id")
	db := lib.DB()
	db.Delete(&model.Plan{}, id)
	c.JSON(http.StatusOK, "OK")
}
