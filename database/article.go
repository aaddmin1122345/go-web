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
func (a *ArticleImpl) CountArticle(keyword string, authorID int) (int, error) {
	var count int
	var err error

	if authorID == 1 && keyword == "" {
		// 如果作者ID为1且关键词为空，则统计所有文章数量
		query := "SELECT COUNT(*) FROM news"
		err = a.db.QueryRow(query).Scan(&count)
		if err != nil {
			return 0, err
		}
	} else {
		// 否则根据关键词和作者ID来统计文章数量
		query := "SELECT COUNT(*) FROM news WHERE"
		var args []interface{}
		if authorID != 0 {
			query += " AuthorID = ?"
			args = append(args, authorID)
		}
		if keyword != "" {
			if len(args) > 0 {
				query += " AND"
			}
			query += " Title LIKE ?"
			args = append(args, "%"+keyword+"%")
		}

		err = a.db.QueryRow(query, args...).Scan(&count)
		if err != nil {
			return 0, err
		}
	}

	return count, nil
}

func (a *ArticleImpl) CountArticleALL(category string) (int, error) {
	var count int
	var err error

	if category != "" {
		query := "SELECT COUNT(*) FROM news where Category = ?"
		err = a.db.QueryRow(query, category).Scan(&count)
		if err != nil {
			return 0, err
		}
	}
	return count, nil
}

func (a *ArticleImpl) CreateArticle(article *model.Article) error {

	query := "INSERT INTO news (Title, Content, ImageURL, Category,Author, AuthorID,File) VALUES (?,?,?, ?, ?, ?,?)"
	_, err := a.db.Exec(query, article.Title, article.Content, article.ImageURL, article.Category, article.Author, article.AuthorID, article.File)
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
		query := "SELECT * FROM news WHERE IsDelete = 0   ORDER BY CreateTime DESC LIMIT ? OFFSET ? "
		rows, err = a.db.Query(query, pageSize, offset)
	} else {
		query := "SELECT * FROM news WHERE Category = ? AND IsDelete = 0  ORDER BY CreateTime DESC LIMIT ? OFFSET ? "
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
		err = rows.Scan(&article.ID, &article.Title, &article.Content, &article.CreateTime, &article.ImageURL, &article.Category, &article.Date, &article.IsDelete, &article.AuthorID, &article.Author, &article.File)
		if err != nil {
			return nil, err
		}
		articles = append(articles, &article)
	}

	return articles, nil
}

// GetArticleByKeyword 模糊查询,主页右上角用
func (a *ArticleImpl) GetArticleByKeyword(keyword string, authorID int, page int) ([]*model.Article, error) {
	var query string
	var args []interface{}

	// 默认页面大小
	pageSize := 10

	// 计算偏移量
	offset := (page - 1) * pageSize

	// 检查 authorID 是否为 1
	if authorID == 1 {
		if keyword == "" {
			query = "SELECT ID, Title, Content, CreateTime, ImageURL, Category, Author, IsDelete, AuthorID FROM news WHERE IsDelete = 0 ORDER BY CreateTime DESC LIMIT ? OFFSET ?"
			args = []interface{}{pageSize, offset}
		} else {
			query = "SELECT ID, Title, Content, CreateTime, ImageURL, Category, Author, IsDelete, AuthorID FROM news WHERE (Title LIKE ? OR Content LIKE ? OR Category LIKE ? OR Author LIKE ?) AND IsDelete = 0 ORDER BY CreateTime DESC LIMIT ? OFFSET ?"
			args = []interface{}{"%" + keyword + "%", "%" + keyword + "%", "%" + keyword + "%", "%" + keyword + "%", pageSize, offset}
		}
	} else {
		if keyword == "" {
			query = "SELECT ID, Title, Content, CreateTime, ImageURL, Category, Author, IsDelete, AuthorID FROM news WHERE IsDelete = 0 AND AuthorID = ? ORDER BY CreateTime DESC LIMIT ? OFFSET ?"
			args = []interface{}{authorID, pageSize, offset}
		} else {
			query = "SELECT ID, Title, Content, CreateTime, ImageURL, Category, Author, IsDelete, AuthorID FROM news WHERE (Title LIKE ? OR Content LIKE ? OR Category LIKE ? OR Author LIKE ?) AND IsDelete = 0 AND AuthorID = ? ORDER BY CreateTime DESC LIMIT ? OFFSET ?"
			args = []interface{}{"%" + keyword + "%", "%" + keyword + "%", "%" + keyword + "%", "%" + keyword + "%", authorID, pageSize, offset}
		}
	}

	rows, err := a.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []*model.Article
	for rows.Next() {
		var article model.Article
		err = rows.Scan(&article.ID, &article.Title, &article.Content, &article.CreateTime, &article.ImageURL, &article.Category, &article.Author, &article.IsDelete, &article.AuthorID)
		if err != nil {
			return nil, err
		}
		articles = append(articles, &article)
	}

	return articles, nil
}

func (a *ArticleImpl) GetArticleByID(id int) (*model.Article, error) {
	query := "SELECT Title, Content, CreateTime, ImageURL, Category,Author,file FROM news WHERE ID = ?"
	row := a.db.QueryRow(query, id)

	var article model.Article
	err := row.Scan(&article.Title, &article.Content, &article.CreateTime, &article.ImageURL, &article.Category, &article.Author, &article.File)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("找不到对应的文章")
		}
		return nil, fmt.Errorf("查询文章失败：%w", err)
	}

	return &article, nil
}

func (a *ArticleImpl) DeleteArticle(id int, userID int) error {
	var query string
	var args []interface{}

	if userID == 1 {
		query = "UPDATE news SET IsDelete = 1 WHERE id = ?"
		args = append(args, id)
	} else {
		query = "UPDATE news SET IsDelete = 1 WHERE id = ? AND AuthorID = ?"
		args = append(args, id, userID)
	}

	_, err := a.db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (a *ArticleImpl) EditArticleByID(id, authorID int) (*model.Article, error) {
	var query string
	var article model.Article
	var args []interface{}

	if authorID == 1 {
		query = "SELECT ID, Title, Content, CreateTime, ImageURL, Category, Author, AuthorID FROM news WHERE ID = ? AND IsDelete = 0"
		args = append(args, id)

	} else {
		query = "SELECT ID, Title, Content, CreateTime, ImageURL, Category, Author, AuthorID FROM news WHERE ID = ? AND AuthorID = ? AND IsDelete = 0"
		args = append(args, id, authorID)
	}

	row := a.db.QueryRow(query, args...)

	err := row.Scan(&article.ID, &article.Title, &article.Content, &article.CreateTime, &article.ImageURL, &article.Category, &article.Author, &article.AuthorID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("找不到对应的文章")
		}
		return nil, fmt.Errorf("查询文章失败：%w", err)
	}

	return &article, nil
}

func (a *ArticleImpl) SaveEditArticle(article *model.Article) error {
	fmt.Println(article)
	query := "UPDATE news SET Title = ?, Content = ?,  Category = ? WHERE ID = ?"
	result, err := a.db.Exec(query, article.Title, article.Content, article.Category, article.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("文章不存在或未更新任何行")
	}

	return nil
}
