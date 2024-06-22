package api

import (
	"github.com/gorilla/sessions"
	"go-web/service"
	"io"
	"net/http"
)

var UserService service.UserServiceImpl
var ArticleServer service.ArticleImpl
var CommentServer service.CommentImpl

// var UserAPI = &UserApiImpl{}
var store = sessions.NewCookieStore([]byte("go-web-session-test"))

// var CommentApi = &api.CommentImpl{}
var userApi = &UserApiImpl{Session: store}

type Article interface {
	GetArticleByKeyword(w http.ResponseWriter, r *http.Request)
	GetArticleByCategory(w http.ResponseWriter, r *http.Request)
	UploadFile(w http.ResponseWriter, r *http.Request)
	DeleteArticle(w http.ResponseWriter, r *http.Request)
}

type Comment interface {
	AddComment(w http.ResponseWriter, r *http.Request)
	DeleteComment(w http.ResponseWriter, r *http.Request)
}

type UserApi interface {
	DecodeJson(w http.ResponseWriter, r *http.Request, newData interface{}) error
	BodyToInit(body io.Reader) (int, error)
	Login(w http.ResponseWriter, r *http.Request)
	GetUserByKeyword(w http.ResponseWriter, r *http.Request)
	AddUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
	GetUserType(w http.ResponseWriter, r *http.Request) (string, error)
}
