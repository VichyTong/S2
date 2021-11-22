package model

import (
	"Server/util"
	"errors"
)

var (
	ErrorExist         = errors.New("username used")
	ErrorNoUser        = errors.New("no such user")
	ErrorWrongPassword = errors.New("wrong password")
	ErrorSession       = errors.New("not logged in")
)
var usernameMap map[string]string
var sessionMap map[string]string

func init() {
	usernameMap = make(map[string]string)
	sessionMap = make(map[string]string)
}

func UserRegister(username string, password string) error {
	_, ok := usernameMap[username]
	if ok {
		return ErrorExist
	}
	usernameMap[username] = password
	return nil
}

func UserCheck(username string, password string) (string, error) {
	Password, ok := usernameMap[username]
	if !ok {
		return "", ErrorNoUser
	}
	if Password != password {
		return "", ErrorWrongPassword
	}
	sessionID := util.RandomString(64)
	sessionMap[username] = sessionID
	return sessionID, nil
}

func SessionCheck(username string, sessionID string) error {
	SessionID, ok := sessionMap[username]
	if !ok {
		return ErrorSession
	}
	if sessionID != SessionID {
		return ErrorSession
	}
	return nil
}
