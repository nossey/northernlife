package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nossey/northernlife/application"
	"github.com/nossey/northernlife/model"
)

// ContentController handles content
type ContentController struct {
}

// ContentCtrl is singleton of ContentController
var ContentCtrl *ContentController

func init() {
	ContentCtrl = &ContentController{}
}

// UploadImageFile godoc
// @Summary Upload a image file of base64 encoded data url type
// @Accept  json
// @Produce  json
// @Param message body model.FileImageUploadModel true "Image Data"
// @Success 201 {object} model.FileImageUploadSuccessResultModel
// @Failure 400 {object} model.ErrorMessage
// @Failure 500 {object} model.ErrorMessage
// @Router /contents [post]
// @Tags Contents
func (c *ContentController) UploadImageFile(ctx *gin.Context) {
	var json model.FileImageUploadModel
	if err := ctx.ShouldBind(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, &model.ErrorMessage{Message: "Invalid request"})
		return
	}

	contentApplication := application.ContentApp
	url, result := contentApplication.Upload(json.Image)

	switch result {
	case application.Success:
		ctx.JSON(http.StatusCreated, &model.FileImageUploadSuccessResultModel{URL: url})
		break
	case application.InvalidEncodedImage:
		ctx.JSON(http.StatusBadRequest, &model.ErrorMessage{Message: "Invalid encoded image data"})
		break
	case application.S3UploadFailed:
		ctx.JSON(http.StatusInternalServerError, &model.ErrorMessage{Message: "S3 upload failed"})
		break
	}
}
