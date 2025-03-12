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

func RouteTest(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Running API Test"})
}

func GetAllTasks(ctx *gin.Context) {
	rows, err := DB.Query("SELECT id, title FROM tasks")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.Id, &task.Title)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tasks = append(tasks, task)
	}

	// If no tasks are found, return a 204 No Content status
	if len(tasks) == 0 {
		ctx.JSON(http.StatusNoContent, gin.H{"message": "No tasks found"})
		return
	}

	// Return all tasks as JSON
	ctx.JSON(http.StatusOK, tasks)
}

func GetTaskById(ctx *gin.Context) {
	id := ctx.Param("id")

	var task Task
	row := DB.QueryRow("SELECT id, title FROM tasks WHERE id = ?", id)

	err := row.Scan(&task.Id, &task.Title)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func CreateTask(ctx *gin.Context) {
	// Bind the JSON request body to a Task struct
	var newTask Task
	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert the new task into the tasks slice
	tasks = append(tasks, newTask)

	// Insert the new task into the database
	stmt, err := DB.Prepare("INSERT INTO tasks (title) VALUES (?)")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer stmt.Close()

	res, err := stmt.Exec(newTask.Title)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newTask.Id = int(id)
	ctx.JSON(http.StatusCreated, newTask)
}

func UpdateTask(ctx *gin.Context) {
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
}

func DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")

	_, err := DB.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
