package route

import (
	"github.com/gorilla/mux"
)

func (r MyRouteImpl) Comment(router *mux.Router) {
	auth := router.PathPrefix("/api").Subrouter()
	auth.Use(r.AuthMiddleware("reader")) // 不登陆就无法使用的功能
	auth.HandleFunc("/createComment", CommentApi.AddComment).Methods("POST")
	auth.HandleFunc("/deleteComment", CommentApi.DeleteComment).Methods("POST")

	//auth.HandleFunc("/createComment", CommentApi.).Methods("POST")

}
