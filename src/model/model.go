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

type Info struct {
	Username               string
	Password               string
	Session                string
	ProfileName            string
	ProfileBio             string
	ProfileBlog            string
	ProfileTwitterUsername string
	ProfileCompany         string
	ProfileLocation        string
	Type                   string
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

func UserRegister(form Info) error {
	inf := Info{}
	err := c.Find(bson.M{"username": form.Username}).One(&inf)
	if err == nil {
		return ErrorExist
	}
	form.Type = "profile"
	err = c.Insert(&form)
	if err != nil {
		return ErrorDatabase
	}
	return nil
}

func UserCheck(username string, password string) (string, error) {
	inf := Info{}
	err := c.Find(bson.M{"username": username}).One(&inf)
	if err != nil {
		return "", ErrorNoUser
	}
	if password != inf.Password {
		return "", ErrorWrongPassword
	}
	sessionID := util.RandomString(64)
	err = c.Update(bson.M{"username": username}, bson.D{{"$set", bson.M{"session": sessionID}}})
	if err != nil {
		return "", ErrorDatabase
	}
	return sessionID, nil
}

func SessionCheck(username string, sessionID string) error {
	inf := Info{}
	err := c.Find(bson.M{"username": username}).One(&inf)
	if err != nil {
		return ErrorSession
	}
	if inf.Session != sessionID {
		return ErrorSession
	}
	return nil
}

func UserUpdate(form Info) error {
	username := form.Username
	inf := Info{}
	err := c.Find(bson.M{"username": username}).One(&inf)
	if err != nil {
		return ErrorNoUser
	}
	if form.Password != "" {
		err = c.Update(bson.M{"username": username}, bson.D{{"$set", bson.M{"password": form.Password}}})
		if err != nil {
			return ErrorDatabase
		}
	}
	if form.ProfileName != "" {
		err = c.Update(bson.M{"username": username}, bson.D{{"$set", bson.M{"profilename": form.ProfileName}}})
		if err != nil {
			return ErrorDatabase
		}
	}
	if form.ProfileBio != "" {
		err = c.Update(bson.M{"username": username}, bson.D{{"$set", bson.M{"profilebio": form.ProfileBio}}})
		if err != nil {
			return ErrorDatabase
		}
	}
	if form.ProfileBlog != "" {
		err = c.Update(bson.M{"username": username}, bson.D{{"$set", bson.M{"profileblog": form.ProfileBlog}}})
		if err != nil {
			return ErrorDatabase
		}
	}
	if form.ProfileTwitterUsername != "" {
		err = c.Update(bson.M{"username": username}, bson.D{{"$set", bson.M{"profiletwitterusername": form.ProfileTwitterUsername}}})
		if err != nil {
			return ErrorDatabase
		}
	}
	if form.ProfileCompany != "" {
		err = c.Update(bson.M{"username": username}, bson.D{{"$set", bson.M{"profilecompany": form.ProfileCompany}}})
		if err != nil {
			return ErrorDatabase
		}
	}
	if form.ProfileLocation != "" {
		err = c.Update(bson.M{"username": username}, bson.D{{"$set", bson.M{"profilelocation": form.ProfileLocation}}})
		if err != nil {
			return ErrorDatabase
		}
	}
	return nil
}

type issue struct {
	Id          string
	Body        string
	Open        bool
	CommentList map[int]string
	Url         string
	Type        string
}

func IssueList(username string) (map[int]issue, error) {

}
