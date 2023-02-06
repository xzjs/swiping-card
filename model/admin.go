package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}
