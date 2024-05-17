package utils

import (
	"go-web/api"
	"net/http"
)

var UserApi = api.UserApiImpl{}

type Route interface {
	Handler()
}

type RouteImpl struct {
}

func (r RouteImpl) Handler() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/sayHello", api.SayHello)
	http.HandleFunc("/api/login", UserApi.Login)
	http.HandleFunc("/api/GetUsersByStudID", UserApi.GetUsersByStudID)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
