package service

import "go-web/model"

type Article interface {
	GetArticleByKeyword(keyword string) ([]*model.Article, error)
	CreateArticle(article *model.Article) error
	GetArticleByCategory(category string, page int, pageSize int) ([]*model.Article, error)
	GetArticleByID(id int) (*model.Article, error)
	CountArticle(category string) (int, error)
}

type ArticleImpl struct{}

func (a *ArticleImpl) CountArticle(category string) (int, error) {
	count, err := MyDatabaseArticle.CountArticle(category)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (a *ArticleImpl) GetArticleByCategory(category string, page int, pageSize int) ([]*model.Article, error) {
	articles, err := MyDatabaseArticle.GetArticleByCategory(category, page, pageSize)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (a *ArticleImpl) GetArticleByKeyword(keyword string) ([]*model.Article, error) {
	articles, err := MyDatabaseArticle.GetArticleByKeyword(keyword)
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
