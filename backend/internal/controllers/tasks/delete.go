package tasks_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (tc *TasksController) DeleteTaskController(c *gin.Context) {
	id, success := c.Params.Get("id")
	if !success {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing id"})
		return
	}

	response, err := tc.Repo.Delete(c.Request.Context(), id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": response})
}
