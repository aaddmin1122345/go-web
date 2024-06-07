package route

import (
	"net/http"
)

// Api 存放全部的api路由
func (r MyRouteImpl) Api() {
	http.HandleFunc("/api/login", UserApi.Login)
	http.HandleFunc("/api/logout", UserApi.Logout)
	http.HandleFunc("/api/getUserByKeyword", UserApi.GetUserByKeyword)
	http.HandleFunc("/api/updateUser", UserApi.UpdateUser)
	http.HandleFunc("/api/register", UserApi.AddUser)
	http.HandleFunc("/api/delUser", UserApi.DeleteUser)
	http.HandleFunc("/api/check", UserApi.ValidUser)
	http.HandleFunc("/api/createArticle", ArticleApi.CreateArticle)
	//http.HandleFunc("/api/getArticleByKeyword", MyTemplate)
}
