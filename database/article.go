package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go-web/model"
)

type Article interface {
	GetArticleByKeyword(keyword string) ([]*model.Article, error)
	CreateArticle(article *model.Article) error
	SetDb(db *sql.DB)
}

type ArticleImpl struct {
	db *sql.DB
}

func (a *ArticleImpl) SetDb(db *sql.DB) {
	a.db = db
}

func (a *ArticleImpl) CreateArticle(article *model.Article) error {

	query := "INSERT INTO news (Title, Content, ImageURL, Category) VALUES (?, ?, ?, ?)"
	_, err := a.db.Exec(query, article.Title, article.Content, article.ImageURL, article.Category)
	if err != nil {
		return err
	}

	return nil
}

// GetArticleByKeyword 模糊查询,主页右上角用
func (a *ArticleImpl) GetArticleByKeyword(keyword string) ([]*model.Article, error) {
	query := "SELECT ID, Title, Content, CreateTime, ImageURL, Category FROM news WHERE Title LIKE ? OR Content LIKE ? OR CreateTime LIKE ? OR ImageURL LIKE ? OR Category LIKE ?"
	rows, err := a.db.Query(query, "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = rows.Close(); err != nil {
			// 处理关闭 rows 时的错误
		}
	}()

	var articles []*model.Article
	for rows.Next() {
		var article model.Article
		err = rows.Scan(&article.ID, &article.Title, &article.Content, &article.CreateTime, &article.ImageURL, &article.Category)
		if err != nil {
			return nil, err
		}
		articles = append(articles, &article)
	}

	return articles, nil
}
