package route

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"go-web/api"
	"go-web/template"
	"net/http"
)

// 不初始化一个session会报错,暂时先这样写,后面在来细优化
var store = sessions.NewCookieStore([]byte("go-web-session-test"))
var ArticleService = api.ArticleImpl{}
var CommentApi = &api.CommentImpl{}
var UserApi = &api.UserApiImpl{Session: store}
var ArticleApi = &api.ArticleImpl{}
var MyTemplate = template.MyTemplateImpl{}

type MyRoute interface {
	static(router *mux.Router)
	Init()
	AuthMiddleware(userType string) mux.MiddlewareFunc
	Comment(router *mux.Router)
	Article(router *mux.Router)
	Admin(router *mux.Router)
}

type MyRouteImpl struct{}

func (r MyRouteImpl) static(router *mux.Router) {
	staticDir := "./static"
	staticHandler := http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir)))
	router.PathPrefix("/static/").Handler(staticHandler).Methods("GET")

	uploadsDir := "./uploads"
	uploadsHandler := http.StripPrefix("/uploads/", http.FileServer(http.Dir(uploadsDir)))
	router.PathPrefix("/uploads/").Handler(uploadsHandler).Methods("GET")
}

func (r MyRouteImpl) Init() {
	// 设置路由
	//r.User()
	route := mux.NewRouter()

	// 调用路由注册函数
	r.MyRoutes(route)
	r.Comment(route)
	r.static(route)
	r.Article(route)
	r.Admin(route)

	// 设置根路径的处理函数
	route.HandleFunc("/", MyTemplate.RenderSwfuPage)

	// 启动 HTTP 服务器
	err := http.ListenAndServe("0.0.0.0:8080", route)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// AuthMiddleware 添加鉴权路由,鉴权直接在路由上处理
func (r MyRouteImpl) AuthMiddleware(userType string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			actualUserType, _ := UserApi.GetUserType(w, req)

			switch userType {
			case "admin":
				if actualUserType != "admin" {
					MyTemplate.PermissionDenied(w, req)
					return
				}
			case "author":
				if actualUserType != "admin" && actualUserType != "author" {
					MyTemplate.PermissionDenied(w, req)
					return
				}
			case "reader":
				if actualUserType == "" {
					MyTemplate.PermissionDenied(w, req)
					return
				}
			default:
				MyTemplate.PermissionDenied(w, req)
				return
			}

			next.ServeHTTP(w, req)
		})
	}
}
