package model

import "html/template"

type Article struct {
	ID         int
	Title      string
	Content    template.HTML
	Date       string
	CreateTime string
	ImageURL   string
	Category   string
	Author     string
	IsDelete   int
	AuthorID   int
	File       string
}

type TemplateDataByCategory struct {
	NewArticle         []*Article
	Article            []*Article
	ArticlesByCategory []*Article
	CurrentPage        int
	TotalPages         int
	PrevPage           int
	NextPage           int
	HasPrev            bool
	HasNext            bool
}

type TemplateDataByID struct {
	Article            *Article
	ArticlesByCategory []*Article
	Comment            []*Comment
	ShowComments       bool
}

type AuthArticle struct {
	Articles    []*Article
	TotalPages  int
	CurrentPage int
	Keyword     string
}

type UserID struct {
	UserID int
}
