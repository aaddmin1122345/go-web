package template

import (
	"fmt"
	"go-web/model"
	"html/template"
	"math"
	"net/http"
	"strconv"
)

func (t MyTemplateImpl) AuthorArticle(w http.ResponseWriter, r *http.Request) {
	Sessions, err := UserApi.GetSessionInfo(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	authorID := Sessions.UserID

	// 获取关键词
	keyword := r.FormValue("keyword")
	if keyword == "" {
		keyword = "" // 或者设置为一个默认值，视情况而定
	}
	// 获取页码
	page, err := strconv.Atoi(r.FormValue("page"))
	if err != nil || page < 1 {
		page = 1 // 默认第一页
	}

	pageSize := 10

	// 查询总文章数（带有关键词搜索）
	totalArticles, err := ArticleServer.CountArticle(keyword, authorID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 计算总页数
	totalPages := int(math.Ceil(float64(totalArticles) / float64(pageSize)))

	// 获取当前页的文章列表（带有关键词搜索）
	Articles, err := ArticleServer.GetArticleByKeyword(keyword, authorID, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 数据封装
	data := struct {
		Articles    []*model.Article
		TotalPages  int
		CurrentPage int
		Keyword     string
		HasPrev     bool
		HasNext     bool
		PrevPage    int
		NextPage    int
	}{
		Articles:    Articles,
		TotalPages:  totalPages,
		CurrentPage: page,
		Keyword:     keyword, // 确保 keyword 是字符串类型
		HasPrev:     page > 1,
		HasNext:     page < totalPages,
		PrevPage:    page - 1,
		NextPage:    page + 1,
	}

	fmt.Printf("Keyword value: %s\n", data.Keyword)

	// 加载并解析模板，同时注册辅助函数
	tmpl, err := template.New("authorArticle.html").Funcs(template.FuncMap{
		"seq": t.seq,
		"add": t.add,
		"min": t.min,
		"sub": t.sub,
	}).ParseFiles("./static/html/admin/authorArticle.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 设置响应头
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// 渲染模板并传递数据
	if err = tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (t MyTemplateImpl) RenderHead(w http.ResponseWriter, r *http.Request) bool {
	Sessions, err := UserApi.GetSessionInfo(r)
	UserType := Sessions.UserType
	UserName := Sessions.Username
	fmt.Println(UserType)
	fmt.Println(UserName)
	if err != nil {
		return false
	}
	t.RenderTemplate(w, "./static/html/head.html", Sessions)

	return true
}

func (t MyTemplateImpl) RenderFoot(w http.ResponseWriter, _ *http.Request) {
	t.RenderTemplate(w, "./static/html/foot.html", nil)
}

//func (t MyTemplateImpl) GetArticleByID(w http.ResponseWriter, r *http.Request) {
//
//	//type TemplateDataByID struct {
//	//	Article            *model.Article
//	//	ArticlesByCategory []*model.Article
//	//	Comment            []*model.Comment
//	//	ShowComments       bool
//	//}
//
//	// 获取关键词
//	ID, _ := strconv.Atoi(r.FormValue("id"))
//	//
//	article, err := ArticleServer.GetArticleByID(ID)
//	if err != nil {
//		t.NotFoundHandler(w, r)
//
//		return
//	}
//	articlesByCategory, err := ArticleServer.GetArticleByCategory("", 1, 10)
//	if err != nil {
//		t.NotFoundHandler(w, r)
//		//http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	//id, err := strconv.Atoi(r.URL.Query().Get("id"))
//	//if err != nil {
//	//	http.Error(w, "文章ID必须是整数", http.StatusBadRequest)
//	//	return
//	//}
//
//	comment, err := CommentServer.GetComments(ID)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	userType, _ := UserApi.GetUserType(w, r)
//
//	showComments := userType != ""
//	fmt.Println(showComments)
//
//	data := &model.TemplateDataByID{
//		Article:            article,
//		ArticlesByCategory: articlesByCategory,
//		Comment:            comment,
//		ShowComments:       showComments,
//	}
//
//	fmt.Println(data.Article.File)
//	t.RenderHead(w, r)
//	t.RenderTemplate(w, "./static/html/article.html", data)
//	t.RenderFoot(w, r)
//}

func (t MyTemplateImpl) GetArticleByID(w http.ResponseWriter, r *http.Request) {
	ID, _ := strconv.Atoi(r.FormValue("id"))

	article, err := ArticleServer.GetArticleByID(ID)
	if err != nil {
		t.NotFoundHandler(w, r)
		return
	}

	articlesByCategory, err := ArticleServer.GetArticleByCategory("", 1, 10)
	if err != nil {
		t.NotFoundHandler(w, r)
		return
	}

	comments, err := CommentServer.GetComments(ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userType, _ := UserApi.GetUserType(w, r)
	showComments := userType != ""

	data := &model.TemplateDataByID{
		Article:            article,
		ArticlesByCategory: articlesByCategory,
		Comment:            comments,
		ShowComments:       showComments,
	}

	t.RenderHead(w, r)
	t.RenderTemplate(w, "./static/html/article.html", data)
	t.RenderFoot(w, r)
}

func (t MyTemplateImpl) GetArticleByCategory(w http.ResponseWriter, r *http.Request) {
	// 获取分类和页码
	category := r.FormValue("category")
	page, err := strconv.Atoi(r.FormValue("page"))
	if err != nil || page < 1 {
		page = 1 // 默认第一页
	}

	pageSize := 10
	//keyword := r.FormValue("keyword")

	// 查询总文章数
	totalArticles, err := ArticleServer.CountArticleALL(category)
	if err != nil {
		http.Error(w, "无法获取文章总数", http.StatusInternalServerError)
		return
	}
	fmt.Println(totalArticles)

	// 计算总页数
	totalPages := int(math.Ceil(float64(totalArticles) / float64(pageSize)))

	// 获取当前页的文章列表
	articles, err := ArticleServer.GetArticleByCategory(category, page, pageSize)
	if err != nil {
		http.Error(w, "无法获取文章列表", http.StatusInternalServerError)
		return
	}
	allArticles, err := ArticleServer.GetArticleByCategory("", 1, 10)
	if err != nil {
		http.Error(w, "无法获取文章列表", http.StatusInternalServerError)
		return
	}

	// 准备数据供模板渲染
	data := &model.TemplateDataByCategory{
		NewArticle:  allArticles,
		Article:     articles,
		CurrentPage: page,
		TotalPages:  totalPages,
		PrevPage:    page - 1,
		NextPage:    page + 1,
		HasPrev:     page > 1,
		HasNext:     page < totalPages,
	}

	// 加载并解析模板
	t.RenderHead(w, r)
	tmpl, err := template.New("category.html").Funcs(template.FuncMap{"seq": t.seq, "add": t.add, "min": t.min, "sub": t.sub}).ParseFiles("./static/html/category.html")
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

func (t MyTemplateImpl) seq(start, end int) []int {
	s := make([]int, end-start+1)
	for i := range s {
		s[i] = start + i
	}
	return s
}

func (t MyTemplateImpl) CreateArticle(w http.ResponseWriter, r *http.Request) {
	//t.RenderHead(w, r)
	Sessions, _ := UserApi.GetSessionInfo(r)
	UserID, _ := strconv.Atoi(strconv.Itoa(Sessions.UserID))
	fmt.Println(UserID)
	data := model.UserID{UserID: UserID}

	t.RenderTemplate(w, "./static/html/createArticle.html", data)
	//ArticleApi.UploadFile(w, r)
	//t.RenderFoot(w, r)
}

func (t MyTemplateImpl) UploadFile(w http.ResponseWriter, r *http.Request) {
	ArticleApi.UploadFile(w, r)
}

//func (t MyTemplateImpl) ArticleManager(w http.ResponseWriter, r *http.Request) {
//	//Article:=model.Article{}
//	Article, err := ArticleServer.GetArticleByKeyword("")
//}

//func (t MyTemplateImpl) DeleteArticle(w http.ResponseWriter, r *http.Request) {
//	ArticleApi.DeleteArticle(w, r)
//}

func (t MyTemplateImpl) EditArticle(w http.ResponseWriter, r *http.Request) {
	Sessions, _ := UserApi.GetSessionInfo(r)
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Error(w, "文章ID必须是整数", http.StatusBadRequest)
		return
	}
	article, err := ArticleServer.EditArticleByID(id, Sessions.UserID)
	fmt.Println(article)
	t.RenderTemplate(w, "./static/html/edit.html", article)
}
