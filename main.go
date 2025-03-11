package main

import (
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
	DB := initDB("tasks.db")
	defer DB.Close()

	router.SetTrustedProxies(nil)

	RegisterRoutes(router)

	router.Run(":3000")
}
