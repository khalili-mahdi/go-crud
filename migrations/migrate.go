package main

import (
	"github.com/khalili-mahdi/go-crud/initializers"
	"github.com/khalili-mahdi/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.TodoItem{})
}
