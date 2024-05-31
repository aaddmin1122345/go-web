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
	root(w http.ResponseWriter, res *http.Request)
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

func (r RouteImpl) root(w http.ResponseWriter, res *http.Request) {
	if res.Method == http.MethodGet {
		renderTemplate(w, "./static/html/swfu.html", nil)
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
	http.HandleFunc("/user/register", UserApi.AddUser)
	http.HandleFunc("/user/logout", UserApi.Logout)
}

func (r RouteImpl) Article() {
	http.HandleFunc("/article/getArticleByKeyword", r.getArticleByKeyword)
	http.HandleFunc("/api/createArticle", ArticleApi.CreateArticle)
	http.HandleFunc("/article/GetArticleByCategory", ArticleApi.GetArticleByCategory)
	http.HandleFunc("/article/GetArticleByID", ArticleApi.GetArticleByID)

}

func (r RouteImpl) Init() {
	http.HandleFunc("/", r.root)
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
