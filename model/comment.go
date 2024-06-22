package model

type Comment struct {
	CommentID       int
	ArticleID       int
	Author          string
	Username        string
	Content         string
	CreateTime      string
	IsDelete        int
	ParentCommentID *int       // 使用指针允许NULL值
	Replies         []*Comment // 保存嵌套评论
}
