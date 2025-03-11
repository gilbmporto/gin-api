package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Created a new Gin router
	router := gin.Default()

	// Initialize the database and create the table if it doesn't exist.
	DB = initDB("tasks.db")
	defer DB.Close()

	router.SetTrustedProxies(nil)

	RegisterRoutes(router)

	router.Run(":3000")
}
