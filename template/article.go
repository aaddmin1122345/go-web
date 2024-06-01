package template

import (
	"net/http"
	"strconv"
)

func (t MyTemplateImpl) GetArticleByKeyword(w http.ResponseWriter, res *http.Request) {
	ArticleApi.GetArticleByKeyword(w, res)
}

func (t MyTemplateImpl) CreateArticle(w http.ResponseWriter, res *http.Request) {
	if res.Method == http.MethodGet {
		t.RenderTemplate(w, "./static/html/createArticle.html", nil)
	}
}

func (t MyTemplateImpl) RenderHead(w http.ResponseWriter, _ *http.Request) {

	t.RenderTemplate(w, "./static/html/head.html", nil)
}

func (t MyTemplateImpl) RenderFoot(w http.ResponseWriter, _ *http.Request) {
	t.RenderTemplate(w, "./static/html/foot.html", nil)
}

func (t MyTemplateImpl) GetArticleByID(w http.ResponseWriter, res *http.Request) {
	if res.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 获取关键词
	ID := res.FormValue("id")
	IntID, err := strconv.Atoi(ID)

	articles, err := ArticleServer.GetArticleByID(IntID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.RenderHead(w, nil)
	t.RenderTemplate(w, "./static/html/article.html", articles)
	t.RenderFoot(w, nil)
}
