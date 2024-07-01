package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/khalili-mahdi/go-crud/dto"
	"github.com/khalili-mahdi/go-crud/initializers"
	"github.com/khalili-mahdi/go-crud/models"
)

func CreateTodoItem(c *gin.Context) {
	var todoItem models.TodoItem
	c.Bind(todoItem)
	result := initializers.DB.Create(&todoItem)
	if result.Error != nil {
		log.Fatal("Creating entity in database failed!")
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"todoItem": todoItem,
	})
}

func GetAllTodoItems(c *gin.Context) {
	var todoItems []dto.TodoDto
	initializers.DB.Find(&todoItems)
	c.JSON(200, gin.H{
		"todoItems": todoItems,
	})
}

func GetTodoItem(c *gin.Context) {

}
