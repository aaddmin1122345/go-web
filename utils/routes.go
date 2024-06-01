package utils

import (
	"go-web/api"
	"html/template"
	"net/http"
)

type Route interface {
	Api()
	Article()
	Init()
	login(w http.ResponseWriter, res *http.Request)
	ArticleByIt(w http.ResponseWriter, res *http.Request)
	static()
	getArticleByKeyword(w http.ResponseWriter, res *http.Request)
}

type RouteImpl struct {
}

func (r RouteImpl) login(w http.ResponseWriter, res *http.Request) {
	if res.Method == http.MethodGet {
		renderTemplate(w, "./static/html/login.html", nil)
	}
}

func (r RouteImpl) register(w http.ResponseWriter, res *http.Request) {
	if res.Method == http.MethodGet {
		renderTemplate(w, "./static/html/register.html", nil)
	}
}

func (r RouteImpl) ArticleByIt(w http.ResponseWriter, res *http.Request) {
	if res.Method != http.MethodGet {
		return
		//renderTemplate(w, "./static/html/swfu.html", nil)
	}
	renderSWFUPage(w, "it")

}

func renderSWFUPage(w http.ResponseWriter, category string) {
	// 这里可以调用获取数据的函数，然后将数据传递给模板进行渲染
	// 在此之前，你需要确保数据库连接和相应的模型等都已准备就绪
	// 以下是一个示例实现，你需要根据实际情况进行调整

	// 获取文章数据

	articles, err := ServiceArticle.GetArticleByCategory(category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 加载 HTML 模板
	tmpl, err := template.ParseFiles("./static/html/swfu.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 渲染模板并将结果写入 ResponseWriter
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmpl.Execute(w, articles)
	if err != nil {
		http.Error(w, "无法生成HTML", http.StatusInternalServerError)
		return
	}
}

func (r RouteImpl) getArticleByKeyword(w http.ResponseWriter, res *http.Request) {

	ArticleApi.GetArticleByKeyword(w, res)

}

func (r RouteImpl) static() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
}

func (r RouteImpl) Api() {
	http.HandleFunc("/sayHello", api.SayHello)
	http.HandleFunc("/api/login", UserApi.Login)
	http.HandleFunc("/api/logout", UserApi.Logout)
	http.HandleFunc("/api/getUserByKeyword", UserApi.GetUserByKeyword)
	http.HandleFunc("/api/updateUser", UserApi.UpdateUser)
	http.HandleFunc("/api/register", UserApi.AddUser)
	http.HandleFunc("/api/delUser", UserApi.DeleteUser)
	http.HandleFunc("/check", UserApi.ValidUser)
	http.HandleFunc("/api/getArticleByKeyword", r.getArticleByKeyword)
}

func (r RouteImpl) User() {
	http.HandleFunc("/user/check", UserApi.ValidUser)
	http.HandleFunc("/user/login", r.login)
	http.HandleFunc("/user/register", r.register)
	http.HandleFunc("/user/logout", UserApi.Logout)
}

func (r RouteImpl) Article() {
	http.HandleFunc("/article/getArticleByKeyword", r.getArticleByKeyword)
	http.HandleFunc("/api/createArticle", ArticleApi.CreateArticle)
	http.HandleFunc("/article/GetArticleByCategory", ArticleApi.GetArticleByCategory)
	http.HandleFunc("/article/GetArticleByID", ArticleApi.GetArticleByID)

}

func (r RouteImpl) Init() {
	http.HandleFunc("/", r.ArticleByIt)
	r.Api()
	r.User()
	r.static()
	r.Article()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
