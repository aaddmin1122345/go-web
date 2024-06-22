package database

import (
	"database/sql"
	"fmt"
	"go-web/model"
)

type CommentImpl struct {
	db *sql.DB
}

func (c *CommentImpl) SetDb(db *sql.DB) {
	c.db = db
}

func (c *CommentImpl) AddComment(comment *model.Comment) error {
	query := "INSERT INTO comment (ArticleID, Username, Content, Author, ParentCommentID) VALUES (?, ?, ?, ?, ?)"
	_, err := c.db.Exec(query, comment.ArticleID, comment.Username, comment.Content, comment.Author, comment.ParentCommentID)
	if err != nil {
		return err
	}
	return nil
}

// GetComments 查询全部评论
func (c *CommentImpl) GetComments(articleID int) ([]*model.Comment, error) {
	// 定义查询语句
	query := "SELECT * FROM comment WHERE ArticleID = ? AND IsDelete = 0 ORDER BY CreateTime"

	// 执行查询，获取行结果
	rows, err := c.db.Query(query, articleID)
	if err != nil {
		return nil, err
	}

	// 存储所有评论的映射key是int,value是*model.Comment
	commentsMap := make(map[int]*model.Comment)
	// 存储顶级评论,父评论的就是顶级评论,父评论是ParentCommentID
	var rootComments []*model.Comment

	// 遍历查询结果
	for rows.Next() {
		// 复制一份新模型,主要是作用于后面的append
		var comment model.Comment

		// 扫描结果到评论结构体
		err := rows.Scan(&comment.CommentID, &comment.ArticleID, &comment.Username, &comment.Content, &comment.CreateTime, &comment.Author, &comment.IsDelete, &comment.ParentCommentID)
		if err != nil {
			return nil, err
		}

		// 将评论添加到映射中
		//把[comment.CommentID]当作key,&comment当作值
		commentsMap[comment.CommentID] = &comment

		// 如果评论没有父评论，则为顶级评论
		if comment.ParentCommentID == nil {
			rootComments = append(rootComments, &comment)
		}
		fmt.Println(&rootComments)
	}

	// 处理子评论
	for _, comment := range commentsMap {
		// 如果有父评论ID，则添加到父评论的回复列表中
		if comment.ParentCommentID != nil {
			parentComment := commentsMap[*comment.ParentCommentID]
			parentComment.Replies = append(parentComment.Replies, comment)
		}
	}

	// 返回顶级评论
	return rootComments, nil
}

func (c *CommentImpl) DeleteComment(id int) error {
	query := "update comment set IsDelete = 1 where CommentID = ?"
	_, err := c.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil

}
