package model

import "gorm.io/gorm"

type Do struct {
	gorm.Model
	PlanID   uint `json:"plan_id"`
	WayID    uint `json:"way_id"`
	Money    int  `json:"money"`
	Finished bool `json:"finished"`
}
