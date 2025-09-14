package controllers

import (
	"Todolist/global"
	"Todolist/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//待实现的api：查询（批量、单个）、增加、更新待办事项信息、删除、修改完成状态）

// 查询：该返回用户的所有待办事项
func GetTasks(ctx *gin.Context) {

	//获取上下文中的userid
	value, exist := ctx.Get("userid")
	if !exist {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user"})
		return
	}
	userId, ok := value.(uint)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": "can not get userid"})
		return
	}

	//查询该id对应的task
	var tasks []models.Task

	if err := global.DB.Where("user_id = ? ", userId).Find(&tasks).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//json格式返回
	ctx.JSON(http.StatusOK, &tasks)
}

// 创建新的任务
func CreatTask(ctx *gin.Context) {
	var task models.Task

	//需要ddl、title、description
	//时间格式：2006-01-02T15:04:05Z
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//获取userid 并关联
	value, exist := ctx.Get("userid")
	if !exist {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user"})
		return
	}
	userId, ok := value.(uint)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": "can not get userid"})
		return
	}
	task.UserID = userId

	//自动迁移
	if err := global.DB.AutoMigrate(&task); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	//数据库中增加记录
	if err := global.DB.Create(&task).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Creat task successfully", "task": task})

}

// 查询单个任务
func GetTaskById(ctx *gin.Context) {

	taskId := ctx.Param("id")

	var task models.Task

	if err := global.DB.Where("id = ? ", taskId).Find(&task).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, task)

}

// 更改某个任务完成情况
func ChangeCompleteStauts(ctx *gin.Context) {

	taskId := ctx.Param("id")

	//查询完成情况
	var task models.Task
	if err := global.DB.Select("completed", "id").Where("id = ?", taskId).First(&task).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	//将完成情况置反
	task.Completed = !task.Completed

	//再重新写入
	//task包含主键，自动使用主键构建查询条件来更新记录
	if err := global.DB.Model(&task).Update("completed", task.Completed).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "change successfully", "new completed status": task.Completed})
}

// 删除某个任务
func DeleteTask(ctx *gin.Context) {

	taskId := ctx.Param("id")

	//这里是软删除
	if err := global.DB.Delete(&models.Task{}, taskId).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//删除成功返回204
	ctx.Status(http.StatusNoContent)
}

// 更改task内容
func Updatetask(ctx *gin.Context) {

	//获取要更新的内容 (title,description,deadline)
	var updatedTask models.Task
	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	//获取taskid
	taskId := ctx.Param("id")
	taskIdUint, err := strconv.ParseUint(taskId, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid taskid"})
		return
	}
	updatedTask.ID = uint(taskIdUint)

	//更新 只更新updatedTask的非零字段 (title,description,deadline)
	if err := global.DB.Model(&updatedTask).Updates(updatedTask).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "update successfully", "updatedTask": updatedTask})

}
