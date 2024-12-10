package main

import (
	database "task-manager/internal/config"
	"task-manager/internal/middlewares"
	"task-manager/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middlewares.CorsMiddleware())

	c := database.Connect()
	routes.RegisterRoutes(r, c)

	r.Run()
}
