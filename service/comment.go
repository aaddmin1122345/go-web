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

func (c *CommentImpl) GetComments(articleID int) ([]*model.Comment, error) {
	comments, err := MyDatabaseComment.GetComments(articleID)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (c *CommentImpl) DeleteComment(id int) error {
	err := MyDatabaseComment.DeleteComment(id)
	if err != nil {
		return err
	}
	return nil

}
