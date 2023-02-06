package model

import "gorm.io/gorm"

type Card struct {
	gorm.Model
	Name   string `json:"name"`
	UserID uint   `json:"-"`
	BankID uint   `json:"bank_id"`
	NO     string `json:"no"`
}
