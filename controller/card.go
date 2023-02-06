package controller

import (
	"net/http"
	"swiping-card/lib"
	"swiping-card/model"

	"github.com/gin-gonic/gin"
)

func CardPost(c *gin.Context) {
	card := model.Card{}
	if err := c.ShouldBindJSON(&card); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	userID := c.GetUint("userID")
	card.UserID = userID
	db := lib.DB()
	result := db.Create(&card)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error.Error())
	} else {
		c.JSON(http.StatusOK, "OK")
	}
}

func CardGet(c *gin.Context) {
	var cards []model.Card
	db := lib.DB()
	userID := c.GetUint("userID")
	result := db.Where("user_id = ?", userID).Find(&cards)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error.Error())
	} else {
		c.JSON(http.StatusOK, cards)
	}
}
