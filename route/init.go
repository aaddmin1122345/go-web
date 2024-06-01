package route

import (
	"github.com/gorilla/sessions"
	"go-web/api"
	"go-web/template"
	"net/http"
)

// 不初始化一个session会报错,暂时先这样写,后面在来细优化
var store = sessions.NewCookieStore([]byte("go-web-session-test"))

var UserApi = &api.UserApiImpl{Session: store}
var ArticleApi = api.ArticleImpl{}
var MyTemplate = template.MyTemplateImpl{}

type MyRoute interface {
}

type MyRouteImpl struct{}

func (r MyRouteImpl) Init() {
	http.HandleFunc("/", MyTemplate.RenderSwfuPage)
	//r.Api()
	r.User()
	r.Api()
	r.static()
	r.Article()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
