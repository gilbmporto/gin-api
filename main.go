package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Task struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

var tasks []Task = []Task{
	{Id: 1, Title: "Surfar na beira mar"},
	{Id: 2, Title: "Conseguir um emprego"},
}

func main() {
	// Created a new Gin router
	router := gin.Default()

	router.SetTrustedProxies(nil)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})

	router.GET("/tasks", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, tasks)
	})

	router.GET("/tasks/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		// Convert the string ID to an integer
		idInt, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		// Find the task with the given ID and return it
		for _, task := range tasks {
			if task.Id == (idInt) {
				ctx.JSON(http.StatusOK, task)
				return
			}
		}

		// If the task is not found, return a 404 error
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
	})

	router.POST("/tasks", func(ctx *gin.Context) {
		var newTask *Task
		if err := ctx.ShouldBindJSON(&newTask); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newTask.Id = len(tasks) + 1
		tasks = append(tasks, *newTask)
		ctx.JSON(http.StatusCreated, *newTask)
	})

	router.Run(":3000")
}
