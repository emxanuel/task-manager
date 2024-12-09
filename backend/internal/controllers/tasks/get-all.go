package tasks_controller

import (
	"context"
	"net/http"
	tasks_repository "task-manager/internal/repository"

	"github.com/gin-gonic/gin"
)

type TasksController struct {
	Repo *tasks_repository.Repository
}

func NewTaskController(repo *tasks_repository.Repository) *TasksController {
	return &TasksController{Repo: repo}
}

func (tc *TasksController) GetAllTasksController(c *gin.Context) {
	tasks, err := tc.Repo.FindAll(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}
