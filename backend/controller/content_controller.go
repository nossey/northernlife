package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nossey/northernlife/application"
)

// ContentController handles content
type ContentController struct {
}

// ContentCtrl is singleton of ContentController
var ContentCtrl *ContentController

func init() {
	ContentCtrl = &ContentController{}
}

// UploadFile uploads a file
func (c *ContentController) UploadFile(ctx *gin.Context) {
	contentApplication := application.ContentApp
	err := contentApplication.Upload()
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{})
}
