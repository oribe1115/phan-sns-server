package model

import (
	"time"
)

type Post struct {
	UserID    string    `json:"user_name"`
	PostID    string    `json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
	Content   string    `json:"content"`
	DateInGame
	CoopStatus
}

func CreatePostsTable() {
	db.CreateTable(&Post{})
}
