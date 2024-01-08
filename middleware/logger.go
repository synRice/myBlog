package middleware

import (
	"fmt"
	"myBlog/log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GinLogger() gin.HandlerFunc { // 是返回 Gin 框架中间件处理函数的函数，用于记录请求的日志
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		// raw := c.Request.URL.RawQuery // 原始查询参数

		c.Next()

		end := time.Now()
		latency := end.Sub(start)                                 // 计算请求处理时间
		clientIP := c.ClientIP()                                  // 获取客户端 IP 地址
		method := c.Request.Method                                // 请求方法
		statusCode := c.Writer.Status()                           // HTTP 状态码
		comment := c.Errors.ByType(gin.ErrorTypePrivate).String() // 可能的错误信息

		// 使用 Logger 记录整条日志
		log.Logger.Info(fmt.Sprintf(" %3d | %13v | %15s | %-7s |%s | %s",
			statusCode,
			latency,
			clientIP,
			method,
			path,
			//raw,
			comment,
		))
	}
}

func GinRecovery() gin.HandlerFunc { // 用于 Gin 框架的中间件函数，处理请求过程中发生的 panic 并进行恢复
	return func(c *gin.Context) {
		defer func() { // 使用 defer 延迟执行的特性，在匿名函数中尝试捕获 panic，它用于确保即使发生 panic，也能在 panic 引发程序崩溃前执行恢复操作
			if err := recover(); err != nil {
				log.Logger.Error("Panic recovered", zap.Any("error", err))
				c.AbortWithStatus(http.StatusInternalServerError) // 中止请求处理，并返回一个 500 内部服务器错误状态码给客户端
			}
		}()

		c.Next()
	}
}
