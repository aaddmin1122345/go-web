package api

import "errors"

type userRepository interface {
	addUser(user *User) (*User, error)
	deleteUser(user *User) (*User, error)
	updateUser(user *User) (*User, error)
	getUserByUsername(username string) (*User, error)
}

type User struct {
	ID       int
	StudID   string
	Username string
	Sex      string
	email    string
}

type userServiceImpl struct {
}

func (u userServiceImpl) addUser(user *User) (*User, error) {
	//if (*user).StudID != '0' {
	return nil, errors.New("user already exists")
	//}
}

func (u userServiceImpl) deleteUser(user *User) (*User, error) {
	return nil, errors.New("not implemented")

}

func (u userServiceImpl) updateUser(user *User) (*User, error) {
	return nil, errors.New("not implemented")
}

func (u userServiceImpl) getUserByUsername(username string) (*User, error) {
	return nil, errors.New("not implemented")
}
