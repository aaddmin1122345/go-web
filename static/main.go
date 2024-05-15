package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"web/api"
	"web/database"
)

func main() {
	//s := dbTest()
	//fmt.Println(s)
	http.Handle("/", http.FileServer(http.Dir("./static/")))
	//不够高效所以弃用
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	// 读取并渲染 index.html 文件
	//	http.ServeFile(w, r, "./static/index.html")
	//})

	http.HandleFunc("/sayHello", helloHandler)
	http.HandleFunc("/userValid", userHandler)
	http.HandleFunc("/sum", sumHandler)
	http.HandleFunc("/userIsValid", userIsValid)
	http.HandleFunc("/stud", dbTest)
	http.HandleFunc("/api/searchUsers", api.SearchUsers)
	//检测数据库连接情况
	_, err := database.DbInit()
	if err != nil {
		fmt.Println("Error connecting to database", err)
	}
	if err != nil {
		fmt.Println("查询数据库错误", err)
	}
	fmt.Println("http://127.0.0.1:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("服务启动失败!", err)
	}

}

func helloHandler(w http.ResponseWriter, _ *http.Request) {
	sendJSONResponse(w, map[string]string{"message": "Hello, World!"})
}

func userIsValid(w http.ResponseWriter, r *http.Request) {
	requestData := make(map[string]string)
	//解码数据
	err := json.NewDecoder(r.Body).Decode(&requestData)
	fmt.Println(requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	username, usernameExists := requestData["userName"]
	password, passwordExists := requestData["passWord"]
	if !usernameExists || !passwordExists {
		http.Error(w, "参数不能为空", http.StatusBadRequest)
		return
	}
	//验证username和password是否正确
	//valid := (username == "admin" || username == "sys" || username == "user") && password == "admin"
	valid := (username == "admin" || username == "sys" || username == "user") && password == "admin"

	//一个布尔值,值由变量valid决定
	sendJSONResponse(w, map[string]bool{"valid": valid})

}

func userHandler(w http.ResponseWriter, r *http.Request) {
	//存储从 HTTP 请求中解析出的 JSON 数据,map存储键值对,适合存储json
	requestData := make(map[string]string)
	//解码数据
	err := json.NewDecoder(r.Body).Decode(&requestData)
	fmt.Println(requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	username, usernameExists := requestData["username"]
	password, passwordExists := requestData["password"]
	//判断用户名或者密码是否存在
	if !usernameExists || !passwordExists {
		http.Error(w, "参数不能为空", http.StatusBadRequest)
		return
	}
	//验证username和password是否正确
	valid := username == "admin" && password == "admin"
	//key:value格式,valid:valid
	sendJSONResponse(w, map[string]bool{"valid": valid})
}

func sendJSONResponse(w http.ResponseWriter, data interface{}) {
	//解码为json
	jsonResponse, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func sumHandler(w http.ResponseWriter, r *http.Request) {
	// 存储从 HTTP 请求中解析出的 JSON 数据，将键值对解析为整数类型
	requestData := make(map[string]int)
	err := json.NewDecoder(r.Body).Decode(&requestData)
	fmt.Println(requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	par1, par1Exists := requestData["par1"]
	par2, par2Exists := requestData["par2"]
	if !par1Exists || !par2Exists {
		http.Error(w, "参数不足", http.StatusBadRequest)
		return
	}

	result := par1 + par2
	sendJSONResponse(w, map[string]int{"result": result})
}

type stud struct {
	Id       int
	StudId   string
	Username string
	Sex      string
	Email    string
}

func dbTest(w http.ResponseWriter, _ *http.Request) {
	s := stud{
		Id:       12,
		StudId:   "1234",
		Username: "admin",
		Sex:      "女",
		Email:    "admin@admin.org",
	}
	fmt.Println("这是s\t:", s)

	sendJSONResponse(w, &s)
}

//func searchUsers(w http.ResponseWriter, r *http.Request) {
//	// 读取请求体
//	body, err := io.ReadAll(r.Body)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	fmt.Println(string(body))
//	// 将请求体中的数据转换为整数
//	studId, err := strconv.Atoi(string(body))
//	if err != nil {
//		http.Error(w, "类型转换错误!", http.StatusBadRequest)
//		fmt.Println("类型转换错误!")
//	}
//
//	// 打印查询的数据
//	fmt.Println("查询的studId：", studId)
//
//	// 调用数据库查询函数
//	err = database.DbSelect(studId)
//	if err != nil {
//		fmt.Println("获取参数传入后端出错:", err)
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//}
