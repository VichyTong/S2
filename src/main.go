package main

import (
	echo "github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	RouterInit(e)
	e.Logger.Fatal(e.Start(":8000"))
}
