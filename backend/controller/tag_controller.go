package controller

import (
	"net/http"

	"github.com/ahmetb/go-linq"
	"github.com/gin-gonic/gin"
	"github.com/nossey/northernlife/application"
	"github.com/nossey/northernlife/domain"
	"github.com/nossey/northernlife/model"
)

// TagController provides endpoints for tag api
type TagController struct {
}

// TagCtrl is single instance of TagController
var TagCtrl *TagController

func init() {
	TagCtrl = &TagController{}
}

// GetTags godocs
// @Summary Get tags
// @Description Get all tags
// @Accept json
// @Produce json
// @Success 200 {object} model.TagsGetResult
// @Router /tags [get]
// @Tags Tags
func (c *TagController) GetTags(ctx *gin.Context) {
	tags := application.TagApp.GetAttachedTags()
	result := model.TagsGetResult{}
	linq.From(tags.Items).SelectT(func(t domain.GetTagsItem) model.TagsGetItem {
		item := model.TagsGetItem{}
		item.AttachedCount = t.Count
		item.Name = t.TagName
		return item
	}).ToSlice(&result.Tags)
	ctx.JSON(http.StatusOK, result)
}
