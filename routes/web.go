package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/richardktran/MyBlogBE/pkg/app"
)

type WebRoute struct {
}

func NewWebRoute() WebRoute {
	return WebRoute{}
}

func (r WebRoute) Setup(router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		app.ResponseSuccess("Welcome to MyBlogBE").Context(ctx)
	})
}
