package database

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-web/model"
)

type ArticleImpl struct {
	db *sql.DB
}

func (a *ArticleImpl) SetDb(db *sql.DB) {
	a.db = db
}

// CountArticle 分页的时候需要用到这个代码来计算有几页面
func (a *ArticleImpl) CountArticle(category string) (int, error) {
	var count int
	var err error

	if category == "" {
		query := "SELECT COUNT(*) FROM news"
		err = a.db.QueryRow(query).Scan(&count)
		if err != nil {
			return 0, err
		}
	} else {
		query := "SELECT COUNT(*) FROM news WHERE Category = ?"
		err = a.db.QueryRow(query, category).Scan(&count)
		if err != nil {
			return 0, err
		}
	}

	return count, nil
}

func (a *ArticleImpl) CreateArticle(article *model.Article) error {

	query := "INSERT INTO news (Title, Content, ImageURL, Category) VALUES (?, ?, ?, ?)"
	_, err := a.db.Exec(query, article.Title, article.Content, article.ImageURL, article.Category)
	if err != nil {
		return err
	}

	return nil
}

func (a *ArticleImpl) GetArticleByCategory(category string, page int, pageSize int) ([]*model.Article, error) {
	var err error
	var rows *sql.Rows

	// 计算偏移量,这个是获得当前页
	offset := (page - 1) * pageSize

	if category == "" {
		query := "SELECT * FROM news ORDER BY CreateTime DESC LIMIT ? OFFSET ?"
		rows, err = a.db.Query(query, pageSize, offset)
	} else {
		query := "SELECT * FROM news WHERE Category = ? ORDER BY CreateTime DESC LIMIT ? OFFSET ?"
		rows, err = a.db.Query(query, category, pageSize, offset)
	}
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
		err = rows.Scan(&article.ID, &article.Title, &article.Content, &article.CreateTime, &article.ImageURL, &article.Category, &article.Date)
		if err != nil {
			return nil, err
		}
		articles = append(articles, &article)
	}

	return articles, nil
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

func (a *ArticleImpl) GetArticleByID(id int) (*model.Article, error) {
	query := "SELECT Title, Content, CreateTime, ImageURL, Category FROM news WHERE ID = ?"
	row := a.db.QueryRow(query, id)

	var article model.Article
	err := row.Scan(&article.Title, &article.Content, &article.CreateTime, &article.ImageURL, &article.Category)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("找不到对应的文章")
		}
		return nil, fmt.Errorf("查询文章失败：%w", err)
	}

	return &article, nil
}
