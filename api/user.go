package api

import (
	"encoding/json"
	"go-web/model"
	"go-web/service"
	"io"
	"net/http"
	"strconv"
)

var UserService service.UserServiceImpl

type UserApi interface {
	bodyToInit(body io.Reader) (int, error)
	GetUsersByStudID(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type UserApiImpl struct{}

// 为后续查询用户做准备,把body里面的str to int
func (u UserApiImpl) bodyToInit(body io.Reader) (int, error) {
	data, err := io.ReadAll(body)
	if err != nil {
		return 0, err
	}

	value, err := strconv.Atoi(string(data))
	if err != nil {
		return 0, err
	}

	return value, nil
}

func (u UserApiImpl) GetUsersByStudID(w http.ResponseWriter, r *http.Request) {
	studID, err := u.bodyToInit(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//fmt.Println("11111111111111")
	//fmt.Println(studID)

	user := UserService.GetUsersByStudID(studID)

	//解码json
	responseData, err := json.Marshal(user)
	//http.StatusInternalServerError 返回特定类型的报错
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(responseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (u UserApiImpl) Login(w http.ResponseWriter, r *http.Request) {
	// 解析请求体中的 JSON 数据到 requestData 变量中
	var requestData model.Login
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		//http.StatusBadRequest 返回特定错误
		http.Error(w, "解码json失败", http.StatusBadRequest)
		return
	}

	// 验证逻辑,报错返回401
	user, err := UserService.Login(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// 构造响应数据
	//匿名结构体 responseData
	//user 不为 nil，说明登录成功，Valid 为 true
	responseData := struct {
		Valid bool `json:"valid"`
	}{
		Valid: user != nil,
	}

	//编码为JSON格式
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(responseData)
	if err != nil {
		http.Error(w, "无法编码数据", http.StatusInternalServerError)
		return
	}
}
