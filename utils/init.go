package utils

import (
	"github.com/gorilla/sessions"
	"go-web/api"
	"go-web/service"
)

// 不初始化一个session会报错,暂时先这样写,后面在来细优化
var store = sessions.NewCookieStore([]byte("go-web-session-test"))

var UserApi = &api.UserApiImpl{Session: store}
var ArticleApi = api.ArticleImpl{}
var ServiceArticle = service.ArticleImpl{}
