package service

import "go-web/model"

type CommentImpl struct {
}

func (c *CommentImpl) AddComment(comment *model.Comment) error {
	err := MyDatabaseComment.AddComment(comment)
	if err != nil {
		return err
	}
	return nil
}

func (c *CommentImpl) GetComments(article int) ([]*model.Comment, error) {
	comment, err := MyDatabaseComment.GetComments(article)
	if err != nil {
		return nil, err
	}
	return comment, nil

}
