package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"

	"github.com/ahmetb/go-linq"
	"github.com/gin-gonic/gin"
	"github.com/nossey/northernlife/application"
	"github.com/nossey/northernlife/domain"
	"github.com/nossey/northernlife/model"
)

// PostController produces endpoints for manage post
type PostController struct {
}

// PostCtrl provides endpoints for post
var PostCtrl *PostController

func init() {
	PostCtrl = &PostController{}
}

// ToPostListItem converts application post to viewmodel post
func ToPostListItem(post domain.PostListItem) (viewmodel model.PostListItem) {
	viewmodel = model.PostListItem{
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
		ID:        strings.Replace(post.ID.String(), "-", "", -1),
		UserID:    post.UserID,
		Title:     post.Title,
		Body:      post.Body,
		PlainBody: post.PlainBody,
		Published: post.Published,
		Thumbnail: post.Thumbnail,
		Tags:      post.Tags,
	}
	return
}

// ToPostSingleItem converts application post to viewmodel post
func ToPostSingleItem(post domain.SinglePostItem) (viewmodel model.PostSingleItem) {
	viewmodel = model.PostSingleItem{
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
		ID:        strings.Replace(post.ID.String(), "-", "", -1),
		UserID:    post.UserID,
		Title:     post.Title,
		Body:      post.Body,
		PlainBody: post.PlainBody,
		Published: post.Published,
		Thumbnail: post.Thumbnail,
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
// @Param tags query []string false "Tags"
// @Param search query string false "Search"
// @Tags Posts
func (postController *PostController) GetPosts(ctx *gin.Context) {
	pageQuery := ctx.Query("page")
	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		page = 1
	}
	tags, _ := ctx.GetQueryArray("tags")
	searchQuery := ctx.Query("search")
	postResult := application.PostApp.GetPostList(tags, page, searchQuery, domain.Published)
	postListViewModel := model.PostListModel{
		TotalCount:   postResult.TotalCount,
		PerPageCount: postResult.PerPageCount,
		Posts:        []model.PostListItem{},
	}
	linq.From(postResult.Posts).SelectT(func(p domain.PostListItem) model.PostListItem {
		return ToPostListItem(p)
	}).ToSlice(&postListViewModel.Posts)

	ctx.JSON(http.StatusOK, postListViewModel)
}

// GetPost godoc
// @Summary Get single post with specific id
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {object} model.PostSingleItem
// @Failure 404 {object} model.ErrorMessage
// @Router /posts/{id} [get]
// @Tags Posts
func (postController *PostController) GetPost(ctx *gin.Context) {
	id := ctx.Param("id")
	postID, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, model.ErrorMessage{Message: "Invalid post id (should be uuid)"})
		return
	}
	post, err := application.GetPost(postID)
	postViewModel := ToPostSingleItem(post)
	if err != nil {
		ctx.JSON(http.StatusNotFound, model.ErrorMessage{Message: "Post not found"})
		return
	}
	ctx.JSON(http.StatusOK, postViewModel)
}
