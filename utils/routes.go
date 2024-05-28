package utils

import (
	"go-web/api"
	"net/http"
)

//var UserApi = api.UserApiImpl{}

type Route interface {
	Api()
	Article()
	Init()
}

type RouteImpl struct {
}

func (r RouteImpl) Api() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/sayHello", api.SayHello)
	http.HandleFunc("/api/login", UserApi.Login)
	http.HandleFunc("/api/getUserByKeyword", UserApi.GetUserByKeyword)
	http.HandleFunc("/api/updateUser", UserApi.UpdateUser)
	http.HandleFunc("/api/register", UserApi.AddUser)
	http.HandleFunc("/api/delUser", UserApi.DeleteUser)

}

func (r RouteImpl) Article() {
	http.HandleFunc("/getArticleByKeyword", ArticleApi.GetArticleByKeyword)
	//http.HandleFunc("/api/getArticleByKeyword", ArticleApi.GetArticleByKeyword)
	http.HandleFunc("/api/createArticle", ArticleApi.CreateArticle)
}

func (r RouteImpl) Init() {
	r.Api()
	r.Article()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}

}
