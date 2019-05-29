package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/oribe1115/phan-sns-server/model"
)

func CreatePostsTableHandler(c echo.Context) error {
	model.CreatePostsTable()
	return c.String(http.StatusOK, "posts table created")
}
