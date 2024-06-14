package api

import (
	"encoding/json"
	"fmt"
	"go-web/model"
	"net/http"
)

//type Comment interface {
//	AddComment(w http.ResponseWriter, r *http.Request)
//}

type CommentImpl struct {
}

func (c *CommentImpl) AddComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "只允许 POST 请求", http.StatusMethodNotAllowed)
		return
	}

	// Decode JSON request body into newComment strucnewComment = {model.Comment} t
	var newComment model.Comment
	err := json.NewDecoder(r.Body).Decode(&newComment)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "解码JSON失败", http.StatusBadRequest)
		return
	}

	// Ensure articleId is provided in comment data
	if newComment.ArticleID == 0 {
		http.Error(w, "未提供有效的文章ID", http.StatusBadRequest)
		return
	}

	// Assuming CommentServer is a global or context-scoped instance
	// Call the service layer method to add the comment
	err = CommentServer.AddComment(&newComment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte("评论添加成功"))
	if err != nil {
		fmt.Println("响应写入失败:", err)
		return
	}
}
