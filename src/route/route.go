package route

import (
	"A/controller"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func RegisterRoute() *gin.Engine {
	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	router.POST("/github.com/session", controller.SessionHandler)
	router.GET("/github.com/login", controller.LoginHandler)
	router.GET("/github.com", controller.GithubHandler)

	return router
}
