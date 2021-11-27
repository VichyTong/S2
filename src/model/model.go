package model

import (
	"Server/util"
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	ErrorExist         = errors.New("username used")
	ErrorNoUser        = errors.New("no such user")
	ErrorWrongPassword = errors.New("wrong password")
	ErrorSession       = errors.New("not logged in")
	ErrorDatabase      = errors.New("database error")
)

type info struct {
	Username string
	Password string
	Session  string
}

var c *mgo.Collection

func Init() {
	url := "mongodb://localhost:27017"
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	defer session.Clone()
	session.SetMode(mgo.Monotonic, true)
	c = session.DB("test").C("info")
}

func UserRegister(username string, password string) error {
	inf := info{}
	err := c.Find(bson.M{"username": username}).One(&inf)
	if err == nil {
		return ErrorExist
	}
	err = c.Insert(&info{username, password, ""})
	if err != nil {
		return ErrorDatabase
	}
	return nil
}

func UserCheck(username string, password string) (string, error) {
	inf := info{}
	err := c.Find(bson.M{"username": username}).One(&inf)
	if err != nil {
		return "", ErrorNoUser
	}
	if password != inf.Password {
		return "", ErrorWrongPassword
	}
	sessionID := util.RandomString(64)
	err = c.Update(bson.M{"username": username}, bson.D{{"$set", bson.M{"session": sessionID}}})
	return sessionID, nil
}

func SessionCheck(username string, sessionID string) error {
	inf := info{}
	err := c.Find(bson.M{"username": username}).One(&inf)
	if err != nil {
		return ErrorSession
	}
	if inf.Session != sessionID {
		return ErrorSession
	}
	return nil
}
