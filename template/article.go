package template

import (
	"html/template"
	"net/http"
)

func (t MyTemplateImpl) RenderSWFUPage(w http.ResponseWriter, category string) {
	// 获取文章数据
	articles, err := ServiceArticle.GetArticleByCategory(category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 加载 HTML 模板
	tpl, err := template.ParseFiles("./static/html/swfu.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 渲染模板并将结果写入 ResponseWriter
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tpl.Execute(w, articles)
	if err != nil {
		http.Error(w, "无法生成HTML", http.StatusInternalServerError)
		return
	}
}

func (t MyTemplateImpl) ArticleByMusic(w http.ResponseWriter, res *http.Request) {
	if res.Method != http.MethodGet {
		return
	}
	t.RenderSWFUPage(w, "music")
}

func (t MyTemplateImpl) ArticleByPe(w http.ResponseWriter, res *http.Request) {
	if res.Method != http.MethodGet {
		return
	}
	t.RenderSWFUPage(w, "pe")
}

func (t MyTemplateImpl) ArticleByShipin(w http.ResponseWriter, res *http.Request) {
	if res.Method != http.MethodGet {
		return
	}
	t.RenderSWFUPage(w, "shipin")
}

func (t MyTemplateImpl) ArticleByIt(w http.ResponseWriter, res *http.Request) {
	if res.Method != http.MethodGet {
		return
	}
	t.RenderSWFUPage(w, "it")
}

func (t MyTemplateImpl) getArticleByKeyword(w http.ResponseWriter, res *http.Request) {
	ArticleApi.GetArticleByKeyword(w, res)
}
