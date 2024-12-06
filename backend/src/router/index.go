package router

import (
	"task-manager/src/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	Router := gin.Default()
	Router.GET("/", controller.Greetings)
	Router.Run("localhost:80")
}
