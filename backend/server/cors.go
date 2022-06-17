package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*, Content-Type, Authorization, X-Requested-With, x-visitor-token, x-token, x-tz")
		c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS, DELETE")
		c.Header("Access-Control-Expose-Headers",
			"Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type",
		)
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		method := c.Request.Method
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
