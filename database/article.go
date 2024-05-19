package database

import (
	"database/sql"
	"go-web/model"
)

type Article interface {
	GetArticleByKeyword(keyword string) ([]*model.Article, error)
}

type ArticleImpl struct {
	Db *sql.DB
}

// GetArticleByKeyword 模糊查询,主页右上角用
func (a *ArticleImpl) GetArticleByKeyword(keyword string) ([]*model.Article, error) {
	query := "SELECT * FROM news WHERE Title LIKE ? OR Content LIKE ? OR PublishDate LIKE ? OR ImageURL LIKE ? OR Category LIKE ? ;"
	rows, err := a.Db.Query(query, "%"+keyword+"%")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			// handle error
		}
	}(rows)

	var articles []*model.Article
	for rows.Next() {
		var article model.Article
		err = rows.Scan(&article.ID, &article.Title, &article.Content, &article.Published, &article.ImageURL, &article.Category)
		if err != nil {
			return nil, err
		}
		articles = append(articles, &article)
	}

	return articles, nil
}
