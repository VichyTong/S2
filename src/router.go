package main

import (
	"Server/controller"
	"github.com/labstack/echo/v4"
)

func RouterInit(e *echo.Echo) {
	e.GET("/", controller.FrontPage)
	e.POST("usr/register", controller.Register)
	e.GET("usr/login", controller.Login)
	e.POST("usr/session", controller.Session)
	e.POST("user/update", controller.Update)
}
