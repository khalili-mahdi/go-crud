package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khalili-mahdi/go-crud/dto"
	"github.com/khalili-mahdi/go-crud/initializers"
	"github.com/khalili-mahdi/go-crud/models"
)

func CreateTodoItem(c *gin.Context) {
	var todoItemDto dto.TodoDto
	if bindError := c.Bind(&todoItemDto); bindError != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": bindError.Error(),
			},
		)
		return
	}
	todoItem := models.TodoItem{Title: todoItemDto.Title, Body: todoItemDto.Body}
	result := initializers.DB.Create(&todoItem)
	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	if &todoItem == nil {
		c.Status(
			http.StatusNoContent,
		)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"data": todoItem,
	})
}

func GetAllTodoItems(c *gin.Context) {
	var todoItems []models.TodoItem
	response := initializers.DB.Find(&todoItems)
	if response.Error != nil {
		c.Status(
			http.StatusInternalServerError,
		)
		return
	}
	if &todoItems == nil {
		c.Status(
			http.StatusNoContent,
		)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": todoItems,
	})
}

func GetTodoItem(c *gin.Context) {
	var todoItem models.TodoItem
	id := c.Param("id")
	response := initializers.DB.First(&todoItem, id)
	if response.Error != nil {
		c.Status(
			http.StatusInternalServerError,
		)
		return
	}
	if &todoItem == nil {
		c.Status(
			http.StatusNoContent,
		)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": todoItem,
	})
}
