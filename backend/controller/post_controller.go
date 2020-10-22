package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nossey/northernlife/application"
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
