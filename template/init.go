package template

import (
	"go-web/api"
	"go-web/model"
	"go-web/service"
	"html/template"
	"net/http"
)

var ArticleApi = api.ArticleImpl{}
var ArticleServer = service.ArticleImpl{}
var ServiceArticle = service.ArticleImpl{}

type MyTemplate interface {
	Index(w http.ResponseWriter, res *http.Request)
	RenderSwfuPage(w http.ResponseWriter, res *http.Request)
	GetArticleByKeyword(w http.ResponseWriter, res *http.Request)
	CreateArticle(w http.ResponseWriter, res *http.Request)
	RenderHead(w http.ResponseWriter, res *http.Request)
	GetArticleByID(w http.ResponseWriter, res *http.Request)
}

type MyTemplateImpl struct {
}

func (t MyTemplateImpl) Index(w http.ResponseWriter, res *http.Request) {
	// 获取文章数据
	itArticles, err := ServiceArticle.GetArticleByCategory("it")

	peArticles, err := ServiceArticle.GetArticleByCategory("pe")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	shipinArticles, err := ServiceArticle.GetArticleByCategory("shipin")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	musicArticles, err := ServiceArticle.GetArticleByCategory("music")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 定义一个数据结构来存储所有类别的文章数据
	type ArticleData struct {
		It     []*model.Article
		Pe     []*model.Article
		Shipin []*model.Article
		Music  []*model.Article
	}

	data := ArticleData{
		It:     itArticles,
		Pe:     peArticles,
		Shipin: shipinArticles,
		Music:  musicArticles,
	}

	t.RenderTemplate(w, "./static/html/swfu.html", data)

}

func (t MyTemplateImpl) RenderSwfuPage(w http.ResponseWriter, res *http.Request) {
	t.RenderHead(w, nil)
	t.Index(w, nil)
	t.RenderFoot(w, nil)

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

//func (t MyTemplateImpl) RenderTemplate2(w http.ResponseWriter, tmpl string, tmpl2 string, data interface{}, data2 interface{}) {
//	tpl, err := template.ParseFiles(tmpl)
//	tpl2, err := template.ParseFiles(tmpl2)
//
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	w.Header().Set("Content-Type", "text/html; charset=utf-8")
//	err = tpl.Execute(w, data)
//	err = tpl2.Execute(w, data2)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//}
