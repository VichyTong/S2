package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Register struct {
	username string `form:"username"`
	password string `form:"password" validator"`
}

type Login struct {
	username string `form:"username"`
	password string `form:"password"`
}
func main() {

	r := gin.Default()

	var Usermap map[string]string
	Usermap = make(map[string]string)
	Usermap["root"] = "admin"

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello!")
		return
	})
	r.POST("/register", func(c *gin.Context){
		var form Register
		form.username = c.PostForm("username")
		form.password = c.PostForm("password")
		if err := c.Bind(&form); err != nil{
			c.String(http.StatusBadRequest, "")
			return
		}
		_ , ok := Usermap[form.username]
		if ok {
			c.String(http.StatusOK, "Username had been used.")
			return
		} else {
			Usermap[form.username] = form.password
			c.String(http.StatusOK, "Registered successfully.")
			return
		}
	})
	r.POST("/login", func(c *gin.Context) {
		var form Login
		form.username = c.PostForm("username")
		form.password = c.PostForm("password")
		if err := c.Bind(&form); err != nil{
			c.String(http.StatusBadRequest, "")
			return
		}
		password, ok := Usermap[form.username]
		if ok {
			if password == form.password {
				//fmt.Println(password)
				//fmt.Println(form.password)
				c.String(http.StatusOK, "Login successfully.")
				return
			} else{
				c.String(http.StatusOK, "Login failed.")
				return
			}
		} else{
			c.String(http.StatusOK, "Username had not been registered.")
			return
		}
	})
	r.Run(":8000")
}
