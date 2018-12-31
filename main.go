package main

import (
	"Golang-API-Testing/controller"
	"github.com/gin-gonic/gin"
)

// Main Application
func main() {
	BaseRoutes := GinEngine()

	BaseRoutes.Run(":8080")
}

// GinEngine func declaration
func GinEngine() *gin.Engine {
	routeDefinition := gin.Default()

	versioningAPIV1 := routeDefinition.Group("/api/v1/")
	{
		versioningAPIV1.GET("books", controller.GetAllBooks)
		versioningAPIV1.GET("books/:UUIDBooks", controller.GetDetailBooks)
		versioningAPIV1.POST("books", controller.CreateNewBooks)
		versioningAPIV1.PUT("books/:UUIDBooks", controller.UpdateBooks)
		versioningAPIV1.DELETE("books/:UUIDBooks", controller.DeleteBooks)
	}

	return routeDefinition
}
