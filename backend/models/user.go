package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Username string `gorm:"unique;not null" binding:"required"`
	Password string `gorm:"not null;size:100" binding:"required"`

	Tasks []Task `gorm:"foreignKey:UserID"`
}
