package model

type User struct {
	ID         int
	PhoneNum   string
	Username   string
	Sex        string
	Email      string
	Password   string
	UserType   string
	CreateTime string
}

type Login struct {
	//PhoneNum string
	Username string
	PhoneNum string
	Password string
}

type Register struct {
	ID       int
	PhoneNum string
	Username string
	Sex      string
	Email    string
	Password string
	UserType string
}
