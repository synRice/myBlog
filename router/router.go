package router

import (
	"myBlog/api"
	"myBlog/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.Use(middleware.GinLogger())
	r.Use(middleware.GinRecovery())
	r.GET("/", api.Hello)
}
