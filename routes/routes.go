package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/provsenya/go-task-tracker/database"
	"github.com/provsenya/go-task-tracker/handlers"
)

func SetupRoutes(r *gin.Engine) {
	database.InitDB()

	r.GET("/tasks", handlers.GetTasks)
	r.POST("/tasks", handlers.CreateTask)
}
