package database

import (
	"database/sql"
	"go-web/model"
)

type CommentImpl struct {
	db *sql.DB
}

func (c *CommentImpl) SetDb(db *sql.DB) {
	c.db = db
}

func (c *CommentImpl) AddComment(comment *model.Comment) error {

	query := "INSERT INTO comment (ArticleID, Username, Content) VALUES (?, ?, ?)"
	_, err := c.db.Exec(query, comment.ArticleID, comment.Username, comment.Content)
	if err != nil {
		return err
	}
	return nil
}

// GetComments 查询全部评论
func (c *CommentImpl) GetComments(articleID int) ([]*model.Comment, error) {
	query := "SELECT ArticleID, Username, Content, CreateTime FROM comment WHERE ArticleID = ?"
	rows, err := c.db.Query(query, articleID)
	if err != nil {
		return nil, err
	}

	var comments []*model.Comment
	for rows.Next() {
		var comment model.Comment
		err = rows.Scan(&comment.ArticleID, &comment.Username, &comment.Content, &comment.CreateTime)
		if err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}

	return comments, nil
}
