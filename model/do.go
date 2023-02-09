package model

import "gorm.io/gorm"

type Do struct {
	gorm.Model
	UserID   uint `json:"-"`
	PlanID   uint `json:"plan_id"`
	Plan     Plan `json:"plan"`
	WayID    uint `json:"way_id"`
	Way      Way  `json:"way"`
	Money    int  `json:"money"`
	Finished bool `json:"finished"`
}
