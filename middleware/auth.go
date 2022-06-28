package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckUserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//
		ginToken := c.Request.Header.Get("GINTOKEN")
		if ginToken == "" {
			c.JSON(http.StatusOK,gin.H{"code":1,"msg":"token 不存在","data":nil})
			c.Abort()
			return
		}
		//此处可以配合jwt
		//解析token
		c.Set("userId",88)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息

	}
}