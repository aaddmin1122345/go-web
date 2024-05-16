package api

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Hello, World!",
	}

	// 将响应结构体编码为 JSON 格式
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 设置响应头
	w.Header().Set("Content-Type", "application/json")

	// 将 JSON 响应写入响应体
	w.Write(jsonResponse)
}
