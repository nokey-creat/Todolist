package middleware

import (
	"Todolist/common/config"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CoreMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{config.AppConfig.CORSConfig.AllowOrigins},  //允许所有来源
		AllowMethods:     []string{"GET", "POST", "DELETE", "PATCH"},          //允许的方法
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"}, //除了简单请求头外，允许的头部
		ExposeHeaders:    []string{"Content-Length"},                          //允许js获取的响应头部（不需要Authorization，这是请求头中的）
		AllowCredentials: true,                                                // 使用*时不能设置为true
		MaxAge:           12 * time.Hour,                                      //本次预检请求的有效期
	})
}
