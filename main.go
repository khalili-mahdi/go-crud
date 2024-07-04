package main

import (
	"github.com/gin-gonic/gin"
	"github.com/khalili-mahdi/go-crud/initializers"
	"github.com/khalili-mahdi/go-crud/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run()
}
