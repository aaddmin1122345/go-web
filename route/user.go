package route

import (
	"net/http"
)

func (r MyRouteImpl) static() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
}

func (r MyRouteImpl) User() {
	http.HandleFunc("/user/check", UserApi.ValidUser)
	http.HandleFunc("/user/login", MyTemplate.Login)
	http.HandleFunc("/user/register", MyTemplate.Register)
	http.HandleFunc("/user/logout", UserApi.Logout)
}
