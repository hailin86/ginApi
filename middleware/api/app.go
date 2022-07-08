package api

import (
	"fmt"
	"ginApi/common"
	"github.com/gin-gonic/gin"
)

func CheckAppInstance() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params = make(map[string]interface{})
		err := c.ShouldBindJSON(&params)
		common.LogInfo(c,"api request params:",params)
		if err != nil {
			common.LogError(c,"api param err:",err)
			common.Failed(c,"api request error",100,nil)
			c.Abort()
		}

		//获取请求方的appKey 并校验
		appKey,ok := params["app_key"].(string)
		fmt.Println(appKey)
		if !ok {
			common.Failed(c,"param app_key not exists",10001,nil)
			c.Abort()
		}
		//通过appKey 查询出服务端是不是服务端下发的 然后取出appKey 对应的appSecret
		appSecret := "123456"
		// do something to check appKey

		//获取请求方的时间戳
		timestamp,ok := params["timestamp"].(string)
		fmt.Println(timestamp)
		if !ok {
			common.Failed(c,"param timestamp not exists",10002,nil)
			c.Abort()
		}
		//比对请求的时间戳与当前时间戳的差距，比如相差10s 则请求不合法
		//do something to check timestamp

		//获取请求方的随机字符串.
		nonceStr,ok := params["nonce_str"].(string)
		fmt.Println(nonceStr)

		if !ok {
			common.Failed(c,"param nonce_str not exists",10003,nil)
			c.Abort()
		}
		//比如可以把该 随机字符串存入redis 校验该随机串是否已经响应过请求
		//do something to check nonceStr


		//获取请求方的sign 然后服务端重新用请求参数加密获取新的sign 与请求方的sign 比对。
		requestSign, ok := params["sign"]
		fmt.Println(requestSign)
		if !ok {
			common.Failed(c,"param sign not exists",10004,nil)
			c.Abort()
		}

		//校验sign
		sign := common.GenerateSign(appSecret,params)

		if sign != requestSign {
			common.Failed(c,"sign error",10005,nil)
			c.Abort()
		}

		delete(params,"nonceStr")
		delete(params,"sign")

		//fmt.Println("api params:",params)
		c.Set("apiParams",params)
		//api.Params = params
	}
}
