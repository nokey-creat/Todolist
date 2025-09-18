package handler

import (
	"Todolist/models"
	"Todolist/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 用户注册
func Register(c *gin.Context) {
	// 处理请求，调用业务处理逻辑

	//获取请求内容
	var userReq models.User
	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Register failed", "error": err.Error()})
		return
	}

	//处理业务
	if err := service.Register(c, &userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Register failed", "error": err.Error()})
		return
	}

}

func Login(c *gin.Context) {

	//获取请求内容
	userReq := &models.User{}

	if err := c.ShouldBindJSON(userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Login failed", "error": err.Error()})
		return
	}

	//调用登录的业务逻辑,传入请求内容
	err := service.Login(c, userReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Login failed", "error": err.Error()})
		return
	}
}
