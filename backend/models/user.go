package models

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username string `gorm:"unique;not null" binding:"required"`
	Password string `gorm:"not null;size:100" binding:"required"`

	Tasks []Task `gorm:"foreignKey:UserID"`
}

func InsertUser(db *gorm.DB, user *User) (*User, error) {
	if err := db.Create(user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, errors.New("username has been used")
		} else {
			return nil, err
		}
	}
	return user, nil
}

func SelectUserByUsername(db *gorm.DB, username string) (*User, error) {
	user := &User{}
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("unkown username, err: %v", err)
		} else {
			return nil, fmt.Errorf("select user error, err: %v", err)
		}
	}

	return user, nil

}

// 查询user是否拥有这个task
func IsOwner(db *gorm.DB, taskId uint, userId uint) (bool, error) {

	var count int64
	if err := db.Model(&Task{}).
		Where("id = ? AND user_id = ?", taskId, userId).
		Count(&count).Error; err != nil {

		return false, fmt.Errorf("select db err: %v", err)
	}

	if count == 0 {
		return false, nil
	}

	return true, nil
}
