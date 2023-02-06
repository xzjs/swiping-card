package model

import (
	"time"

	"gorm.io/gorm"
)

type Plan struct {
	gorm.Model
	UserID    uint      `json:"-"`
	CardID    uint      `json:"card_id"`
	Sum       int       `json:"sum"` //总金额
	CycleID   uint      `json:"cycle_id"`
	Cycle     Cycle     `json:"cycle"`
	Total     int       `json:"total"` //总次数
	Frequency int       `json:"frequency"`
	Floor     int       `json:"floor"` //下限
	Ways      []Way     `json:"ways" gorm:"many2many:plan_ways;"`
	Start     time.Time `json:"start"`
	End       time.Time `json:"end"`
}
