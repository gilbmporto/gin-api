package main

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", RouteTest)

	router.GET("/tasks", GetAllTasks)

	router.GET("/tasks/:id", GetTaskById)

	router.POST("/tasks", CreateTask)

	router.PATCH("/tasks/:id", UpdateTask)

	router.DELETE("/tasks/:id", DeleteTask)
}
