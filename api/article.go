package api

import (
	"encoding/json"
	"fmt"
	"go-web/model"
	"html/template"
	"net/http"
	"strconv"
)

type Article interface {
	GetArticleByKeyword(w http.ResponseWriter, r *http.Request)
	CreateArticle(w http.ResponseWriter, r *http.Request)
	GetArticleByID(w http.ResponseWriter, r *http.Request)
	GetArticleByCategory(w http.ResponseWriter, r *http.Request, category string)
}

type ArticleImpl struct{}

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

	articles, err := ArticleServer.GetArticleByCategory(category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 加载 HTML 模板
	tmpl, err := template.ParseFiles("./static/html/test.html")
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

func (u ArticleImpl) GetArticleByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 获取关键词
	ID := r.FormValue("id")
	IntID, err := strconv.Atoi(ID)

	articles, err := ArticleServer.GetArticleByID(IntID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 加载 HTML 模板
	tmpl, err := template.ParseFiles("./static/html/article.html")
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
