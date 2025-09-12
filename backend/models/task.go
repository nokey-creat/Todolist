package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model

	UserID string `gorm:"not null"`

	Title       string    `gorm:"not null;size:100"`
	Description string    `gorm:"size:1000"`
	Deadline    time.Time `time_format:"2006-01-02"`
	Completed   bool      `gorm:"default:false"`
}
