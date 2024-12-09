package routes

import (
	tasks_routes "task-manager/internal/routes/tasks"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterRoutes(router *gin.Engine, client *mongo.Client) {
	tasks_routes.RegisterTasksRoutes(router, client)
}
