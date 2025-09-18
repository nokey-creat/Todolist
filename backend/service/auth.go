package service

import (
	"Todolist/common/utils"
	"Todolist/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 处理用户注册逻辑,有错误向handler返回
func Register(c *gin.Context, userReq *models.User) error {

	//先加密密码
	EncryptedPwd, err := utils.GetEncryptedPassword(userReq.Password)
	if err != nil {
		return fmt.Errorf("Register user error : %v", err)
	}

	//写入数据库
	user, err := models.InsertUser(models.GetDB(), &models.User{
		Username: userReq.Username,
		Password: EncryptedPwd,
	})
	if err != nil {
		return fmt.Errorf("Register user error : %v", err)

	}

	c.JSON(http.StatusCreated, gin.H{"message": "register successfully", "user_id": user.ID})
	return nil
}

// 用户登录
func Login(c *gin.Context, userReq *models.User) error {

	//查找用户信息
	user, err := models.SelectUserByUsername(models.GetDB(), userReq.Username)
	if err != nil {
		//服务器错误导致登录失败，返回err
		return fmt.Errorf("user login err: %v", err)
	}

	//验证密码
	isPwdOK := utils.CheckPassword(user.Password, userReq.Password)
	if !isPwdOK {
		//到这里是完成了登录（密码错误），返回响应并返回nil
		c.JSON(http.StatusOK, gin.H{"message": "wrong password"})
		return nil
	}

	//密码正确，返回JWT
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		//服务器错误导致登录失败，返回err
		return fmt.Errorf("user login err: %v", err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "login successfully", "jwt": token})

	return nil
}
