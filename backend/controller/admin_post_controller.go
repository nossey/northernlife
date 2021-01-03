package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/ahmetb/go-linq"
	"github.com/google/uuid"

	jwt "github.com/appleboy/gin-jwt/v2"

	"github.com/gin-gonic/gin"
	"github.com/nossey/northernlife/application"
	"github.com/nossey/northernlife/domain"
	"github.com/nossey/northernlife/model"
)

func toPostListItem(post domain.PostListItem) (viewmodel model.AdminPostListItem) {
	viewmodel = model.AdminPostListItem{
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

// AdminPostController produces endpoints for manage post
type AdminPostController struct {
}

// AdminPostCtrl provides endpoints for post
var AdminPostCtrl *AdminPostController

func init() {
	AdminPostCtrl = &AdminPostController{}
}

// GetAdminPosts godoc
// @Summary Get posts with pagination
// @Accept  json
// @Produce  json
// @Success 200 {object} model.AdminPostListModel
// @Failure 400 {object} model.ErrorMessage
// @Failure 401 {object} model.UnauthorizedMessage
// @Router /admin/posts [get]
// @Param page query int false "Page"
// @Param type query string false "string enums" Enums(all, published, draft)
// @Tags AdminPosts
func (ctrl *AdminPostController) GetAdminPosts(ctx *gin.Context) {
	getType := ctx.Query("type")
	getPostType := domain.All

	switch getType {
	case "draft":
		getPostType = domain.Draft
		break
	case "published":
		getPostType = domain.Published
		break
	case "all":
		getPostType = domain.All
		break
	}

	pageQuery := ctx.Query("page")
	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		page = 1
	}

	claims := jwt.ExtractClaims(ctx)
	userID := claims[identityKey].(string)

	postApplication := application.PostApp
	result := postApplication.GetAdminPosts(page, userID, getPostType)

	viewmodel := model.AdminPostListModel{
		TotalCount:   result.TotalCount,
		PerPageCount: result.PerPageCount,
		Posts:        []model.AdminPostListItem{},
	}

	linq.From(result.Posts).SelectT(func(p domain.PostListItem) model.AdminPostListItem {
		return toPostListItem(p)
	}).ToSlice(&viewmodel.Posts)

	ctx.JSON(http.StatusOK, viewmodel)

	return
}

// CreateAdminPost godocs
// @Summary Create single post
// @Accept json
// @Produce json
// @Param message body model.PostCreateBody true "Post Data"
// @Success 201 {object} model.PostCreateResult
// @Failure 400 {object} model.ErrorMessage
// @Failure 401 {object} model.UnauthorizedMessage
// @Router /posts [post]
// @Security ApiKeyAuth
// @Tags AdminPosts
func (ctrl *AdminPostController) CreateAdminPost(ctx *gin.Context) {
	claims := jwt.ExtractClaims(ctx)
	postApp := application.PostApp

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
	postID, err := postApp.CreatePost(create)
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
	ctx.JSON(http.StatusCreated, successResult)
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
// @Router /admin/posts/{id} [put]
// @Security ApiKeyAuth
// @Tags AdminPosts
func (ctrl *AdminPostController) UpdatePost(ctx *gin.Context) {
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
		Published: json.Published,
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
