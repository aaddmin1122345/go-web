package service

import (
	"fmt"
	"go-web/database"
)

var (
	DbInit            = database.DbInitImpl{}
	MyDatabaseUser    = &database.MyDatabaseImpl{} //能用就别动,不使用指针数据都传输不过去
	MyDatabaseArticle = &database.ArticleImpl{}
)

func init() {
	conn, err := DbInit.Conn()
	if err != nil {
		fmt.Println("初始化数据库连接错误:", err)
		return
	}
	//DbInit.SetDb(conn)
	MyDatabaseUser.SetDb(conn)
	MyDatabaseArticle.SetDb(conn)
	//fmt.Println("service/init初始化错误")
}
