package router

import (
	"net/http"
	"search/handler"
	"search/router/middleware"

	"github.com/gin-gonic/gin"
)

func LoadRouters(r *gin.Engine) {

	//加载中间件和失败路由

	r.Use(middleware.NoCache, middleware.Options, middleware.Secure)
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	//路由

	userrouter := r.Group("/api/v1/map")
	{
		userrouter.POST("", handler.Add)
		userrouter.GET("", handler.Sea)
		userrouter.DELETE("", handler.Del)
	}

}
