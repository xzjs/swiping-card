package controller

import (
	"net/http"
	"strconv"
	"swiping-card/lib"
	"swiping-card/model"

	"github.com/gin-gonic/gin"
)

func OrderPost(c *gin.Context) {
	order := model.Order{}
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	db := lib.DB()
	userID := c.GetUint("userID")
	order.UserID = userID
	order.Status = 0
	var cartIDs []uint
	for _, value := range order.Carts {
		cartIDs = append(cartIDs, value.ID)
	}
	var carts []model.Cart
	db.Where("user_id = ?", userID).Preload("Good").Find(&carts, cartIDs)
	for _, value := range carts {
		order.Num += value.Num
		order.Total += value.Num * int(value.Good.Price)
	}
	result := db.Save(&order)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error.Error())
		return
	}
	c.JSON(http.StatusOK, "OK")
}

func OrderGet(c *gin.Context) {
	var Orders []model.Order
	db := lib.DB()
	userid := c.GetUint("userID")
	status, err := strconv.Atoi(c.DefaultQuery("status", "0"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	result := db.Where("user_id = ? AND status = ?", userid, status).Preload("Carts.Good").Order("created_at desc").Find(&Orders)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error.Error())
	} else {
		c.JSON(http.StatusOK, Orders)
	}
}

func OrderPut(c *gin.Context) {
	Order := model.Order{}
	if err := c.ShouldBindJSON(&Order); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	id := c.Param("id")
	db := lib.DB()
	var OrderDB model.Order
	userID := c.GetUint("userID")
	result := db.Where("user_id = ?", userID).First(&OrderDB, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error.Error())
		return
	}
	OrderDB.Status = Order.Status
	result = db.Save(&OrderDB)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error.Error())
		return
	}
	c.JSON(http.StatusOK, "OK")
}

func OrderDelete(c *gin.Context) {
	id := c.Param("id")
	db := lib.DB()
	userID := c.GetInt("userID")
	result := db.Where("user_id = ?", userID).Delete(&model.Order{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error.Error())
		return
	}
	c.JSON(http.StatusOK, "OK")
}
