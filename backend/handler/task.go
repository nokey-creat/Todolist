package handler

import (
	"Todolist/models"
	"Todolist/service"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//待实现的api：查询（批量、单个）、增加、更新待办事项信息、删除、修改完成状态）

// 查询：该返回用户的所有待办事项
func GetTasks(c *gin.Context) {

	//获取上下文中的userid
	value, exist := c.Get("userid")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user"})
		return
	}
	userId, ok := value.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "can not get userid"})
		return
	}

	err := service.GetTasks(c, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

}

// 创建新的任务
func CreatTask(c *gin.Context) {

	//从请求上下文获取ddl、title、description
	//时间格式：2006-01-02T15:04:05Z
	var taskReq models.Task
	if err := c.ShouldBindJSON(&taskReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//从请求上下文获取userid
	value, exist := c.Get("userid")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user"})
		return
	}
	userId, ok := value.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "can not get userid"})
		return
	}

	//creatTask业务
	err := service.CreatTask(c, userId, &taskReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

}

// 查询单个任务
func GetTaskById(c *gin.Context) {

	taskId := c.Param("id")

	//getTask业务
	err := service.GetTaskById(c, taskId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

}

// 更改某个任务完成情况
func ChangeCompleteStauts(c *gin.Context) {
	//处理请求
	taskId := c.Param("id")

	//ChangeCompleteStauts业务
	err := service.ChangeCompleteStauts(c, taskId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

// 删除某个任务
func DeleteTask(c *gin.Context) {

	taskId := c.Param("id")

	err := service.DeleteTask(c, taskId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

// 更改task的内容(title,description,deadline)
func Updatetask(c *gin.Context) {

	//获取要更新的内容 (title,description,deadline)
	taskReq := &models.Task{}
	if err := c.ShouldBindJSON(taskReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	//获取taskid(uint64)
	taskId := c.Param("id")
	taskIdUint64, err := strconv.ParseUint(taskId, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid taskid"})
		return
	}

	//更新
	err = service.Updatetask(c, uint(taskIdUint64), taskReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
