package main

import (
	"Golang-API-Testing/controller"
	"github.com/gin-gonic/gin"
)

// Main Application
func main() {
	BaseRoutes := GinEngine()

	BaseRoutes.Run(":8000")
}

// GinEngine func declaration
func GinEngine() *gin.Engine {
	routeDefinition := gin.Default()

	versioningAPIV1 := routeDefinition.Group("/api/v1/")
	{
		versioningAPIV1.GET("product", controller.GetAllProducts)
		versioningAPIV1.GET("product/:UUIDProducts", controller.GetDetailProducts)
	}

	return routeDefinition
}
