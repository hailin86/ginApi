package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(c *gin.Context,data interface{})  {
	Failed(c,"success",200,data)
}

func Failed(c *gin.Context,msg string,code int,data interface{})  {
	c.JSON(http.StatusOK,gin.H{
		"msg":msg,
		"code":code,
		"data":data,
		"trace_id":c.GetString(TraceKey),
	})
}


func ResultSuccess(c *gin.Context,data interface{})  {
	c.JSON(http.StatusOK,gin.H{
		"msg":"success",
		"code":200,
		"data":data,
		"trace_id":c.GetString(TraceKey),
	})
}

func ResultFailed(c *gin.Context,msg string,code int,data interface{})  {
	c.JSON(http.StatusOK,gin.H{
		"msg":msg,
		"code":code,
		"data":data,
		"trace_id":c.GetString(TraceKey),
	})
}