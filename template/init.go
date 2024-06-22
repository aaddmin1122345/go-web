package template

import (
	"fmt"
	"github.com/gorilla/sessions"
	"go-web/api"
	"go-web/model"
	"go-web/service"
	"html/template"
	"math"
	"net/http"
	"strconv"
)

// 不初始化一个session会报错,暂时先这样写,后面在来细优化
var store = sessions.NewCookieStore([]byte("go-web-session-test"))

// var CommentApi = &api.CommentImpl{}
var UserApi = &api.UserApiImpl{Session: store}

// var UserApi = &api.UserApiImpl{}
var ArticleApi = &api.ArticleImpl{}
var ArticleServer = &service.ArticleImpl{}
var ServiceArticle = service.ArticleImpl{}
var UserServer = service.UserServiceImpl{}
var CommentServer = service.CommentImpl{}

type MyTemplate interface {
	Index(w http.ResponseWriter, res *http.Request)
	RenderSwfuPage(w http.ResponseWriter, res *http.Request)
	GetArticleByKeyword(w http.ResponseWriter, res *http.Request)
	CreateArticle(w http.ResponseWriter, res *http.Request)
	RenderHead(w http.ResponseWriter, res *http.Request)
	GetArticleByID(w http.ResponseWriter, res *http.Request)
	until(start, end int) []int
	NotFoundHandler(w http.ResponseWriter, req *http.Request)
}

type MyTemplateImpl struct {
}

func (t MyTemplateImpl) Index(w http.ResponseWriter, res *http.Request) {
	// 获取文章数据
	itArticles, err := ServiceArticle.GetArticleByCategory("it", 1, 10)

	peArticles, err := ServiceArticle.GetArticleByCategory("pe", 1, 10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	shipinArticles, err := ServiceArticle.GetArticleByCategory("shipin", 1, 10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	musicArticles, err := ServiceArticle.GetArticleByCategory("music", 1, 10)
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
	path := res.URL.Path

	// 如果请求的路径不是根路径，则返回 404
	if path != "/" {
		t.NotFoundHandler(w, res)
		return
	}
	t.RenderHead(w, res)
	t.Index(w, res)
	t.RenderFoot(w, res)

}

func (t MyTemplateImpl) NotFoundHandler(w http.ResponseWriter, req *http.Request) {
	//w.WriteHeader(http.StatusNotFound)

	t.RenderHead(w, req)
	t.RenderTemplate(w, "./static/html/404.html", nil)
	t.RenderFoot(w, req)
	// 设置响应状态码为 404
}

func (t MyTemplateImpl) PermissionDenied(w http.ResponseWriter, req *http.Request) {
	//w.WriteHeader(http.StatusNotFound)

	t.RenderHead(w, req)
	t.RenderTemplate(w, "./static/html/403.html", req)
	t.RenderFoot(w, req)
	// 设置响应状态码为 404
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

func (t MyTemplateImpl) AdminUser(w http.ResponseWriter, req *http.Request) {

	allUser, err := UserServer.GetUserByKeyword("")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//allArticle, err := ArticleServer.GetArticleByCategory("", 1, 10)
	//if err != nil {
	//	t.NotFoundHandler(w, req)
	//	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

	data := &model.ALL{
		ALLUser: allUser,
		//ALLArticle: allArticle,
	}

	t.RenderTemplate(w, "./static/html/admin/user.html", data)
}

func (t MyTemplateImpl) AdminArticle(w http.ResponseWriter, r *http.Request) {
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
	tmpl, err := template.New("article.html").Funcs(template.FuncMap{
		"seq": t.seq,
		"add": t.add,
		"min": t.min,
		"sub": t.sub,
	}).ParseFiles("./static/html/admin/article.html")
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
