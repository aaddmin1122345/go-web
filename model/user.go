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
	IsDelete   int
}

type Login struct {
	//PhoneNum string
	ID         int
	Username   string
	PhoneNum   string
	Email      string
	UserType   string
	Password   string
	RememberMe bool
	IsDelete   int
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

type SessionInfo struct {
	Username string
	UserType string
	UserID   int
}
