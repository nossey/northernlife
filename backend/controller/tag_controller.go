package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nossey/northernlife/application"
)

// CreateTagController creates tag controller
func CreateTagController() *TagController {
	return &TagController{}
}

// TagController routes rest api for tags
type TagController struct {
}

// GetTags godocs
// @Summary Get tags
// @Description Get all tags
// @Accept json
// @Produce json
// @Success 200 {array} string
func (c *TagController) GetTags(ctx *gin.Context) {
	tags := application.GetTags()
	ctx.JSON(http.StatusOK, tags)
}
