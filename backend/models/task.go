package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model

	UserID uint `gorm:"not null"`

	Title       string    `gorm:"not null;size:100" binding:"required"`
	Description string    `gorm:"size:1000" binding:"required"`
	Deadline    time.Time `binding:"required"`
	Completed   bool      `gorm:"default:false"`
}
