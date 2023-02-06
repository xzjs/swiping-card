package controller

import (
	"net/http"
	"swiping-card/lib"
	"swiping-card/model"

	"github.com/gin-gonic/gin"
)

func BankPost(c *gin.Context) {
	bank := model.Bank{}
	if err := c.ShouldBindJSON(&bank); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	db := lib.DB()
	result := db.Create(&bank)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error.Error())
	} else {
		c.JSON(http.StatusOK, "OK")
	}
}

func BankGet(c *gin.Context) {
	var banks []model.Bank
	db := lib.DB()
	result := db.Find(&banks)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error.Error())
	} else {
		c.JSON(http.StatusOK, banks)
	}
}
