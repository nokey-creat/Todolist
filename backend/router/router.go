package router

import (
	"Todolist/common/middleware"
	"Todolist/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	//在根路由添加cors中间件，处理跨域请求
	router.Use(middleware.CoreMiddleware())

	{
		auth := router.Group("/api/v1/auth")
		auth.POST("/register", handler.Register)
		auth.POST("/login", handler.Login)

	}
	{
		tasks := router.Group("/api/v1/tasks")
		tasks.Use(middleware.CheckAuthMiddleware())

		tasks.POST("", handler.CreatTask) // 不能是"/"
		tasks.GET("", handler.GetTasks)

		tasks.GET("/:id", middleware.CheckPermissionMiddleware(), handler.GetTaskById)
		tasks.PATCH("/:id/completed", middleware.CheckPermissionMiddleware(), handler.ChangeCompleteStauts)
		tasks.DELETE("/:id", middleware.CheckPermissionMiddleware(), handler.DeleteTask)
		tasks.PATCH("/:id", middleware.CheckPermissionMiddleware(), handler.Updatetask)

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
