package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go-web/model"
)

type DbInit interface {
	Conn() (*sql.DB, error)
	Close() error
}

type Article interface {
	GetArticleByKeyword(keyword string) ([]*model.Article, error)
	CreateArticle(article *model.Article) error
	SetDb(db *sql.DB)
	GetArticleByCategory(category string, page int, pageSize int) ([]*model.Article, error)
	GetArticleByID(id int) (*model.Article, error)
	CountArticle(category string) (int, error)
}

type Comment interface {
	AddComment(comment *model.Comment) error
	GetComments(articleID int) ([]*model.Comment, error)
}

type Database interface {
	HashPassword(password string) ([]byte, error)
	CheckUser(user *model.Register) error
	CheckDelUser(id int) error
	GetUserByKeyword(username string) ([]*model.User, error)
	GetUserByPhoneNum(phoneNum string) (*model.User, error)
	AddUser(user *model.Register) error
	UpdateUser(user *model.Register) error
	DeleteUser(id int) error
	Login(login *model.Login) (*model.Login, error)
	SetDb(db *sql.DB)
}

type DbInitImpl struct {
	Db *sql.DB
}

//func (d DbInitImpl) SetDb(db *sql.DB) {
//	d.Db = db
//}

func (d *DbInitImpl) Conn() (*sql.DB, error) {
	dbConn := "root:123456@tcp(127.0.0.1:3306)/web"
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		_ = d.Close()
		return nil, err
	}
	d.Db = db
	return d.Db, nil
}

func (d *DbInitImpl) Close() error {
	if d.Db != nil {
		return d.Db.Close()
	}
	return nil
}
