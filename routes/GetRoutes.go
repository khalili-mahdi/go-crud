package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/khalili-mahdi/go-crud/controllers"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/todos", controllers.GetAllTodoItems)
		api.GET("/todos/:id", controllers.GetTodoItem)
		api.POST("/todos", controllers.CreateTodoItem)
	}
}
