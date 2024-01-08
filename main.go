package main

import (
	"myBlog/log"
	"myBlog/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建路由
	r := gin.Default()

	// 注册中间件

	// 绑定路由规则，执行的函数
	log.InitLogger()
	router.InitRouter(r)

	// 监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8000")
}
