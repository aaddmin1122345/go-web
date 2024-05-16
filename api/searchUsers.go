package api

//
//import (
//	"fmt"
//	"io"
//	"net/http"
//	"strconv"
//	"web/database"
//)
//
//func SearchUsers(w http.ResponseWriter, r *http.Request) {
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
