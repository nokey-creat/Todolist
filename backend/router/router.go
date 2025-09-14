package router

import (
	"Todolist/config"
	"Todolist/controllers"
	"Todolist/middleware"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	//在根路由添加core中间件，处理跨域请求
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{config.AppConfig.CORSConfig.AllowOrigins},  //允许所有来源
		AllowMethods:     []string{"GET", "POST", "DELETE", "PATCH"},          //允许的方法
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"}, //除了简单请求头外，允许的头部
		ExposeHeaders:    []string{"Content-Length"},                          //允许js获取的响应头部（不需要Authorization，这是请求头中的）
		AllowCredentials: true,                                                // 使用*时不能设置为true
		MaxAge:           12 * time.Hour,                                      //本次预检请求的有效期
	}))

	{
		auth := router.Group("/api/auth")
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)

	}

	{
		tasks := router.Group("/api/tasks")
		tasks.Use(middleware.CheckAuthMiddleware())

		tasks.POST("", controllers.CreatTask)
		tasks.GET("", controllers.GetTasks)

		tasks.GET("/:id", middleware.CheckPermissionMiddleware(), controllers.GetTaskById)
		tasks.PATCH("/:id/completed", middleware.CheckPermissionMiddleware(), controllers.ChangeCompleteStauts)
		tasks.DELETE("/:id", middleware.CheckPermissionMiddleware(), controllers.DeleteTask)
		tasks.PATCH("/:id", middleware.CheckPermissionMiddleware(), controllers.Updatetask)

	}

	// 自定义 404 处理
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":  "API endpoint not found",
			"status": 404,
		})
	})

	return router
}
