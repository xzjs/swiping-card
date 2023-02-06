package model

import (
	"gorm.io/gorm"
)

type Way struct {
	gorm.Model
	Name string `json:"name"`
}
