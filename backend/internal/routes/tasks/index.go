package tasks_routes

import (
	tasks_controller "task-manager/internal/controllers/tasks"
	tasks_repository "task-manager/internal/repository"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterTasksRoutes(router *gin.Engine, client *mongo.Client) {
	tasks := router.Group("/api/tasks/")
	repository := tasks_repository.NewTaskRepository(client)
	controller := tasks_controller.NewTaskController(repository)
	{
		tasks.GET("/", controller.GetAllTasksController)
		tasks.POST("/", controller.CreateTaskController)
		tasks.DELETE("/:id", controller.DeleteTaskController)
	}
}
