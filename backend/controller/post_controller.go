package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"

	jwt "github.com/appleboy/gin-jwt/v2"

	"github.com/ahmetb/go-linq"
	"github.com/gin-gonic/gin"
	"github.com/nossey/northernlife/application"
	"github.com/nossey/northernlife/model"
)

// ToPostViewModel converts application post to viewmodel post
func ToPostViewModel(post application.Post) (viewmodel model.Post) {
	viewmodel = model.Post{
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
		ID:        strings.Replace(post.ID.String(), "-", "", -1),
		UserID:    post.UserID,
		Title:     post.Title,
		Body:      post.Body,
		PlainBody: post.PlainBody,
		Published: post.Published,
		Tags:      post.Tags,
	}
	return
}

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
	postListViewModel := model.PostListModel{
		TotalCount:   postResult.TotalCount,
		PerPageCount: postResult.PerPageCount,
		Posts:        []model.Post{},
	}
	linq.From(postResult.Posts).SelectT(func(p application.Post) model.Post {
		return ToPostViewModel(p)
	}).ToSlice(&postListViewModel.Posts)

	ctx.JSON(http.StatusOK, postListViewModel)
}

// GetPost godoc
// @Summary Get single post with specific id
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {object} model.Post
// @Failure 404 {object} model.ErrorMessage
// @Router /posts/{id} [get]
// @Tags Posts
func (c *Controller) GetPost(ctx *gin.Context) {
	id := ctx.Param("id")
	postID, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, model.ErrorMessage{Message: "Invalid post id (should be uuid)"})
		return
	}
	post, err := application.GetPost(postID)
	postViewModel := ToPostViewModel(post)
	if err != nil {
		ctx.JSON(http.StatusNotFound, model.ErrorMessage{Message: "Post not found"})
		return
	}
	ctx.JSON(http.StatusOK, postViewModel)
}

// CreatePost godoc
// @Summary Create single post
// @Accept json
// @Produce json
// @Param message body model.PostCreateBody true "Post Data"
// @Success 200 {object} model.PostCreateResult
// @Failure 400 {object} model.ErrorMessage
// @Failure 401 {object} model.UnauthorizedMessage
// @Router /posts [post]
// @Security ApiKeyAuth
// @Tags Posts
func (c *Controller) CreatePost(ctx *gin.Context) {
	claims := jwt.ExtractClaims(ctx)

	var json model.PostCreateBody
	if err := ctx.ShouldBindJSON(&json); err != nil {
		errorMessage := model.ErrorMessage{
			Message: err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, errorMessage)
		return
	}
	create := application.PostCreate{
		UserID:    claims[identityKey].(string),
		Title:     json.Title,
		Body:      json.Body,
		PlainBody: json.PlainBody,
		Published: true,
		Thumbnail: json.Thumbnail,
		Tags:      json.Tags,
	}
	postID, err := application.CreatePost(create)
	if err != nil {
		errorMessage := model.ErrorMessage{
			Message: err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, errorMessage)
		return
	}
	id := strings.Replace(postID.String(), "-", "", -1)
	successResult := model.PostCreateResult{
		PostID: id,
	}
	ctx.JSON(http.StatusOK, successResult)
}
