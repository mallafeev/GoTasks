package main

import (
	"task9/internal/db"
	"task9/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()

	r := gin.Default()

	r.POST("/groups", handlers.CreateGroup)
	r.GET("/groups", handlers.GetGroups)
	r.PUT("/groups/:id", handlers.UpdateGroup)
	r.DELETE("/groups/:id", handlers.DeleteGroup)
	r.POST("/people", handlers.CreatePerson)
	r.PUT("/people/:id", handlers.UpdatePerson)
	r.DELETE("/people/:id", handlers.DeletePerson)
	r.GET("/groups/:id/people", handlers.GetPeopleByGroup)

	r.Run(":8080")
}
