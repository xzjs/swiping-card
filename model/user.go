package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	OpenID     string `json:"-"`
	SessionKey string `json:"-"`
}
