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
// @Tags Posts
func (postController *PostController) GetPosts(ctx *gin.Context) {
	pageQuery := ctx.Query("page")
	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		page = 1
	}
	postResult := application.GetPosts(page)
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
func (postController *PostController) CreatePost(ctx *gin.Context) {
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

// UpdatePost godoc
// @Summary Update single post
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Param message body model.PostUpdateModel true "Post Data"
// @Success 200 {object} model.PostUpdateResult
// @Failure 400 {object} model.ErrorMessage
// @Failure 401 {object} model.UnauthorizedMessage
// @Failure 404 {object} model.ErrorMessage
// @Router /posts/{id} [put]
// @Security ApiKeyAuth
// @Tags Posts
func (postController *PostController) UpdatePost(ctx *gin.Context) {
	id := ctx.Param("id")
	postID, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, model.ErrorMessage{Message: "Invalid post id (should be uuid)"})
		return
	}

	var json model.PostUpdateModel
	err = ctx.ShouldBindJSON(&json)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorMessage{
			Message: "Invalid Request",
		})
		return
	}

	claims := jwt.ExtractClaims(ctx)
	userID := claims[identityKey].(string)
	postApplication := application.PostApp
	err = postApplication.UpdatePost(application.PostUpdate{
		PostID:    postID,
		UserID:    userID,
		Body:      json.Body,
		Title:     json.Title,
		PlainBody: json.PlainBody,
		Tags:      json.Tags,
		Thumbnail: json.Thumbnail,
		Published: *json.Published,
	})

	if err != nil {
		ctx.JSON(http.StatusNotFound, model.ErrorMessage{
			Message: "Post Not Found",
		})
		return
	}

	ctx.JSON(http.StatusOK, model.PostUpdateResult{
		PostID: id,
	})
}

// DeletePost godoc
// @Summary Delete single post
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Success 204 {object} model.PostDeleteResult
// @Failure 401 {object} model.UnauthorizedMessage
// @Failure 404 {object} model.ErrorMessage
// @Router /posts/{id} [delete]
// @Security ApiKeyAuth
// @Tags Posts
func (postController *PostController) DeletePost(ctx *gin.Context) {
	claims := jwt.ExtractClaims(ctx)
	userID := claims[identityKey].(string)
	postID := ctx.Param("id")

	postApplicaion := application.PostApp
	err := postApplicaion.DeletePost(userID, postID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, model.ErrorMessage{Message: "NotFound Post"})
		return
	}

	ctx.JSON(http.StatusNoContent, model.PostDeleteResult{
		PostID: postID,
	})
}
