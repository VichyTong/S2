package main

import (
	"A/route"
	"github.com/gin-gonic/gin"
)
func main() {
	gin.SetMode(gin.ReleaseMode)
	router := route.RegisterRoute()
	router.Run(":8000")
}
