package main

import (
	database "task-manager/internal/config"
	"task-manager/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	c := database.Connect()
	routes.RegisterRoutes(r, c)

	r.Run()
}
