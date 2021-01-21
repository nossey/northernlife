package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nossey/northernlife/application"
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
// @Success 200 {array} string
// @Router /tags [get]
// @Tags Tags
func (c *TagController) GetTags(ctx *gin.Context) {
	tags := application.TagApp.GetAttachedTags()
	ctx.JSON(http.StatusOK, tags)
}
