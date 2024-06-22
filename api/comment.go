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
	var newComment *model.Comment
	err := json.NewDecoder(r.Body).Decode(&newComment)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "解码JSON失败", http.StatusBadRequest)
		return
	}

	if newComment.ArticleID <= 0 {
		http.Error(w, "未提供有效的文章ID", http.StatusBadRequest)
		return
	}

	sessions, err := userApi.GetSessionInfo(r)
	newComment.Author = sessions.Username
	err = CommentServer.AddComment(newComment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte("评论添加成功"))
	if err != nil {
		fmt.Println("响应写入失败:", err)
		return
	}
}

func (c *CommentImpl) DeleteComment(w http.ResponseWriter, r *http.Request) {
	id, err := userApi.BodyToInit(r.Body)
	if err != nil {
		http.Error(w, "id错误", http.StatusBadRequest)
		return
	}
	err = CommentServer.DeleteComment(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
