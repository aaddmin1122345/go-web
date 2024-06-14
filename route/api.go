package route

//
//import (
//	"github.com/gorilla/mux"
//)
//
//// Api 存放全部的api路由
//func (r MyRouteImpl) Api(router *mux.Router) {
//	router.HandleFunc("/api/login", UserApi.Login).Methods("POST")
//	router.HandleFunc("/api/logout", UserApi.Logout).Methods("POST")
//	router.HandleFunc("/api/getUserByKeyword", UserApi.GetUserByKeyword).Methods("POST")
//	router.HandleFunc("/api/updateUser", UserApi.UpdateUser).Methods("POST")
//	router.HandleFunc("/api/register", UserApi.AddUser).Methods("POST")
//	router.HandleFunc("/api/delUser", UserApi.DeleteUser).Methods("POST")
//	//http.HandleFunc("/api/check", UserApi.ValidUser)
//	router.HandleFunc("/api/createArticle", ArticleApi.CreateArticle)
//	router.HandleFunc("/api/upload", ArticleApi.UploadFile)
//	router.HandleFunc("/api/createComment", CommentApi.AddComment)
//	//http.HandleFunc("/api/test", UserApi.SomeOtherHandler)
//	//http.HandleFunc("/api/getComment", CommentApi.GetComment)
//
//	//http.HandleFunc("/api/getArticleByKeyword", MyTemplate)
//}
