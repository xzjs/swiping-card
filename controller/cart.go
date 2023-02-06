package controller

import (
	"net/http"
	"swiping-card/lib"
	"swiping-card/model"

	"github.com/gin-gonic/gin"
)

func CartPost(c *gin.Context) {
	cart := model.Cart{}
	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	db := lib.DB()
	userID := c.GetUint("userID")
	result := db.Where("user_id = ? AND good_id = ? AND order_id IS NULL", userID, cart.GoodID).First(&cart)
	if result.Error != nil {
		cart.Num = 1
		cart.UserID = uint(userID)
		result = db.Create(&cart)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, result.Error.Error())
			return
		}
	} else {
		cart.Num++
		if result = db.Save(&cart); result.Error != nil {
			c.JSON(http.StatusInternalServerError, result.Error.Error())
			return
		}
	}
	c.JSON(http.StatusOK, "OK")
}

func CartGet(c *gin.Context) {
	var Carts []model.Cart
	db := lib.DB()
	userid := c.GetUint("userID")
	result := db.Where("user_id = ? and order_id is NULL", userid).Preload("Good").Order("created_at desc").Find(&Carts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error.Error())
	} else {
		c.JSON(http.StatusOK, Carts)
	}
}

func CartPut(c *gin.Context) {
	Cart := model.Cart{}
	if err := c.ShouldBindJSON(&Cart); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	id := c.Param("id")
	db := lib.DB()
	var CartDB model.Cart
	userID := c.GetUint("userID")
	result := db.Where("user_id = ?", userID).First(&CartDB, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error.Error())
		return
	}
	if Cart.Num == 0 {
		db.Delete(&CartDB)
	} else {
		CartDB.Num = Cart.Num
		db.Save(&CartDB)
	}
	c.JSON(http.StatusOK, "OK")
}

func CartDelete(c *gin.Context) {
	id := c.Param("id")
	db := lib.DB()
	userID := c.GetInt("userID")
	result := db.Where("user_id = ?", userID).Delete(&model.Cart{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error.Error())
		return
	}
	c.JSON(http.StatusOK, "OK")
}
