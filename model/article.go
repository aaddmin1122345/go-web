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
}

type TemplateDataByCategory struct {
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
