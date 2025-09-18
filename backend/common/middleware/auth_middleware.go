package middleware

import (
	"Todolist/common/utils"
	"Todolist/models"
	"net/http"
	"strconv"
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
		taskIdStr := ctx.Param("id")
		taskId, err := strconv.ParseUint(taskIdStr, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "get taskId error"})
			ctx.Abort()
			return
		}

		//获取userid
		value, exist := ctx.Get("userid")
		if !exist {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user"})
			ctx.Abort()
			return
		}
		userId, ok := value.(uint)
		if !ok {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "can not get userid"})
			ctx.Abort()
			return
		}

		//验证所有权
		isOwner, err := models.IsOwner(models.GetDB(), uint(taskId), userId)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "can not vervified userid"})
			ctx.Abort()
			return
		}

		//无权访问
		if !isOwner {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Permission deny or task does not exist"})
			ctx.Abort()
			return
		}

		ctx.Next()

	}
}
