package main

import (
	"Server/controller"
	"github.com/labstack/echo/v4"
)

func RouterInit(e *echo.Echo) {
	e.GET("/", controller.FrontPage)
	e.POST("/register", controller.Register)
	e.GET("/login", controller.Login)
	e.POST("/session", controller.Session)
}
