package template

import (
	"go-web/model"
	"html/template"
	"math"
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
	ID, _ := strconv.Atoi(r.FormValue("id"))
	//
	article, err := ArticleServer.GetArticleByID(ID)
	if err != nil {
		t.NotFoundHandler(w, r)

		return
	}
	articlesByCategory, err := ArticleServer.GetArticleByCategory("", 1, 10)
	if err != nil {
		t.NotFoundHandler(w, r)
		//http.Error(w, err.Error(), http.StatusInternalServerError)
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
		CurrentPage        int
		TotalPages         int
		PrevPage           int
		NextPage           int
		HasPrev            bool
		HasNext            bool
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 获取关键词
	category := r.FormValue("category")
	page, err := strconv.Atoi(r.FormValue("page"))
	if err != nil {
		page = 1
	}

	pageSize := 10

	// 查询总文章数
	totalArticles, err := ArticleServer.CountArticle(category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 计算总页数
	totalPages := int(math.Ceil(float64(totalArticles) / float64(pageSize)))

	article, err := ArticleServer.GetArticleByCategory(category, page, pageSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//查询全部数据,不根据分类来排序,不过只是10条
	articlesByCategory, err := ArticleServer.GetArticleByCategory("", page, pageSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := TemplateData{
		Article:            article,
		ArticlesByCategory: articlesByCategory,
		CurrentPage:        page,
		TotalPages:         totalPages,
		PrevPage:           page - 1,
		NextPage:           page + 1,
		HasPrev:            page > 1,
		HasNext:            page < totalPages,
	}

	// 加载并解析模板
	tmpl, err := template.New("category.html").Funcs(template.FuncMap{"until": until, "add": add, "minimum": minimum}).ParseFiles("./static/html/category.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 渲染模板
	if err = tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func until(start, end int) []int {
	var result []int
	for i := start; i <= end; i++ {
		result = append(result, i)
	}
	return result
}

func add(a, b int) int {
	return a + b
}

func minimum(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (t MyTemplateImpl) UploadFile(w http.ResponseWriter, r *http.Request) {
	t.RenderTemplate(w, "./static/html/upload.html", nil)
	ArticleApi.UploadFile(w, r)
}
