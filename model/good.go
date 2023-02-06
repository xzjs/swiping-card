package model

import "gorm.io/gorm"

type Good struct {
	gorm.Model
	Name        string  `json:"name"`
	Img         string  `json:"img"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}
