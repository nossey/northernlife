package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nossey/northernlife/model"
)

// GetPosts godoc
// @Summary Retrive posts
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Post
// @Router /posts/ [get]
func (c *Controller) GetPosts(ctx *gin.Context) {
	post := model.Post{}
	post.ID = 1
	ctx.JSON(http.StatusOK, post)
}
