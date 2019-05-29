package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Post struct {
	UserID    string    `json:"user_name"`
	PostID    string    `json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
	Content   string    `json:"content"`
	DateInGame
	CoopStatus
}

type NewPostData struct {
	UserID  string `json:"user_name"`
	Content string `json:"content"`
	DateInGame
	CoopStatus
}

func CreatePostsTable() {
	db.CreateTable(&Post{})
}

func NewPost(postedData NewPostData) error {
	newPost := Post{}
	newPost.UserID = postedData.UserID
	newPost.Content = postedData.Content
	newPost.DateInGame = postedData.DateInGame
	newPost.CoopStatus = postedData.CoopStatus

	u, err := uuid.NewRandom()
	if err != nil {
		return errors.New("faild to make uuid")
	}
	newPost.PostID = u.String()

	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	newPost.CreatedAt = time.Now().In(jst)

	errDB := db.Table("posts").Create(&newPost)
	if errDB.Error != nil {
		return errors.New("faild to post to DB")
	}

	return nil
}
