package main

import (
	"NetLikePlate/controllers"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
)

func Start() {
	//InitializeLogging("logs/log.log")
	startup()
}

// Route NOTE: Define Route Here
func Route(e *echo.Echo) {
	group := e.Group("/test")
	group.Use(BasicMiddleware)
	group.GET("/testing", controllers.HelloWorld)
}
