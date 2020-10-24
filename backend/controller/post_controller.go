package controller

import (
	"net/http"
	"strconv"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"github.com/nossey/northernlife/application"
	"github.com/nossey/northernlife/model"
)

// GetPosts godoc
// @Summary Get posts with pagination
// @Accept  json
// @Produce  json
// @Success 200 {object} model.PostListModel
// @Router /posts [get]
// @Param page query int false "Page"
// @Tags Posts
func (c *Controller) GetPosts(ctx *gin.Context) {
	pageQuery := ctx.Query("page")
	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		page = 1
	}
	postResult := application.GetPosts(page)
	ctx.JSON(http.StatusOK, postResult)
}

// GetPost godoc
// @Summary Get single post with specific id
// @Accept json
// @Produce json
// @Success 200 {object} model.Post
// @Failure 400 {object} model.ErrorMessage
// @Failure 404 {object} model.ErrorMessage
// @Router /posts/{id} [get]
// @Tags Posts
func (c *Controller) GetPost(ctx *gin.Context) {
	id := ctx.Param("id")
	postID, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorMessage{Message: "Invalid post id"})
		return
	}
	post, err := application.GetPost(postID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, model.ErrorMessage{Message: "Post not found"})
		return
	}
	ctx.JSON(http.StatusOK, post)
}
