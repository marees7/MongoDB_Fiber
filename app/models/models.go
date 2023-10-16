package models

import "time"

type Article struct {
	ArticleID    int
	Title        string
	Content      string
	Nickname     string
	CreationDate time.Time
}

type Comment struct {
	ArticleID       uint
	CommentID       uint
	ParentCommentID uint
	Content         string
	Nickname        string
	CreationDate    time.Time
	ParentComment   *Comment
}
