package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserCreate is model to create a user
type UserCreate struct {
	UserID   string `json:"user_id" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

// CreateUser creates a user by specified input
func (c *Controller) CreateUser(ctx *gin.Context) {
	var userCreate UserCreate
	if err := ctx.ShouldBindJSON(&userCreate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
