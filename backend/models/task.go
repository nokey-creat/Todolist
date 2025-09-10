package models

import "gorm.io/gorm"

type List struct {
	gorm.Model
	UserId      uint //外键
	Title       string
	Description string
	Completed   bool
}
