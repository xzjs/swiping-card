package model

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	Good    Good  `json:"good"`
	GoodID  uint  `json:"good_id"`
	Num     int   `json:"num"`
	UserID  uint  `json:"-"`
	OrderID *uint `json:"order_id"`
}
