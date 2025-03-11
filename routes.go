package main

import (
	"net/http"
	"slices"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
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
		newTask.Id = tasks[len(tasks)-1].Id + 1
		tasks = append(tasks, *newTask)
		ctx.JSON(http.StatusCreated, *newTask)
	})

	router.PATCH("/tasks/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		// Convert the string ID to an integer
		idInt, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		var updatedTask *Task
		if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Find the task with the given ID and update it
		for index, task := range tasks {
			if task.Id == (idInt) {
				tasks[index].Title = updatedTask.Title
				ctx.JSON(http.StatusOK, tasks[index])
				return
			}
		}

		// If the task is not found, return a 404 error
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
	})

	router.DELETE("/tasks/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		// Convert the string ID to an integer
		idInt, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		// Find the task with the given ID and remove it
		for index, task := range tasks {
			if task.Id == (idInt) {
				tasks = slices.Delete(tasks, index, index+1)
				ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
				return
			}
		}

		// If the task is not found, return a 404 error
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
	})
}
