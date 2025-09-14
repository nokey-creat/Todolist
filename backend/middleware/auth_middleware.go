package middleware

import (
	"Todolist/global"
	"Todolist/models"
	"Todolist/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 检查是否通过身份验证
func CheckAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("authorization")
		if tokenString == "" {
			//fmt.Println("tokenstring empty")
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			ctx.Abort()
			return
		}

		token := strings.TrimPrefix(tokenString, "Bearer ")

		userid, err := utils.ParseJWT(token)
		if err != nil {
			//fmt.Println("parse wrong")
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token", "error": err.Error()})
			ctx.Abort()
			return
		}

		ctx.Set("userid", userid)

		ctx.Next()

	}
}

// 检查用户是否有权限访问该task
func CheckPermissionMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取任务id
		taskId := ctx.Param("id")

		//获取userid
		value, exist := ctx.Get("userid")
		if !exist {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user"})
			return
		}
		userId, ok := value.(uint)
		if !ok {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "can not get userid"})
			return
		}

		//查询任务并验证所有权
		var count int64
		if err := global.DB.Model(&models.Task{}).
			Where("id = ? AND user_id = ?", taskId, userId).
			Count(&count).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		if count == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Permission deny or task does not exist"})
			ctx.Abort()
			return
		}

		ctx.Next()

	}
}
