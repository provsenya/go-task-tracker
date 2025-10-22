package main

import (
	"github.com/gin-gonic/gin"
	"github.com/provsenya/go-task-tracker/routes"
)

func main() {
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
