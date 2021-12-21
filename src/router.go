package main

import (
	"Server/controller"
	"github.com/labstack/echo/v4"
)

func RouterInit(e *echo.Echo) {
	e.POST("usr/register", controller.Register)
	e.GET("usr/login", controller.Login)
	e.POST("usr/session", controller.Session)
	e.POST("usr/update", controller.Update)
	e.GET("issues", controller.Issue)
	e.POST(":username/:repository/issues", controller.IssueAdd)
	e.POST(":username/:repository/issues/:number", controller.IssueModify)
	e.POST(":username/:repository/issue_comments", controller.IssueAddComment)
}
