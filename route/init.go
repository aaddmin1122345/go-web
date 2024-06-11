package route

import (
	"fmt"
	"github.com/gorilla/sessions"
	"go-web/api"
	"go-web/template"
	"net/http"
)

// 不初始化一个session会报错,暂时先这样写,后面在来细优化
var store = sessions.NewCookieStore([]byte("go-web-session-test"))

var UserApi = &api.UserApiImpl{Session: store}
var ArticleApi = &api.ArticleImpl{}
var MyTemplate = template.MyTemplateImpl{}

type MyRoute interface {
}

type MyRouteImpl struct{}

func (r MyRouteImpl) static() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))
}

//func (r MyRouteImpl) NotFoundHandler(w http.ResponseWriter, req *http.Request) {
//	http.NotFound(w, req)
//}

func (r MyRouteImpl) Init() {
	// 设置路由
	r.User()
	r.Api()
	r.static()
	r.Article()

	// 设置默认的 404 处理器
	//http.HandleFunc("/404", r.NotFoundHandler)

	// 设置根路径的处理函数
	http.HandleFunc("/", MyTemplate.RenderSwfuPage)

	// 启动 HTTP 服务器
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
