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
