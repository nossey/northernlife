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
// @Success 200 {array} model.Post
// @Router /posts [get]
// @Tags Posts
func (c *Controller) GetPosts(ctx *gin.Context) {
	var posts []model.Post
	posts = append(posts, model.Post{ID: 1})
	posts = append(posts, model.Post{ID: 2})
	posts = append(posts, model.Post{ID: 3})
	ctx.JSON(http.StatusOK, gin.H{"posts": posts})
}
