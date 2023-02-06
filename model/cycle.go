package model

import "gorm.io/gorm"

type Cycle struct {
	gorm.Model
	Name string `json:"name"`
}
