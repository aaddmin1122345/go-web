package template

import (
	"fmt"
	"go-web/model"
	"net/http"
	"strconv"
)

func (t MyTemplateImpl) GetArticleByKeyword(w http.ResponseWriter, r *http.Request) {
	ArticleApi.GetArticleByKeyword(w, r)
}

func (t MyTemplateImpl) CreateArticle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t.RenderTemplate(w, "./static/html/createArticle.html", nil)
	}
}

func (t MyTemplateImpl) RenderHead(w http.ResponseWriter, _ *http.Request) {

	t.RenderTemplate(w, "./static/html/head.html", nil)
}

func (t MyTemplateImpl) RenderFoot(w http.ResponseWriter, _ *http.Request) {
	t.RenderTemplate(w, "./static/html/foot.html", nil)
}

func (t MyTemplateImpl) GetArticleByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	type TemplateData struct {
		Article            *model.Article
		ArticlesByCategory []*model.Article
	}

	// 获取关键词
	ID := r.FormValue("id")
	IntID, err := strconv.Atoi(ID)
	//
	article, err := ArticleServer.GetArticleByID(IntID)
	articlesByCategory, err := ArticleServer.GetArticleByCategory("")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := TemplateData{
		Article:            article,
		ArticlesByCategory: articlesByCategory,
	}

	t.RenderHead(w, nil)
	t.RenderTemplate(w, "./static/html/article.html", data)
	//t.RenderTemplate(w, "./static/html/latestArticle.html", articlesByCategory)
	//t.RenderTemplate(w, "./static/html/article.html", articlesByCategory)

	//t.RenderTemplate(w, "./static/html/latestArticle.html", articlesByCategory)
	t.RenderFoot(w, nil)
}

func (t MyTemplateImpl) GetArticleByCategory(w http.ResponseWriter, r *http.Request) {
	type TemplateData struct {
		Article            []*model.Article
		ArticlesByCategory []*model.Article
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// 获取关键词
	category := r.FormValue("category")

	article, err := ArticleServer.GetArticleByCategory(category)
	articlesByCategory, err := ArticleServer.GetArticleByCategory("")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := TemplateData{
		Article:            article,
		ArticlesByCategory: articlesByCategory,
	}
	for _, article := range articlesByCategory {
		fmt.Println(*article)

	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 加载 HTML 模板
	t.RenderTemplate(w, "./static/html/category.html", data)
}
