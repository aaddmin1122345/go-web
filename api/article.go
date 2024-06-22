package api

import (
	"encoding/json"
	"fmt"
	"go-web/model"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

//type Article interface {
//	GetArticleByKeyword(w http.ResponseWriter, r *http.Request)
//	//CreateArticle(w http.ResponseWriter, r *http.Request)
//	GetArticleByCategory(w http.ResponseWriter, r *http.Request)
//	UploadFile(w http.ResponseWriter, r *http.Request)
//}

type ArticleImpl struct{}

func (u ArticleImpl) UploadFile(w http.ResponseWriter, r *http.Request) {
	const MaxUploadSize = 10 << 20 // 10MB
	err := r.ParseMultipartForm(MaxUploadSize)
	if err != nil {
		http.Error(w, "上传文件太大了", http.StatusBadRequest)
		return
	}

	// 获取上传的文件句柄
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "读取数据错误", http.StatusInternalServerError)
		return
	}
	defer func(file multipart.File) {
		err = file.Close()
		if err != nil {
			fmt.Println("关闭文件失败:", err)
		}
	}(file)

	// 检查文件类型是否允许上传
	allowedTypes := map[string]bool{
		".docx": true,
		".pdf":  true,
		".xls":  true,
		".xlsx": true,
		".png":  true,
		".jpg":  true,
		".jpeg": true,
		".gif":  true,
		".bmp":  true,
	}
	fileExt := strings.ToLower(path.Ext(handler.Filename))
	if !allowedTypes[fileExt] {
		http.Error(w, "只允许上传 .docx, .pdf, .xls, .xlsx, .png, .jpg, .jpeg, .gif, .bmp 文件", http.StatusBadRequest)
		return
	}

	// 获取当前日期
	now := time.Now()
	year, month, day := now.Date()

	// 根据日期创建路径
	uploadDir := fmt.Sprintf("./uploads/%d/%02d/%02d", year, month, day)

	// 如果目录不存在，则创建目录
	if _, err = os.Stat(uploadDir); os.IsNotExist(err) {
		if err = os.MkdirAll(uploadDir, 0755); err != nil {
			http.Error(w, "无法创建目录", http.StatusInternalServerError)
			fmt.Println("无法创建目录:", err)
			return
		}
	}

	// 创建一个新文件
	filePath := path.Join(uploadDir, handler.Filename)
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "文件创建失败", http.StatusInternalServerError)
		fmt.Println("文件创建失败:", err)
		return
	}
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			fmt.Println("关闭文件流失败:", err)
		}
	}(f)

	// 将上传的文件内容复制到新文件中
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "复制文件失败", http.StatusInternalServerError)
		fmt.Println("无法复制文件:", err)
		return
	}

	response := map[string]string{"message": "文件上传成功", "path": filePath}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "JSON编码失败", http.StatusInternalServerError)
		fmt.Println("JSON编码失败:", err)
		return
	}
}

func (u ArticleImpl) CreateArticle(w http.ResponseWriter, r *http.Request) {
	var newArticle *model.Article
	err := json.NewDecoder(r.Body).Decode(&newArticle)
	if err != nil {
		http.Error(w, "解码json失败", http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	Sessions, err := userApi.GetSessionInfo(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newArticle.AuthorID = Sessions.UserID
	newArticle.Author = Sessions.Username

	err = ArticleServer.CreateArticle(newArticle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 返回成功响应
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte("文章发布成功"))
	if err != nil {
		fmt.Println("文章发布失败!", err)
		return
	}
}

func (u ArticleImpl) GetArticleByCategory(w http.ResponseWriter, r *http.Request) {
	//Sessions, err := UserAPI.GetSessionInfo(r)
	//author := Sessions.Username
	// 获取关键词
	category := r.FormValue("category")
	//page := r.FormValue("page")
	//pageSize := r.FormValue("pageSize")
	page, err := strconv.Atoi(r.FormValue("page"))
	pageSize, err := strconv.Atoi(r.FormValue("pageSize"))
	//Sessions, err := UserAPI.GetSessionInfo(r)
	//author := Sessions.Username

	_, err = ArticleServer.GetArticleByCategory(category, page, pageSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (u ArticleImpl) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	Sessions, err := userApi.GetSessionInfo(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userID := Sessions.UserID
	id, err := userApi.BodyToInit(r.Body)
	if err != nil {
		http.Error(w, "文章ID必须是整数", http.StatusBadRequest)
		return
	}
	err = ArticleServer.DeleteArticle(id, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
	return

}

func (u ArticleImpl) SaveEditArticle(w http.ResponseWriter, r *http.Request) {
	var article model.Article
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		http.Error(w, "无效的请求体", http.StatusBadRequest)
		return
	}
	//id, err := strconv.Atoi(r.FormValue("id"))
	//article.ID = id
	err = ArticleServer.SaveEditArticle(&article)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 返回成功响应
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "文章编辑成功")
}
