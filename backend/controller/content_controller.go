package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/nossey/northernlife/application"
)

// ContentController handles content
type ContentController struct {
}

// UploadFile uploads a file
func (c *ContentController) UploadFile(ctx *gin.Context) {
	contentApplication := application.ContentApp
	contentApplication.Upload()
}
