package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// InitMiddleware 接受服务实例，并存到gin.Key中
func InitMiddleware(service []interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {
		// 将实例存在gin.Keys中
		context.Keys = make(map[string]interface{})
		context.Keys["user"] = service[0]
		context.Keys["task"] = service[1]
		context.Keys["python_server"] = service[2]
		context.Keys["java_server"] = service[3]
		context.Keys["book"] = service[4]
		context.Next()
	}
}

// ErrorMiddleware 错误处理中间件
func ErrorMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				context.JSON(200, gin.H{
					"code": 404,
					"msg":  fmt.Sprintf("%s", r),
				})
				context.Abort()
			}
		}()
		context.Next()
	}
}
