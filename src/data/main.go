package data

type User struct{
	UserName string
	PassWord string
	ID int
}

type Database struct{
	Userdata map[string]User
	UserCnt int
}

var Data Database

func AddUser(username string, password string){
	Data.UserCnt++
	NewUser := new(User)
	NewUser.UserName = username
	NewUser.PassWord = password
	NewUser.ID = Data.UserCnt
	Data.Userdata[username] = *NewUser
}
