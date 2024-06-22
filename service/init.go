package service

import (
	"fmt"
	"go-web/database"
	"go-web/model"
)

var (
	DbInit            = &database.DbInitImpl{}
	MyDatabaseUser    = &database.MyDatabaseImpl{} //能用就别动,不使用指针sql连接数据都传输不过去
	MyDatabaseArticle = &database.ArticleImpl{}
	MyDatabaseComment = &database.CommentImpl{}

	// var CommentApi = &api.CommentImpl{}
)

type Article interface {
	GetArticleByKeyword(keyword string) ([]*model.Article, error)
	CreateArticle(article *model.Article) error
	GetArticleByCategory(category string, page int, pageSize int) ([]*model.Article, error)
	GetArticleByID(id int) (*model.Article, error)
	CountArticle(category string) (int, error)
	DeleteArticle(id int) error
	editArticleByID(id, authorID int)
}

type Comment interface {
	AddComment(comment *model.Comment) error
	GetComments(article int) ([]*model.Comment, error)
	DeleteComment(id int) error
}

type UserService interface {
	AddUser(user *model.Register) error
	GetUserByKeyword(username string) ([]*model.User, error)
	Login(login *model.Login) (*model.Login, error)
	GetUserByPhoneNum(PhoneNum string) *model.User
	DeleteUser(StudID int) error
	UpdateUser(user *model.Register) error
}

func init() {
	conn, err := DbInit.Conn()
	if err != nil {
		fmt.Println("初始化数据库连接错误:", err)
		return
	}
	//DbInit.SetDb(conn)
	MyDatabaseUser.SetDb(conn)
	MyDatabaseArticle.SetDb(conn)
	MyDatabaseComment.SetDb(conn)
	//fmt.Println("service/init初始化错误")
}
