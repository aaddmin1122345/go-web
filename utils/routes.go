package utils

//
//import (
//	"go-web/api"
//	"go-web/template"
//	"net/http"
//)
//
//type Route interface {
//	Api()
//	Article()
//	Init()
//	login(w http.ResponseWriter, res *http.Request)
//	ArticleByIt(w http.ResponseWriter, res *http.Request)
//	static()
//	getArticleByKeyword(w http.ResponseWriter, res *http.Request)
//	ArticleByMusic(w http.ResponseWriter, res *http.Request)
//	ArticleByShipin(w http.ResponseWriter, res *http.Request)
//	ArticleByPe(w http.ResponseWriter, res *http.Request)
//}
//
//type RouteImpl struct {
//	Template.Template
//}
//
//func (r RouteImpl) login(w http.ResponseWriter, res *http.Request) {
//	if res.Method == http.MethodGet {
//		r.Template.RenderTemplate(w, "./static/html/login.html", nil)
//	}
//}
//
//func (r RouteImpl) register(w http.ResponseWriter, res *http.Request) {
//	if res.Method == http.MethodGet {
//		r.Template.RenderTemplate(w, "./static/html/register.html", nil)
//	}
//}
//
//func (r RouteImpl) ArticleByMusic(w http.ResponseWriter, res *http.Request) {
//	if res.Method != http.MethodGet {
//		return
//	}
//	r.Template.RenderSWFUPage(w, "music")
//}
//
//func (r RouteImpl) ArticleByPe(w http.ResponseWriter, res *http.Request) {
//	if res.Method != http.MethodGet {
//		return
//	}
//	r.Template.RenderSWFUPage(w, "pe")
//}
//
//func (r RouteImpl) ArticleByShipin(w http.ResponseWriter, res *http.Request) {
//	if res.Method != http.MethodGet {
//		return
//	}
//	r.Template.RenderSWFUPage(w, "shipin")
//}
//
//func (r RouteImpl) ArticleByIt(w http.ResponseWriter, res *http.Request) {
//	if res.Method != http.MethodGet {
//		return
//	}
//	r.Template.RenderSWFUPage(w, "it")
//}
//
//func (r RouteImpl) getArticleByKeyword(w http.ResponseWriter, res *http.Request) {
//	ArticleApi.GetArticleByKeyword(w, res)
//}
//
//func (r RouteImpl) static() {
//	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
//}
//
//func (r RouteImpl) Api() {
//	http.HandleFunc("/sayHello", api.SayHello)
//	http.HandleFunc("/api/login", UserApi.Login)
//	http.HandleFunc("/api/logout", UserApi.Logout)
//	http.HandleFunc("/api/getUserByKeyword", UserApi.GetUserByKeyword)
//	http.HandleFunc("/api/updateUser", UserApi.UpdateUser)
//	http.HandleFunc("/api/register", UserApi.AddUser)
//	http.HandleFunc("/api/delUser", UserApi.DeleteUser)
//	http.HandleFunc("/check", UserApi.ValidUser)
//	http.HandleFunc("/api/getArticleByKeyword", r.getArticleByKeyword)
//}
//
//func (r RouteImpl) User() {
//	http.HandleFunc("/user/check", UserApi.ValidUser)
//	http.HandleFunc("/user/login", r.login)
//	http.HandleFunc("/user/register", r.register)
//	http.HandleFunc("/user/logout", UserApi.Logout)
//}
//
//func (r RouteImpl) Article() {
//	http.HandleFunc("/article/getArticleByKeyword", r.getArticleByKeyword)
//	http.HandleFunc("/api/createArticle", ArticleApi.CreateArticle)
//	http.HandleFunc("/article/GetArticleByCategory", ArticleApi.GetArticleByCategory)
//	http.HandleFunc("/article/GetArticleByID", ArticleApi.GetArticleByID)
//}
//
//func (r RouteImpl) Init() {
//	http.HandleFunc("/", r.ArticleByIt)
//	r.Api()
//	r.User()
//	r.static()
//	r.Article()
//	err := http.ListenAndServe(":8080", nil)
//	if err != nil {
//		return
//	}
//}
