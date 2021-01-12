package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nossey/northernlife/controller"
	"github.com/nossey/northernlife/infrastructure"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9000
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// User is a single result
type User struct {
	ID             string `gorm:"id"`
	Email          string `gorm:"email"`
	HashedPassword string `gorm:"hashed_password"`
}

func main() {
	defer infrastructure.Db.Close()

	r := gin.Default()

	handler := controller.GetAuthHandler()
	v1 := r.Group("/api/v1")
	{
		post := v1.Group("/posts")
		{
			pc := controller.PostCtrl
			post.GET("", pc.GetPosts)
			post.GET("/:id", pc.GetPost)
			post.Use(handler.MiddlewareFunc())
			{
				post.POST("", pc.CreatePost)
				post.PUT("/:id", pc.UpdatePost)
				post.DELETE("/:id", pc.DeletePost)
			}
		}

		auth := v1.Group("/auth")
		{
			ac := controller.AuthCtrl
			auth.POST("/login", ac.Login)
			auth.POST("/logout", handler.LogoutHandler)
			auth.Use(handler.MiddlewareFunc())
			{
				auth.GET("/user", ac.GetUser)
			}
			auth.GET("/refresh", ac.Refresh)
		}

		tag := v1.Group("/tags")
		{
			tc := controller.TagCtrl
			tag.GET("", tc.GetTags)
			tag.Use(handler.MiddlewareFunc())
			{
				tag.POST("", tc.CreateTag)
			}
		}

		content := v1.Group("/contents")
		{
			cc := controller.ContentCtrl
			content.Use(handler.MiddlewareFunc())
			{
				content.POST("", cc.UploadImageFile)
			}
		}

		admin := v1.Group("/admin")
		{
			admin.Use(handler.MiddlewareFunc())
			{
				adminPost := admin.Group("/posts")
				adminPost.Group("")
				{
					apc := controller.AdminPostCtrl
					adminPost.GET("", apc.GetAdminPosts)
					adminPost.GET("/:id", apc.GetAdminPost)
					adminPost.POST("", apc.CreateAdminPost)
					adminPost.PUT("/:id", apc.UpdatePost)
					adminPost.DELETE("/:id", apc.DeletePost)
				}
			}
		}
	}

	srv := &http.Server{
		Addr:    ":9000",
		Handler: r,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds. quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shuting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
}
