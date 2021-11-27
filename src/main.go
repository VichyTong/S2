package main

import (
	"Server/model"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	RouterInit(e)
	model.Init()
	e.Logger.Fatal(e.Start(":8000"))
}
