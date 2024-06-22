package route

import (
	"github.com/gorilla/mux"
)

func (r MyRouteImpl) MyRoutes(router *mux.Router) {

	// 使用 Gorilla Mux 注册路由
	router.HandleFunc("/user/login", MyTemplate.Login).Methods("GET")
	router.HandleFunc("/user/register", MyTemplate.Register).Methods("GET")
	router.HandleFunc("/api/login", UserApi.Login).Methods("POST")
	router.HandleFunc("/api/register", UserApi.AddUser).Methods("POST")
	//router.HandleFunc("/api/AAtest", UserApi.AAtest).Methods("GET")

	//router.HandleFunc("/user/logout", UserApi.Logout).Methods("POST")

	// 需要鉴权的路由
	auth := router.PathPrefix("/api").Subrouter()
	auth.Use(r.AuthMiddleware("reader"))
	auth.HandleFunc("/logout", UserApi.Logout).Methods("GET")

	//authorAuth := router.PathPrefix("/api").Subrouter()
	//authorAuth.Use(r.AuthMiddleware("author"))
	//authorAuth.HandleFunc("/deleteUser", UserApi.DeleteUser).Methods("POST")

	adminAuth := router.PathPrefix("/api").Subrouter()
	adminAuth.Use(r.AuthMiddleware("admin"))
	adminAuth.HandleFunc("/deleteUser", UserApi.DeleteUser).Methods("POST")
	adminAuth.HandleFunc("/getUserByKeyword", UserApi.GetUserByKeyword).Methods("POST")
	adminAuth.HandleFunc("/updateUser", UserApi.UpdateUser).Methods("POST")

}
