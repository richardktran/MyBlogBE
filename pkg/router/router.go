package router

import (
	"github.com/gin-gonic/gin"
	"github.com/richardktran/MyBlogBE/pkg/app"
	"github.com/richardktran/MyBlogBE/routes"
	"go.uber.org/fx"
)

// Module exports dependency to container
var Module = fx.Options(
	fx.Provide(routes.NewApiV1Route),
	fx.Provide(routes.NewWebRoute),
)

type Routes []Route

type Route interface {
	Setup(*gin.Engine)
}

// NewRoutes sets up routes
func NewRoutes(
	apiRoute routes.ApiV1Route,
	webRoute routes.WebRoute,
) Routes {
	return Routes{
		apiRoute,
		webRoute,
	}
}

// Setup all the route
func (r Routes) Setup() *gin.Engine {
	if app.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	}
	route := gin.Default()

	return route
}

func (r Routes) RegisterRoutes(routeInit *gin.Engine) *gin.Engine {
	for _, route := range r {
		route.Setup(routeInit)
	}

	return routeInit
}
