package template

import (
	"go-web/api"
	"go-web/service"
	"html/template"
	"net/http"
)

var ArticleApi = api.ArticleImpl{}
var ServiceArticle = service.ArticleImpl{}

type MyTemplate interface {
	RenderTemplate(w http.ResponseWriter, tmpl string, data interface{})
	Login(w http.ResponseWriter, res *http.Request)
	Register(w http.ResponseWriter, res *http.Request)
	RenderSWFUPage(w http.ResponseWriter, category string)
	ArticleByPe(w http.ResponseWriter, res *http.Request)
	ArticleByMusic(w http.ResponseWriter, res *http.Request)
	ArticleByShipin(w http.ResponseWriter, res *http.Request)
	ArticleByIt(w http.ResponseWriter, res *http.Request)
	getArticleByKeyword(w http.ResponseWriter, res *http.Request)
}

type MyTemplateImpl struct {
}

func (t MyTemplateImpl) RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tpl, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
