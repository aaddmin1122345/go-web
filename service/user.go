package service

import (
	"go-web/model"
)

//var dbInit = DbInit
//var MyDatabaseUser = MyDatabaseUser

type UserServiceImpl struct {
}

func (u *UserServiceImpl) UpdateUser(user *model.Register) error {

	err := MyDatabaseUser.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil

}

func (u *UserServiceImpl) GetUserByKeyword(username string) ([]*model.User, error) {
	users, err := MyDatabaseUser.GetUserByKeyword(username)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// 必须初始化数据库连接
//func init() {
//	conn, err := DbInit.Conn()
//	if err != nil {
//		fmt.Println("初始化数据库连接错误")
//		return
//	}
//	MyDatabaseUser.Db = conn // 将连接对象保存到 MyDatabaseUser 结构体中的 db 字段
//}

//func (u UserServiceImpl) init() {
//	conn, err := dbInit.Conn()
//	if err != nil {
//		fmt.Println("初始化数据库连接错误")
//		return
//	}
//	MyDatabaseUser.Db = conn // 将连接对象保存到 MyDatabaseUser 结构体中的 db 字段
//}

func (u *UserServiceImpl) Login(login *model.Login) (*model.Login, error) {
	user, err := MyDatabaseUser.Login(login)
	//fmt.Println(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserServiceImpl) GetUserByPhoneNum(PhoneNum string) *model.User {
	phoneNum, err := MyDatabaseUser.GetUserByPhoneNum(PhoneNum)
	if err != nil {
		//fmt.Println("检测StudID是否正确")
		return nil
	}
	return phoneNum
}

func (u *UserServiceImpl) AddUser(user *model.Register) error {
	err := MyDatabaseUser.AddUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserServiceImpl) DeleteUser(StudID int) error {
	err := MyDatabaseUser.DeleteUser(StudID)
	if err != nil {
		//fmt.Println("检测StudID是否正确")
		return err
	}
	return nil
}
