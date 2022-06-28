package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckAppInstance() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header
		appKey := header.Get("Ginappkey")
		if appKey =="" {
			c.JSON(http.StatusOK,gin.H{	"code":1,"msg":"appKey not exists","data":nil})
			c.Abort()
			return
		}
		//获取到appKey 时间戳 等做一下校验
		if appKey != "1000000" {
			c.JSON(http.StatusOK,gin.H{"code":1,"msg":"appKey not exists1","data":nil})
			c.Abort()
			return
		}
		//做一些 sign签名与 timestamp 时间戳的校验
		//...

		//c.Next()
	}
}
