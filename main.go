package main

import (
	"github.com/gin-gonic/gin"
	"github.com/khalili-mahdi/go-crud/controllers"
	"github.com/khalili-mahdi/go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()
	r.GET("/todos", controllers.GetAllTodoItems)
	r.GET("/todos/{id}", controllers.GetTodoItem)
	r.POST("/todos", controllers.CreateTodoItem)
	r.Run()
}
