package controller

import (
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/nossey/northernlife/application"
	"github.com/nossey/northernlife/model"
)

var authHandler *jwt.GinJWTMiddleware
var identityKey = "id"

func init() {
	authHandler, _ = jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "Northernlife Auth",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					identityKey: v.UserID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &model.User{
				UserID: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var login model.Login
			if err := c.ShouldBind(&login); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := login.UserID
			password := login.Password
			if application.IsValidUser(userID, password) {
				return &model.User{
					UserID: userID,
				}, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
		LoginResponse: func(ctx *gin.Context, status int, token string, expiredAt time.Time) {
			message := model.LoginSuccessMessage{Token: token, ExpiredAt: expiredAt}
			ctx.JSON(http.StatusOK, message)
		},
		Unauthorized: func(ctx *gin.Context, code int, message string) {
			msg := model.UnauthorizedMessage{Code: code, Message: message}
			ctx.JSON(code, msg)
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
	})
}

// Login godoc
// @Summary Login
// @Accept  json
// @Produce  json
// @Param login body model.Login true "Login"
// @Success 200 {object} model.LoginSuccessMessage
// @Failure 401 {object} model.UnauthorizedMessage
// @Router /auth/login [post]
// @Tags Auth
func (c *Controller) Login(ctx *gin.Context) {
	authHandler.LoginHandler(ctx)
}

// GetAuthHandler return global handler
func GetAuthHandler() *jwt.GinJWTMiddleware {
	return authHandler
}

// GetUser godoc
// @Summary Login
// @Accept  json
// @Produce  json
// @Success 200 {object} model.User
// @Router /auth/user [get]
// @Security ApiKeyAuth
// @Tags Auth
func (c *Controller) GetUser(ctx *gin.Context) {
	//claims := jwt.ExtractClaims(ctx)
	user, _ := ctx.Get(identityKey)
	//ctx.JSON(http.StatusOK, gin.H{
	//	"key":    claims[identityKey],
	//	"userID": user.(*model.User).UserID,
	//})
	ctx.JSON(http.StatusOK, model.User{
		UserID: user.(*model.User).UserID,
	})
}
