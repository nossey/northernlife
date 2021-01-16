package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nossey/northernlife/application"
	"github.com/nossey/northernlife/model"
)

// AdminTagsController provides endpoints to manage tags
type AdminTagsController struct {
}

// AdminTagsCtrl is singleton of AdminTagsController
var AdminTagsCtrl *AdminTagsController

func init() {
	AdminTagsCtrl = &AdminTagsController{}
}

// GetTags godoc
// @Summary Get Tags
// @Accept json
// @Produce json
// @Success 200 {object} model.AdminTagsGetModel
// @Failure 401 {object} model.UnauthorizedMessage
// @Router /admin/tags [get]
// @Security ApiKeyAuth
// @Tags AdminTags
func (ctrl *AdminTagsController) GetTags(ctx *gin.Context) {
	tagApp := application.TagApp
	tags := tagApp.GetTags()
	ctx.JSON(http.StatusOK, model.AdminTagsGetModel{Tags: tags})
}
