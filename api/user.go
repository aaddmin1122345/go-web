package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/sessions"
	"go-web/model"
	"io"
	"net/http"
	"strconv"
	"time"
)

type UserApiImpl struct {
	Session sessions.Store
}

func (u *UserApiImpl) DecodeJson(w http.ResponseWriter, r *http.Request, newData interface{}) error {
	err := json.NewDecoder(r.Body).Decode(newData)
	if err != nil {
		http.Error(w, "解码json失败", http.StatusBadRequest)
		return err
	}
	return nil
}

func (u *UserApiImpl) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var newUser model.Register
	err := u.DecodeJson(w, r, &newUser)
	if err != nil {
		return
	}

	// 调用服务层方法更新用户信息
	err = UserService.UpdateUser(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 返回成功响应
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("修改信息成功"))
	if err != nil {
		fmt.Println("修改信息失败", err)
		return
	}
}

// 为后续查询用户做准备,把body里面的str to int
func (u *UserApiImpl) bodyToInit(body io.Reader) (int, error) {
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

//func (u *UserApiImpl) GetUserByPhoneNum(w http.ResponseWriter, r *http.Request) {
//	//studID, err := u.bodyToInit(r.Body)
//	//if err != nil {
//	//	http.Error(w, err.Error(), http.StatusBadRequest)
//	//	return
//	//}
//
//	//fmt.Println("11111111111111")
//	//fmt.Println(studID)
//
//	data, err := io.ReadAll(r.Body)
//	if err != nil {
//		return
//	}
//	user := UserService.GetUserByPhoneNum(string(data))
//
//	//解码json
//	responseData, err := json.Marshal(user)
//	//http.StatusInternalServerError 返回特定类型的报错
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//
//	_, err = w.Write(responseData)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//}

func (u *UserApiImpl) GetUserByKeyword(w http.ResponseWriter, r *http.Request) {
	// 解析请求体中的 JSON 数据到 requestData 变量中
	var requestData struct {
		Username string `json:"username"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "解码json失败", http.StatusBadRequest)
		return
	}

	users, err := UserService.GetUserByKeyword(requestData.Username)

	//fmt.Println(&users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 将结果编码为 JSON 格式并发送到客户端
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, "无法编码数据", http.StatusInternalServerError)
		return
	}
}

func (u *UserApiImpl) AddUser(w http.ResponseWriter, r *http.Request) {
	var newUser model.Register
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "解码json失败", http.StatusBadRequest)
		return
	}

	// 调用服务层方法添加用户
	err = UserService.AddUser(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 返回成功响应
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte("用户添加成功"))
	if err != nil {
		fmt.Println("添加失败!", err)
		return
	}
}

func (u *UserApiImpl) DeleteUser(w http.ResponseWriter, r *http.Request) {
	studID, err := u.bodyToInit(r.Body)

	fmt.Println(studID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(studID)
	err = UserService.DeleteUser(studID)
	if err != nil {
		http.Error(w, "删除用户失败", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("删除成功"))
	if err != nil {
		return
	}
}

func (u *UserApiImpl) Login(w http.ResponseWriter, r *http.Request) {
	// 解析请求体中的 JSON 数据到 requestData 变量中
	var requestData model.Login
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "解码json失败", http.StatusBadRequest)
		return
	}

	// 验证逻辑，报错返回401
	user, err := UserService.Login(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// 创建或获取 session
	session, err := u.Session.Get(r, "session")
	if err != nil {
		http.Error(w, "无法获取session", http.StatusInternalServerError)
		return
	}

	// 将用户信息存储到 session 中
	session.Values["username"] = user.Username
	session.Values["usertype"] = user.UserType

	// 设置 session 的 Options
	if requestData.RememberMe {
		session.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   168 * 3600, // 记住 7 天，以秒为单位
			HttpOnly: true,
		}
	} else {
		session.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   0, // 关闭浏览器后失效
			HttpOnly: true,
		}
	}

	// 保存 session
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, "无法保存session", http.StatusInternalServerError)
		return
	}

	// 构造响应数据
	responseData := struct {
		Valid bool `json:"valid"`
	}{
		Valid: user != nil,
	}

	// 编码为 JSON 格式
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(responseData)
	if err != nil {
		http.Error(w, "无法编码数据", http.StatusInternalServerError)
		return
	}
}

func (u *UserApiImpl) GetSessionInfo(r *http.Request) (model.SessionInfo, error) {
	var sessionInfo model.SessionInfo

	// 获取 session
	session, err := u.Session.Get(r, "session")
	if err != nil {
		return sessionInfo, err
	}

	// 从 session 中读取用户名和用户类型
	username, usernameOK := session.Values["username"].(string)
	usertype, usertypeOK := session.Values["usertype"].(string)

	if !usernameOK || !usertypeOK || username == "" || usertype == "" {
		return sessionInfo, errors.New("未授权访问")
	}

	// 将读取到的值赋给 sessionInfo
	sessionInfo = model.SessionInfo{
		Username: username,
		UserType: usertype,
	}

	return sessionInfo, nil
}

func (u *UserApiImpl) Logout(w http.ResponseWriter, r *http.Request) {
	// 删除名为 "session" 的cookie，使其立即过期
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    "",
		Path:     "/",                          // 确保路径正确，通常设置为根路径"/"，适用于整个网站
		Expires:  time.Now().AddDate(0, 0, -1), // 设置过期时间为当前时间之前的一天，即立即过期
		HttpOnly: true,                         // 确保cookie不能被客户端脚本访问
	})

	// 返回注销成功的响应
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("注销成功"))
	if err != nil {
		http.Error(w, "无法发送响应", http.StatusInternalServerError)
		return
	}
}

func (u *UserApiImpl) GetUserType(w http.ResponseWriter, r *http.Request) (string, error) {
	// 获取 session
	sessionInfo, _ := u.GetSessionInfo(r)

	// 返回当前用户的 userType
	return sessionInfo.UserType, nil
}

//func (u *UserApiImpl) bodyToInit(body io.Reader) (string, error) {
//	data, err := io.ReadAll(body)
//	if err != nil {
//		return "", err
//	}
//	return string(data), nil
//}
