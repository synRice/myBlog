package router

import (
	"myBlog/api"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.GET("/", api.Hello)
}
