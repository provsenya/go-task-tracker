package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/provsenya/go-task-tracker/database"
	"github.com/provsenya/go-task-tracker/models"
	"net/http"
)

func GetTasks(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, title, completed FROM tasks")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var t models.Task
		err := rows.Scan(&t.ID, &t.Title, &t.Completed)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tasks = append(tasks, t)
	}
	c.JSON(http.StatusOK, tasks)
}

func CreateTask(c *gin.Context) {
	var t models.Task
	if err := c.BindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := database.DB.Exec("INSERT INTO tasks (title) VALUES (?)", t.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "task added"})
}
