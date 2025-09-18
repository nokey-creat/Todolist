package service

import (
	"Todolist/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetTasks(ctx *gin.Context, userId uint) error {

	//查询该id对应的task
	tasks, err := models.SelectTasksByUserid(models.GetDB(), userId)
	if err != nil {
		return fmt.Errorf("get tasks err: %v", err)
	}

	//json格式返回
	ctx.JSON(http.StatusOK, &tasks)
	return nil
}

func CreatTask(c *gin.Context, userId uint, taskReq *models.Task) error {

	//关联user和task
	task, err := models.CreatTask(models.GetDB(), &models.Task{
		UserID:      userId,
		Title:       taskReq.Title,
		Description: taskReq.Description,
		Deadline:    taskReq.Deadline,
	})
	if err != nil {
		return fmt.Errorf("creat task error: %v", err)
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Creat task successfully", "task": task}) //task自动解引用？
	return nil
}

func GetTaskById(c *gin.Context, taskId string) error {

	task, err := models.SelectTasksById(models.GetDB(), taskId)
	if err != nil {
		return fmt.Errorf("get task error: %v", err)
	}
	c.JSON(http.StatusOK, task)
	return nil
}

func ChangeCompleteStauts(c *gin.Context, taskId string) error {
	//查询完成情况
	task, err := models.SelectTasksById(models.GetDB(), taskId)
	if err != nil {
		return fmt.Errorf("get task error: %v", err)
	}
	//置反
	task.Completed = !task.Completed

	//更新数据库
	_, err = models.UpdatetaskCompleted(models.GetDB(), task)
	if err != nil {
		return fmt.Errorf("update task error: %v", err)
	}

	//返回响应
	c.JSON(http.StatusCreated, gin.H{"message": "change successfully", "new completed status": task.Completed})
	return nil

}

func Updatetask(c *gin.Context, taskId uint, taskReq *models.Task) error {

	//更新task
	updatedTask, err := models.Updatetask(models.GetDB(), &models.Task{
		Model: gorm.Model{
			ID: taskId,
		},
		Title:       taskReq.Title,
		Description: taskReq.Description,
		Deadline:    taskReq.Deadline,
	})

	if err != nil {
		return fmt.Errorf("update task error: %v", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "update successfully",
		"updatedTask": updatedTask})

	return nil
}

func DeleteTask(c *gin.Context, taskId string) error {

	err := models.DeleteTask(models.GetDB(), taskId)
	if err != nil {
		return fmt.Errorf("delete task error: %v", err)
	}

	//删除成功，返回响应
	c.Status(http.StatusNoContent)
	return nil
}
