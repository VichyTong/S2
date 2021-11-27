package controller

import (
	"Server/model"
	"errors"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	ErrorCookieMissing = errors.New("cookies missing")
)

func MakeCookie(name string, value string) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = value
	cookie.Expires = time.Now().Add(24 * 10 * time.Hour)
	return cookie
}

func FrontPage(c echo.Context) error {
	r, err := http.Get("https://github.com")
	if err != nil {
		return c.JSON(http.StatusForbidden, nil)
	}
	s, err := ioutil.ReadAll(r.Body)
	return c.JSON(http.StatusOK, s)
}

func Register(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	err := model.UserRegister(username, password)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "registered successfully")
}

func Session(c echo.Context) error {
	username := c.FormValue("login")
	password := c.FormValue("password")

	sessionID, err := model.UserCheck(username, password)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	c.SetCookie(MakeCookie("user_session", sessionID))
	c.SetCookie(MakeCookie("logged_in", "yes"))
	c.SetCookie(MakeCookie("dotcom_user", username))

	return c.JSON(http.StatusFound, nil)
}
func Login(c echo.Context) error {
	loggedIn, err := c.Cookie("logged_in")
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorCookieMissing.Error())
	}
	if loggedIn.Value != "yes" {
		return c.JSON(http.StatusOK, nil)
	}
	username, err := c.Cookie("dotcom_user")
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorCookieMissing.Error())
	}
	session, err := c.Cookie("user_session")
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorCookieMissing.Error())
	}

	err = model.SessionCheck(username.Value, session.Value)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	c.Response().Header().Add("location", "http://github.com")
	return c.JSON(http.StatusFound, nil)
}
