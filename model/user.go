package model

type User struct {
	ID       int
	StudID   string
	Username string
	Sex      string
	Email    string
}

type Login struct {
	Username string
	Password string
}

type Register struct {
	ID       int
	StudID   string
	Username string
	Sex      string
	Email    string
	Password string
}
