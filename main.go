package main

import (
	"awesomeGolang/controller"

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
	}

	return routeDefinition
}
