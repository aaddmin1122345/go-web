package utils

import (
	"go-web/api"
	"net/http"
)

func Handler() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/sayHello", api.SayHello)
	http.HandleFunc("/api/GetUsersByStudID", api.GetUsersByStudID)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
