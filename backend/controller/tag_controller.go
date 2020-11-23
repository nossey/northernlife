package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nossey/northernlife/application"
	"github.com/nossey/northernlife/model"
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
// @Router /tags [get]
// @Tags Tags
func (c *TagController) GetTags(ctx *gin.Context) {
	tags := application.GetTags()
	ctx.JSON(http.StatusOK, tags)
}

// CreateTag godocs
// @Summary Create tag
// @Description Create a new tag
// @Accept json
// @Produce json
// @Success 200 {object} model.TagCreatedResult
// @Failure 400 {object} model.TagCreateFailResult
// @Failure 401 {object} model.UnauthorizedMessage
// @Failure 409 {object} model.TagCreateFailResult
// @Router /tags [post]
// @Tags Tags
func (c *TagController) CreateTag(ctx *gin.Context) {
	var createRequest model.TagCreateRequest
	if err := ctx.ShouldBind(&createRequest); err != nil {
		result := model.TagCreateFailResult{
			Message: "Tag creation failed",
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	user, _ := ctx.Get(identityKey)
	userID := user.(*model.User).UserID

	err := application.CreateTag(createRequest.TagName, userID)
	if err != nil {
		result := model.TagCreateFailResult{
			Message: "Tag " + createRequest.TagName + " is already exists.",
		}
		ctx.JSON(http.StatusConflict, result)
		return
	}
	result := model.TagCreatedResult{
		TagName: createRequest.TagName,
	}
	ctx.JSON(http.StatusCreated, result)
}
