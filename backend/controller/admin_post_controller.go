package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/ahmetb/go-linq"

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
