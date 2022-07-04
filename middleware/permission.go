package middleware

import (
	"ginApi/common"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Permission() gin.HandlerFunc {
	return func(c *gin.Context) {

		ginToken := c.Request.Header.Get("GINTOKEN")
		if ginToken == "" {
			common.Failed(c,"token 不存在",1,nil)
			//c.JSON(http.StatusOK,gin.H{"code":1,"msg":"","data":nil})
			c.Abort()
			return
		}
		//此处可以配合jwt解析token
		userId,_ := strconv.Atoi(ginToken)

		// 请求的path
		p := c.Request.URL.Path
		// 请求的方法
		m := c.Request.Method

		role := "guest"
		switch userId {
		case 1:
			role = "superAdmin"
		case 2:
			role="product"
		case 3:
			role="order"
		}


		//role := "superAdmin"
		//role:="product"
		//role:="guest"


		// 检查用户权限
		isPass, err := common.Enforcer.Enforce(role, p, m)
		if err != nil {
			c.Abort()
			common.ResultFailed(c,err.Error(),1,nil)
			return
		}
		if isPass {
			c.Next()
		} else {
			c.Abort()
			common.ResultFailed(c,"无访问权限",-1,nil)
			return
		}
	}
}

