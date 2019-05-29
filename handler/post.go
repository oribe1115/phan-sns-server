package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"

	"github.com/oribe1115/phan-sns-server/model"
)

func CreatePostsTableHandler(c echo.Context) error {
	model.CreatePostsTable()
	return c.String(http.StatusOK, "posts table created")
}

func NewPostHandler(c echo.Context) error {
	postData := model.NewPostData{}
	c.Bind(&postData.Content)

	sess, err := session.Get("sessions", c)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "something wrong in getting session")
	}
	if sess.Values["userID"] == nil {
		return c.String(http.StatusOK, "faild to get userID from session")
	}
	postData.UserID = sess.Values["userID"].(string)

	err = model.NewPost(postData)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "faild to post")
	}

	return c.String(http.StatusOK, "Succeded to post!")
}
