package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
	Phone   string `json:"phone" binding:"required"`
	Remark  string `json:"remark"`
	Carts   []Cart `json:"carts"`
	UserID  uint   `json:"-"`
	Status  int    `json:"status"`
	Total   int    `json:"total"`
	Num     int    `json:"num"`
}
