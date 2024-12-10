package tasks_controller

import (
	"net/http"
	task_model "task-manager/internal/models"

	"github.com/gin-gonic/gin"
)

func (tc *TasksController) CreateTaskController(c *gin.Context) {
	var body task_model.Task

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTask, err := tc.Repo.Create(c.Request.Context(), body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdTask)
}
