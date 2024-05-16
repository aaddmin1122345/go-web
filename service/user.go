package service

import (
	"fmt"
	"go-web/database"
	"go-web/model"
)

var dbInit = database.DbInitImpl{}
var myDatabase = database.MyDatabaseImpl{}

func GetUsersByStudID(StudID int) *model.User {
	id, err := myDatabase.GetUserByStudID(StudID)
	if err != nil {
		fmt.Println("检测StudID是否正确")
		return nil
	}
	return id
}

func init() {
	conn, err := dbInit.Conn()
	if err != nil {
		fmt.Println("初始化数据库连接错误")
		return
	}
	myDatabase.Db = conn // 将连接对象保存到 myDatabase 结构体中的 db 字段
}
