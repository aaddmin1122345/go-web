package service

import (
	"fmt"
	"go-web/database"
	"go-web/model"
)

var dbInit = database.DbInitImpl{}
var myDatabase = database.MyDatabaseImpl{}

type UserService interface {
	AddUser(user *model.Register) error
	GetUserByUserName(username string) ([]*model.User, error)
	Login(login *model.Login) (*model.Login, error)
	GetUsersByStudID(StudID int) *model.User
	DeleteUser(StudID int) error
	UpdateUser(user *model.Register) error
}

type UserServiceImpl struct{}

func (u *UserServiceImpl) UpdateUser(user *model.Register) error {
	err := myDatabase.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil

}

func (u *UserServiceImpl) GetUserByUserName(username string) ([]*model.User, error) {
	users, err := myDatabase.GetUserByUserName(username)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// 必须初始化数据库连接
func init() {
	conn, err := dbInit.Conn()
	if err != nil {
		fmt.Println("初始化数据库连接错误")
		return
	}
	myDatabase.Db = conn // 将连接对象保存到 myDatabase 结构体中的 db 字段
}

//func (u UserServiceImpl) init() {
//	conn, err := dbInit.Conn()
//	if err != nil {
//		fmt.Println("初始化数据库连接错误")
//		return
//	}
//	myDatabase.Db = conn // 将连接对象保存到 myDatabase 结构体中的 db 字段
//}

func (u *UserServiceImpl) Login(login *model.Login) (*model.Login, error) {
	user, err := myDatabase.Login(login)
	//fmt.Println(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserServiceImpl) GetUsersByStudID(StudID int) *model.User {
	id, err := myDatabase.GetUserByStudID(StudID)
	if err != nil {
		//fmt.Println("检测StudID是否正确")
		return nil
	}
	return id
}

func (u *UserServiceImpl) AddUser(user *model.Register) error {
	err := myDatabase.AddUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserServiceImpl) DeleteUser(StudID int) error {
	err := myDatabase.DeleteUser(StudID)
	if err != nil {
		//fmt.Println("检测StudID是否正确")
		return err
	}
	return nil
}
