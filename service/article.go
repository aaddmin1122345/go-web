package service

import (
	"go-web/model"
)

type ArticleImpl struct{}

func (a *ArticleImpl) CountArticle(keyword string, authorID int) (int, error) {
	count, err := MyDatabaseArticle.CountArticle(keyword, authorID)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (a *ArticleImpl) CountArticleALL(category string) (int, error) {
	count, err := MyDatabaseArticle.CountArticleALL(category)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (a *ArticleImpl) GetArticleByCategory(category string, page int, pageSize int) ([]*model.Article, error) {
	//var r *http.Request
	//Sessions, err := UserAPI.GetSessionInfo(r)
	//author := Sessions.Username
	articles, err := MyDatabaseArticle.GetArticleByCategory(category, page, pageSize)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (a *ArticleImpl) GetArticleByKeyword(keyword string, authorID int, page int) ([]*model.Article, error) {
	articles, err := MyDatabaseArticle.GetArticleByKeyword(keyword, authorID, page)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (a *ArticleImpl) CreateArticle(article *model.Article) error {
	err := MyDatabaseArticle.CreateArticle(article)
	if err != nil {
		return err
	}
	return nil
}

func (a *ArticleImpl) GetArticleByID(id int) (*model.Article, error) {
	articles, err := MyDatabaseArticle.GetArticleByID(id)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (a *ArticleImpl) DeleteArticle(id int, userID int) error {
	err := MyDatabaseArticle.DeleteArticle(id, userID)
	if err != nil {
		return err
	}
	return nil
}

func (a *ArticleImpl) EditArticleByID(id, authorID int) (*model.Article, error) {
	article, err := MyDatabaseArticle.EditArticleByID(id, authorID)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (a *ArticleImpl) SaveEditArticle(article *model.Article) error {
	err := MyDatabaseArticle.SaveEditArticle(article)
	if err != nil {
		return err
	}
	return nil
}
