package route

import (
	"github.com/gorilla/mux"
)

func (r MyRouteImpl) Article(router *mux.Router) {
	//http.HandleFunc("/article/getArticleByKeyword", r.getArticleByKeyword)
	router.HandleFunc("/article/getArticleByCategory", MyTemplate.GetArticleByCategory).Methods("GET")
	router.HandleFunc("/article/getArticleByID", MyTemplate.GetArticleByID).Methods("GET")
	router.HandleFunc("/api/getArticleByKeyword", ArticleApi.GetArticleByKeyword).Methods("POST")

	//router.HandleFunc("/article/createArticle", MyTemplate.CreateArticle).Methods("GET")
	//router.HandleFunc("/article/createArticle", MyTemplate.CreateArticle).Methods("GET")

	auth := router.PathPrefix("/api").Subrouter()
	auth.Use(r.AuthMiddleware("author"))
	//auth.HandleFunc("/getArticleByKeyword", ArticleApi.GetArticleByKeyword).Methods("POST")
	auth.HandleFunc("/uploadFile", ArticleApi.UploadFile).Methods("POST")
	//auth.HandleFunc("/getArticleByKeyword", ArticleApi.GetArticleByKeyword).Methods("POST")

	article := router.PathPrefix("/article").Subrouter()
	article.Use(r.AuthMiddleware("author"))
	article.HandleFunc("/createArticle", MyTemplate.CreateArticle).Methods("GET")

}
