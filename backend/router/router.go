package router

import (
	"Todolist/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	{
		auth := router.Group("/api/auth")
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)

	}

	return router
}
