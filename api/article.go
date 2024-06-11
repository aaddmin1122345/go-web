package api

import (
	"encoding/json"
	"fmt"
	"go-web/model"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
)

type Article interface {
	GetArticleByKeyword(w http.ResponseWriter, r *http.Request)
	CreateArticle(w http.ResponseWriter, r *http.Request)
	GetArticleByCategory(w http.ResponseWriter, r *http.Request)
	UploadFile(w http.ResponseWriter, r *http.Request)
}

type ArticleImpl struct{}

func (u ArticleImpl) UploadFile(w http.ResponseWriter, r *http.Request) {
	// ParseMultipartForm解析来自表单的请求，解析的结果将存储在Request的FormFile方法中
	const MaxUploadSize = 10 << 20 // 10MB
	err := r.ParseMultipartForm(MaxUploadSize)
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// 获取上传的文件句柄
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file from form data", http.StatusInternalServerError)
		fmt.Println("Error retrieving file from form data:", err)
		return
	}
	defer func(file multipart.File) {
		err = file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	// 获取当前日期
	now := time.Now()
	year, month, day := now.Date()

	// 创建目录路径
	uploadDir := fmt.Sprintf("./uploads/%d/%02d/%02d", year, month, day)

	// 如果目录不存在，则创建目录
	if _, err = os.Stat(uploadDir); os.IsNotExist(err) {
		if err = os.MkdirAll(uploadDir, 0755); err != nil {
			http.Error(w, "Error creating directory", http.StatusInternalServerError)
			fmt.Println("Error creating directory:", err)
			return
		}
	}

	// 创建一个新文件
	filePath := path.Join(uploadDir, handler.Filename)
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error creating file", http.StatusInternalServerError)
		fmt.Println("Error creating file:", err)
		return
	}
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(f)

	// 将上传的文件内容复制到新文件中
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error copying file", http.StatusInternalServerError)
		fmt.Println("Error copying file:", err)
		return
	}

	response := map[string]string{"message": "文件上传成功", "path": filePath}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
	fmt.Println(err)
}

func (u ArticleImpl) GetArticleByKeyword(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 获取关键词
	keyword := r.FormValue("keyword")

	articles, err := ArticleServer.GetArticleByKeyword(keyword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 加载 HTML 模板
	tmpl, err := template.ParseFiles("./static/html/searchArticle.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 渲染模板并将结果写入 ResponseWriter
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmpl.Execute(w, articles)
	if err != nil {
		http.Error(w, "无法生成HTML", http.StatusInternalServerError)
		return
	}
}

func (u ArticleImpl) CreateArticle(w http.ResponseWriter, r *http.Request) {
	var newArticle model.Article
	err := json.NewDecoder(r.Body).Decode(&newArticle)
	if err != nil {
		http.Error(w, "解码json失败", http.StatusBadRequest)
		return
	}

	// 调用服务层方法添加用户
	//err = ArticleService.AddUser(&newArticle)
	err = ArticleServer.CreateArticle(&newArticle)
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
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 获取关键词
	category := r.FormValue("category")
	//page := r.FormValue("page")
	//pageSize := r.FormValue("pageSize")
	page, err := strconv.Atoi(r.FormValue("page"))
	pageSize, err := strconv.Atoi(r.FormValue("pageSize"))

	fmt.Println(category, page, pageSize)

	_, err = ArticleServer.GetArticleByCategory(category, page, pageSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//func (u ArticleImpl) GetArticleByID(id int, w http.ResponseWriter, r *http.Request) {
//	_, err := ArticleServer.GetArticleByID(id)
//	if err != nil {
//		return
//	}
//}
