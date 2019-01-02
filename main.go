package main

import (
	"Golang-API-Testing/controller"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

// Main Application
func main() {
	BaseRoutes := GinEngine()

	BaseRoutes.Run(":8080")
}

// GinEngine func declaration
func GinEngine() *gin.Engine {
	// Write log into file
	logAsFile, err := os.Create("application.log")

	// Error handling
	if err != nil {
		panic(err.Error())
	}

	// Write as file
	gin.DefaultWriter = io.MultiWriter(logAsFile, os.Stdout)

	routeDefinition := gin.Default()

	versioningAPIV1 := routeDefinition.Group("/api/v1/")
	{
		versioningAPIV1.GET("books", controller.GetAllBooks)
		versioningAPIV1.GET("books/:UUIDBooks", controller.GetDetailBooks)
		versioningAPIV1.POST("books", controller.CreateNewBooks)
		versioningAPIV1.PUT("books/:UUIDBooks", controller.UpdateBooks)
		versioningAPIV1.DELETE("books/:UUIDBooks", controller.DeleteBooks)

		versioningAPIV1.GET("testing", controller.ExampleTesting)
		versioningAPIV1.GET("people", controller.GetPeople)
	}

	return routeDefinition
}
