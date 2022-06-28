package middleware

import (
	 "ginApi/common"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CheckUserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//
		ginToken := c.Request.Header.Get("GINTOKEN")
		if ginToken == "" {
			common.Failed(c,"token 不存在",1,nil)
			//c.JSON(http.StatusOK,gin.H{"code":1,"msg":"","data":nil})
			c.Abort()
			return
		}
		//此处可以配合jwt解析token
		userId,_ := strconv.Atoi(ginToken)
		c.Set("userId",userId)
		//c.Next() // 后续的处理函数可以用过c.Get("userId")来获取当前请求的用户信息

	}
}