package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Task struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func main() {
	// Created a new Gin router
	router := gin.Default()

	router.SetTrustedProxies(nil)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})

	router.Run(":3000")
}
