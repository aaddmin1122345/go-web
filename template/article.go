package template

import (
	"fmt"
	"go-web/model"
	"html/template"
	"math"
	"net/http"
	"strconv"
)

//func (t MyTemplateImpl) GetArticleByKeyword(w http.ResponseWriter, r *http.Request) {
//	ArticleApi.GetArticleByKeyword(w, r)
//}

//func (t MyTemplateImpl) CreateArticle(w http.ResponseWriter, r *http.Request) {
//	if r.Method == http.MethodGet {
//		t.RenderTemplate(w, "./static/html/createArticle.html", nil)
//	}
//}

func (t MyTemplateImpl) RenderHead(w http.ResponseWriter, _ *http.Request) bool {
	t.RenderTemplate(w, "./static/html/head.html", nil)
	return true
}

func (t MyTemplateImpl) RenderFoot(w http.ResponseWriter, _ *http.Request) {
	t.RenderTemplate(w, "./static/html/foot.html", nil)
}

func (t MyTemplateImpl) GetArticleByID(w http.ResponseWriter, r *http.Request) {

	//type TemplateDataByID struct {
	//	Article            *model.Article
	//	ArticlesByCategory []*model.Article
	//	Comment            []*model.Comment
	//	ShowComments       bool
	//}

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

	//id, err := strconv.Atoi(r.URL.Query().Get("id"))
	//if err != nil {
	//	http.Error(w, "文章ID必须是整数", http.StatusBadRequest)
	//	return
	//}

	comment, err := CommentServer.GetComments(ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userType, _ := UserApi.GetUserType(w, r)

	showComments := userType != ""
	fmt.Println(showComments)

	data := &model.TemplateDataByID{
		Article:            article,
		ArticlesByCategory: articlesByCategory,
		Comment:            comment,
		ShowComments:       showComments,
	}

	t.RenderHead(w, r)
	t.RenderTemplate(w, "./static/html/article.html", data)
	t.RenderFoot(w, r)
}

func (t MyTemplateImpl) GetArticleByCategory(w http.ResponseWriter, r *http.Request) {

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

	data := &model.TemplateDataByCategory{
		Article:            article,
		ArticlesByCategory: articlesByCategory,
		//当前页
		CurrentPage: page,
		//总页面
		TotalPages: totalPages,
		//上一页
		PrevPage: page - 1,
		//下一页
		NextPage: page + 1,
		//是否有后面的页数
		HasPrev: page > 1,
		//后面没页码了
		HasNext: page < totalPages,
	}

	// 加载并解析模板
	t.RenderHead(w, r)
	tmpl, err := template.New("category.html").Funcs(template.FuncMap{"until": t.until, "add": t.add, "minimum": t.min, "sub": t.sub}).ParseFiles("./static/html/category.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 渲染模板
	if err = tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.RenderFoot(w, r)

}

func (t MyTemplateImpl) until(start, end int) []int {
	var result []int
	for i := start; i <= end; i++ {
		result = append(result, i)
	}
	return result
}

func (t MyTemplateImpl) add(a, b int) int {
	return a + b
}

func (t MyTemplateImpl) min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (t MyTemplateImpl) sub(a, b int) int {
	return a - b
}

func (t MyTemplateImpl) CreateArticle(w http.ResponseWriter, r *http.Request) {
	//t.RenderHead(w, r)
	t.RenderTemplate(w, "./static/html/createArticle.html", nil)
	//ArticleApi.UploadFile(w, r)
	//t.RenderFoot(w, r)
}

func (t MyTemplateImpl) UploadFile(w http.ResponseWriter, r *http.Request) {
	ArticleApi.UploadFile(w, r)
}
