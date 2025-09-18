package models

import (
	"fmt"
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

// 查询用户的所有task
func SelectTasksByUserid(db *gorm.DB, userId uint) ([]Task, error) {

	var tasks []Task

	if err := db.Where("user_id = ? ", userId).Find(&tasks).Error; err != nil {
		return nil, fmt.Errorf("select tasks err: %v", err)
	}

	return tasks, nil
}

// 向数据库中写入task
func CreatTask(db *gorm.DB, task *Task) (*Task, error) {

	if err := db.Create(task).Error; err != nil {
		return nil, fmt.Errorf("create task error: %v", err)
	}
	//creat后数据库会回填相应值
	return task, nil
}

// 查询单个task
func SelectTasksById(db *gorm.DB, taskId string) (*Task, error) {
	task := &Task{}
	err := db.Where("id = ? ", taskId).Find(task).Error
	if err != nil {
		return nil, fmt.Errorf("select task error: %v", err)
	}
	return task, nil
}

func Updatetask(db *gorm.DB, task *Task) (*Task, error) {

	//通过task主键查询要更新哪条记录
	//默认更新非零字段，不能用来更新completed
	err := db.Model(task).Updates(task).Error
	if err != nil {
		return nil, fmt.Errorf("update task error: %v", err)
	}
	return task, nil
}

func DeleteTask(db *gorm.DB, taskId string) error {

	//这里是软删除
	if err := db.Delete(&Task{}, taskId).Error; err != nil {
		return fmt.Errorf("delete task error: %v", err)
	}
	return nil
}

func UpdatetaskCompleted(db *gorm.DB, task *Task) (*Task, error) {

	//更新completed
	err := db.Model(task).Update("completed", task.Completed).Error
	if err != nil {
		return nil, fmt.Errorf("update task error: %v", err)
	}
	return task, nil
}
