package api

import (
	"encoding/json"
	"fmt"
	"go-web/model"
	"html/template"
	"net/http"
)

type Article interface {
	GetArticleByKeyword(w http.ResponseWriter, r *http.Request)
	CreateArticle(w http.ResponseWriter, r *http.Request)
}

type ArticleImpl struct{}

func (u ArticleImpl) GetArticleByKeyword(w http.ResponseWriter, r *http.Request) {
	// 解析表单数据
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "解析表单数据失败", http.StatusBadRequest)
		return
	}

	// 获取关键词
	keyword := r.Form.Get("keyword")

	articles, err := ArticleServer.GetArticleByKeyword(keyword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 加载模板文件
	tmpl, err := template.ParseFiles("./static/html/html.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 渲染模板并将结果写入 ResponseWriter
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmpl.Execute(w, articles)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
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
