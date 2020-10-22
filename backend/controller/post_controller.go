package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nossey/northernlife/application"
)

// GetPosts godoc
// @Summary Retrive posts
// @Accept  json
// @Produce  json
// @Success 200 {array} model.PostListModel
// @Router /posts [get]
// @Tags Posts
func (c *Controller) GetPosts(ctx *gin.Context) {
	postResult := application.GetPosts(1)
	ctx.JSON(http.StatusOK, postResult)
}
