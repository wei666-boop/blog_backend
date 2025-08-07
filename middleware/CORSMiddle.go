package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
浏览器跨域请求问题
浏览器会先发送向服务器发送预检请求
如果允许这样做则才会发送正式请求
*/
func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//允许任何来源的跨域请求
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		//告诉浏览器可以缓存预检请求的结果24小时
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		//允许所有HTTP方法
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		//允许所有请求头字段
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		//允许浏览器在跨域请求是携带认证信息
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		//判断是不是预检请求
		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusOK)
		} else {
			ctx.Next()
		}
	}
}
