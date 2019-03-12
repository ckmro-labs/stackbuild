package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/gin"
	h "github.com/laidingqing/stackbuild/api/handlers"
	"github.com/sirupsen/logrus"
)

//CreateRouter create a gin router.
func CreateRouter() http.Handler {
	e := gin.New()
	initMiddlewares(e)
	initRoutes(e)
	return e
}

func initMiddlewares(e gin.IRouter) {
	e.Use(ginrus.Ginrus(logrus.WithField("component", "gin"), time.RFC3339, true))
}

func initRoutes(e gin.IRouter) {
	v1 := e.Group("/v1")
	{
		v1.GET("/status", nil)
	}

	oauth := e.Group("/oauth")
	{
		oauth.GET("/auth/:provider", h.OAuth)
		oauth.GET("/callbacks/:provider", h.OAuthCallback)
	}
}
