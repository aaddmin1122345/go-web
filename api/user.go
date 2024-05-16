package api

import (
	"encoding/json"
	"fmt"
	"go-web/service"
	"net/http"
)

func GetUsersByStudID(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.)

	user := service.GetUsersByStudID(1000)

	// 将学生信息编码为 JSON 格式
	responseData, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 设置响应头
	w.Header().Set("Content-Type", "application/json")

	// 将 JSON 响应写入到 ResponseWriter 中
	_, err = w.Write(responseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}
