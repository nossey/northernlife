package controller

import (
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
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

// CreateTag godoc
// @Summary Get Tags
// @Accept json
// @Produce json
// @Param message body model.AdminTagsCreateModel true "Tag Data"
// @Success 201 {object} model.AdminTagsCreateResult
// @Failure 400 {object} model.ErrorMessage
// @Failure 401 {object} model.UnauthorizedMessage
// @Failure 409 {object} model.ErrorMessage
// @Router /admin/tags [post]
// @Security ApiKeyAuth
// @Tags AdminTags
func (ctrl *AdminTagsController) CreateTag(ctx *gin.Context) {
	var json model.AdminTagsCreateModel
	err := ctx.ShouldBindJSON(&json)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorMessage{
			Message: "Invalid Request",
		})
		return
	}

	claims := jwt.ExtractClaims(ctx)
	userID := claims[identityKey].(string)
	app := application.TagApp
	err = app.CreateTag(application.TagCreate{
		Name:   json.TagName,
		UserID: userID,
	})
	if err != nil {
		ctx.JSON(http.StatusConflict, model.ErrorMessage{
			Message: "Failed to create a tag",
		})
		return
	}

	ctx.JSON(http.StatusCreated, model.AdminTagsCreateResult{})
}
